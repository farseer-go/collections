package collections

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/farseer-go/fs/parse"
	"strings"
	"sync"
)

// List 集合
type List[T any] struct {
	source        *[]T // 集合
	IList[T]           // 对集合做修改操作
	Enumerable[T]      // 对集合做读操作
}

// NewList 创建集合
func NewList[T any](source ...T) List[T] {
	var lst = List[T]{}
	lst.New()
	lst.Add(source...)
	return lst
}

// AsEnumerable 返回enumerable类型
func (receiver *List[T]) AsEnumerable() Enumerable[T] {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	source := *receiver.source
	return Enumerable[T]{
		source: &source,
		lock:   &sync.RWMutex{},
	}
}

// New 返回enumerable类型
func (receiver *List[T]) New() {
	if receiver.source == nil {
		var lock sync.RWMutex
		source := &[]T{}

		receiver.source = source
		receiver.IList.source = source
		receiver.IList.Collection.source = source
		receiver.IList.Collection.lock = &lock
		receiver.Enumerable.source = source
		receiver.Enumerable.lock = &lock
	}
}

// Value return json value, implement driver.Valuer interface
func (receiver List[T]) Value() (driver.Value, error) {
	if receiver.source == nil {
		//return nil, nil
		return "{}", nil
	}
	ba, err := receiver.MarshalJSON()
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (receiver *List[T]) Scan(val any) error {
	if val == nil {
		*receiver = NewList[T]()
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}

	// 按,号分隔
	*receiver = NewList[T]()
	dataSplits := strings.Split(string(ba), ",")
	var defValue T
	for _, data := range dataSplits {
		receiver.Add(parse.Convert(data, defValue))
	}
	return nil
}

// IsNil 是否未初始化
func (receiver *List[T]) IsNil() bool {
	return receiver.source == nil
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver List[T]) MarshalJSON() ([]byte, error) {
	if receiver.IsNil() {
		return []byte("{}"), nil
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()
	return json.Marshal(receiver.source)
}

// UnmarshalJSON to deserialize []byte
func (receiver *List[T]) UnmarshalJSON(b []byte) error {
	if receiver.IsNil() {
		receiver.New()
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()
	return json.Unmarshal(b, &receiver.source)
}
