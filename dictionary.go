package collections

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// Dictionary 字典
type Dictionary[TKey comparable, TValue any] struct {
	ReadonlyDictionary[TKey, TValue]
}

// NewDictionary 创建一个字典
func NewDictionary[TKey comparable, TValue any]() Dictionary[TKey, TValue] {
	return Dictionary[TKey, TValue]{
		ReadonlyDictionary: NewReadonlyDictionary[TKey, TValue](),
	}
}

// NewDictionaryFromMap 创建一个字典
func NewDictionaryFromMap[TKey comparable, TValue any](source map[TKey]TValue) Dictionary[TKey, TValue] {
	return Dictionary[TKey, TValue]{
		ReadonlyDictionary: NewReadonlyDictionaryFromMap[TKey, TValue](source),
	}
}

// AddMap 添加元素
func (receiver *Dictionary[TKey, TValue]) AddMap(source map[TKey]TValue) {
	for key, value := range source {
		receiver.Add(key, value)
	}
}

// Add 添加元素
func (receiver *Dictionary[TKey, TValue]) Add(key TKey, value TValue) {
	receiver.source.Store(key, value)
}

// Update 更新元素
func (receiver *Dictionary[TKey, TValue]) Update(key TKey, f func(value *TValue)) {
	if v, exists := receiver.source.Load(key); exists {
		v2 := v.(TValue)
		f(&v2)
		receiver.source.Store(key, v2)
	}
}

// Clear 清除元素
func (receiver *Dictionary[TKey, TValue]) Clear() {
	receiver.source = &sync.Map{}
}

// Remove 移除元素
func (receiver *Dictionary[TKey, TValue]) Remove(key TKey) {
	receiver.source.Delete(key)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (receiver *Dictionary[TKey, TValue]) Scan(val any) error {
	if val == nil {
		*receiver = NewDictionary[TKey, TValue]()
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

// UnmarshalJSON to deserialize []byte
func (receiver *Dictionary[TKey, TValue]) UnmarshalJSON(ba []byte) error {
	t := map[TKey]TValue{}
	err := json.Unmarshal(ba, &t)
	*receiver = NewDictionaryFromMap(t)
	return err
}

// ToReadonlyDictionary 转成ReadonlyDictionary对象
func (receiver *Dictionary[TKey, TValue]) ToReadonlyDictionary() ReadonlyDictionary[TKey, TValue] {
	return receiver.ReadonlyDictionary
}

// New 初始化（用于反映时使用）
func (receiver *Dictionary[TKey, TValue]) New() {
	if receiver.source == nil {
		receiver.source = &sync.Map{}
	}
}

// Foreach for range操作
func (receiver *Dictionary[TKey, TValue]) Foreach(itemFn func(TKey, TValue)) {
	if receiver.source == nil {
		return
	}
	receiver.source.Range(func(key, value any) bool {
		itemFn(key.(TKey), value.(TValue))
		return true
	})
}
