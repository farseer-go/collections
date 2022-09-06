package collections

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	source := *receiver.source
	return Enumerable[T]{
		source: &source,
	}
}

// New 返回enumerable类型
func (receiver *List[T]) New() {
	if receiver.source == nil {
		source := &[]T{}
		receiver.source = source
		receiver.IList.source = source
		receiver.IList.Collection.source = source
		receiver.Enumerable.source = source
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

	var t []T
	err := json.Unmarshal(ba, &t)
	*receiver = NewList[T](t...)
	return err
}

// IsNil 是否未初始化
func (receiver *List[T]) IsNil() bool {
	return receiver.source == nil
}
