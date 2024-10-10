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
	source *sync.Map
}

// NewReadonlyDictionary 创建一个字典
func NewReadonlyDictionary[TKey comparable, TValue any]() ReadonlyDictionary[TKey, TValue] {
	return ReadonlyDictionary[TKey, TValue]{
		source: &sync.Map{},
	}
}

// NewReadonlyDictionaryFromMap 创建一个字典
func NewReadonlyDictionaryFromMap[TKey comparable, TValue any](source map[TKey]TValue) ReadonlyDictionary[TKey, TValue] {
	rdc := ReadonlyDictionary[TKey, TValue]{
		source: &sync.Map{},
	}
	for key, value := range source {
		rdc.source.Store(key, value)
	}
	return rdc
}

// Values 获取字典的value
func (receiver ReadonlyDictionary[TKey, TValue]) Values() List[TValue] {
	if receiver.source == nil {
		return NewList[TValue]()
	}

	lst := NewList[TValue]()
	receiver.source.Range(func(key, value any) bool {
		lst.Add(value.(TValue))
		return true
	})
	return lst
}

// Keys 获取字典的Keys
func (receiver ReadonlyDictionary[TKey, TValue]) Keys() List[TKey] {
	if receiver.source == nil {
		return NewList[TKey]()
	}

	lst := NewList[TKey]()
	receiver.source.Range(func(key, value any) bool {
		lst.Add(key.(TKey))
		return true
	})
	return lst
}

// Count 获取字典数量
func (receiver ReadonlyDictionary[TKey, TValue]) Count() int {
	if receiver.source == nil {
		return 0
	}
	count := 0
	receiver.source.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// ContainsKey 是否存在KEY
func (receiver ReadonlyDictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	if receiver.source == nil {
		return false
	}

	_, exists := receiver.source.Load(key)
	return exists
}

// ContainsValue 是否存在指定的Value
func (receiver ReadonlyDictionary[TKey, TValue]) ContainsValue(v TValue) bool {
	if receiver.source == nil {
		return false
	}

	result := false
	receiver.source.Range(func(key, value any) bool {
		result = parse.IsEqual(v, value.(TValue))
		return !result
	})
	return result
}

// GetValue 获取指定KEY的Value
func (receiver ReadonlyDictionary[TKey, TValue]) GetValue(key TKey) TValue {
	if receiver.source == nil {
		var val TValue
		return val
	}

	if v, exists := receiver.source.Load(key); exists {
		return v.(TValue)
	}
	
	var val TValue
	return val
}

// ToMap 转成map对象
func (receiver ReadonlyDictionary[TKey, TValue]) ToMap() map[TKey]TValue {
	if receiver.source == nil {
		return make(map[TKey]TValue)
	}

	m := make(map[TKey]TValue)
	receiver.source.Range(func(key, value any) bool {
		m[key.(TKey)] = value.(TValue)
		return true
	})
	return m
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
	if receiver.source == nil {
		return []byte("{}"), nil
	}

	return json.Marshal(receiver.ToMap())
}

// GormDataType gorm common data type
func (receiver ReadonlyDictionary[TKey, TValue]) GormDataType() string {
	return "JSON"
}

// IsNil 是否未初始化
func (receiver ReadonlyDictionary[TKey, TValue]) IsNil() bool {
	return receiver.source == nil
}

// ToDictionary 返回只写字典
func (receiver ReadonlyDictionary[TKey, TValue]) ToDictionary() Dictionary[TKey, TValue] {
	return Dictionary[TKey, TValue]{
		ReadonlyDictionary: receiver,
	}
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
