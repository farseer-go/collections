package collections

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/farseer-go/fs/parse"
	"sync"
)

// ReadonlyDictionary 只读字典
type ReadonlyDictionary[TKey comparable, TValue any] struct {
	// source array
	source map[TKey]TValue
	lock   *sync.RWMutex
}

// NewReadonlyDictionary 创建一个字典
func NewReadonlyDictionary[TKey comparable, TValue any]() ReadonlyDictionary[TKey, TValue] {
	return ReadonlyDictionary[TKey, TValue]{
		source: make(map[TKey]TValue),
		lock:   &sync.RWMutex{},
	}
}

// NewReadonlyDictionaryFromMap 创建一个字典
func NewReadonlyDictionaryFromMap[TKey comparable, TValue any](source map[TKey]TValue) ReadonlyDictionary[TKey, TValue] {
	return ReadonlyDictionary[TKey, TValue]{
		source: source,
		lock:   &sync.RWMutex{},
	}
}

// Values 获取字典的value
func (receiver ReadonlyDictionary[TKey, TValue]) Values() List[TValue] {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := NewList[TValue]()
	for _, v := range receiver.source {
		lst.Add(v)
	}
	return lst
}

// Keys 获取字典的Keys
func (receiver ReadonlyDictionary[TKey, TValue]) Keys() List[TKey] {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := NewList[TKey]()
	for k := range receiver.source {
		lst.Add(k)
	}
	return lst
}

// Count 获取字典数量
func (receiver ReadonlyDictionary[TKey, TValue]) Count() int {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return len(receiver.source)
}

// ContainsKey 是否存在KEY
func (receiver ReadonlyDictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	_, exists := receiver.source[key]
	return exists
}

// ContainsValue 是否存在指定的Value
func (receiver ReadonlyDictionary[TKey, TValue]) ContainsValue(value TValue) bool {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, itemValue := range receiver.source {
		if parse.IsEqual(value, itemValue) {
			return true
		}
	}
	return false
}

// GetValue 获取指定KEY的Value
func (receiver ReadonlyDictionary[TKey, TValue]) GetValue(key TKey) TValue {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return receiver.source[key]
}

// ToMap 转成map对象
func (receiver ReadonlyDictionary[TKey, TValue]) ToMap() map[TKey]TValue {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return receiver.source
}

// Value return json value, implement driver.Valuer interface
func (receiver ReadonlyDictionary[TKey, TValue]) Value() (driver.Value, error) {
	if receiver.source == nil {
		return nil, nil
	}
	ba, err := receiver.MarshalJSON()
	return string(ba), err
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver ReadonlyDictionary[TKey, TValue]) MarshalJSON() ([]byte, error) {
	if receiver.source == nil || receiver.lock == nil {
		return []byte("{}"), nil
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return json.Marshal(receiver.source)
}

// GormDataType gorm common data type
func (receiver ReadonlyDictionary[TKey, TValue]) GormDataType() string {
	return "jsonmap"
}

// IsNil 是否未初始化
func (receiver ReadonlyDictionary[TKey, TValue]) IsNil() bool {
	if receiver.lock == nil {
		return true
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return receiver.source == nil
}

// ToDictionary 返回只写字典
func (receiver ReadonlyDictionary[TKey, TValue]) ToDictionary() Dictionary[TKey, TValue] {
	return NewDictionaryFromMap(receiver.source)
}

//// GormDBDataType gorm db data type
//func (Dictionary[TKey, TValue]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
//	switch db.Dialector.Name() {
//	case "sqlite":
//		return "JSON"
//	case "mysql":
//		return "JSON"
//	case "postgres":
//		return "JSONB"
//	case "sqlserver":
//		return "NVARCHAR(MAX)"
//	}
//	return ""
//}
