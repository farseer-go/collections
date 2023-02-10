package collections

import (
	"encoding/json"
	"errors"
	"fmt"
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
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	receiver.source[key] = value
}

// Clear 清除元素
func (receiver *Dictionary[TKey, TValue]) Clear() {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	receiver.source = make(map[TKey]TValue)
}

// Remove 移除元素
func (receiver *Dictionary[TKey, TValue]) Remove(key TKey) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	delete(receiver.source, key)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (receiver *Dictionary[TKey, TValue]) Scan(val any) error {
	if receiver.lock != nil {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()
	}
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
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return receiver.ReadonlyDictionary
}
