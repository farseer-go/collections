package collections

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"sync"

	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snc"
)

// List 集合
type List[T any] struct {
	source        *[]T // 集合
	IList[T]           // 对集合做修改操作
	Enumerable[T]      // 对集合做读操作
}

// ToList 将arr按,号分隔，转换成List[T]
func ToList[T any](arr string) List[T] {
	return parse.Convert(arr, NewList[T]())
}

// NewList 创建集合
func NewList[T any](source ...T) List[T] {
	var lst = List[T]{}
	var lock sync.RWMutex

	//*receiver.source = make([]T, len(item))
	//copy(*receiver.source, item)
	lst.source = &source
	lst.IList.source = &source
	lst.IList.Collection.source = &source
	lst.IList.Collection.lock = &lock
	lst.Enumerable.source = &source
	lst.Enumerable.lock = &lock
	return lst
}

// NewListFromChan 创建集合，将chan中的数据填充到集合
func NewListFromChan[T any](c chan T) List[T] {
	var lst = List[T]{}
	lst.New()
	for len(c) > 0 {
		lst.Add(<-c)
	}
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

	return receiver.UnmarshalJSON(ba)
}

// IsNil 是否未初始化
func (receiver *List[T]) IsNil() bool {
	return receiver.source == nil
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver List[T]) MarshalJSON() ([]byte, error) {
	if receiver.IsNil() {
		return []byte("[]"), nil
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()
	return snc.Marshal(receiver.source)
}

// UnmarshalJSON to deserialize []byte
func (receiver *List[T]) UnmarshalJSON(b []byte) error {
	if receiver.IsNil() {
		receiver.New()
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()
	return snc.Unmarshal(b, receiver.source)
}

// GormDataType gorm common data type
func (receiver *List[T]) GormDataType() string {
	return "json"
}

// Copy 克隆出新的集合
func (receiver List[T]) Copy() List[T] {
	arr := *receiver.source
	return NewList(arr...)
}
