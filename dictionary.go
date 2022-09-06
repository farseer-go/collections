package collections

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/farseer-go/fs/parse"
)

// Dictionary 字典
type Dictionary[TKey comparable, TValue any] struct {
	// source array
	source map[TKey]TValue
}

// NewDictionary 创建一个字典
func NewDictionary[TKey comparable, TValue any]() Dictionary[TKey, TValue] {
	return Dictionary[TKey, TValue]{
		source: make(map[TKey]TValue),
	}
}

// NewDictionaryFromMap 创建一个字典
func NewDictionaryFromMap[TKey comparable, TValue any](source map[TKey]TValue) Dictionary[TKey, TValue] {
	return Dictionary[TKey, TValue]{
		source: source,
	}
}

// Values 获取字典的value
func (receiver Dictionary[TKey, TValue]) Values() List[TValue] {
	lst := NewList[TValue]()
	for _, v := range receiver.source {
		lst.Add(v)
	}
	return lst
}

// Keys 获取字典的Keys
func (receiver Dictionary[TKey, TValue]) Keys() List[TKey] {
	lst := NewList[TKey]()
	for k := range receiver.source {
		lst.Add(k)
	}
	return lst
}

// Count 获取字典数量
func (receiver Dictionary[TKey, TValue]) Count() int {
	return len(receiver.source)
}

// AddMap 添加元素
func (receiver Dictionary[TKey, TValue]) AddMap(source map[TKey]TValue) {
	for key, value := range source {
		receiver.Add(key, value)
	}
}

// Add 添加元素
func (receiver Dictionary[TKey, TValue]) Add(key TKey, value TValue) {
	receiver.source[key] = value
}

// Clear 清除元素
func (receiver *Dictionary[TKey, TValue]) Clear() {
	receiver.source = make(map[TKey]TValue)
}

// Remove 移除元素
func (receiver Dictionary[TKey, TValue]) Remove(key TKey) {
	delete(receiver.source, key)
}

// ContainsKey 是否存在KEY
func (receiver Dictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	_, exists := receiver.source[key]
	return exists
}

// ContainsValue 是否存在指定的Value
func (receiver Dictionary[TKey, TValue]) ContainsValue(value TValue) bool {
	for _, itemValue := range receiver.source {
		if parse.IsEqual(value, itemValue) {
			return true
		}
	}
	return false
}

// GetValue 获取指定KEY的Value
func (receiver Dictionary[TKey, TValue]) GetValue(key TKey) TValue {
	return receiver.source[key]
}

// ToMap 转成map对象
func (receiver Dictionary[TKey, TValue]) ToMap() map[TKey]TValue {
	return receiver.source
}

// Value return json value, implement driver.Valuer interface
func (receiver Dictionary[TKey, TValue]) Value() (driver.Value, error) {
	if receiver.source == nil {
		return nil, nil
	}
	ba, err := receiver.MarshalJSON()
	return string(ba), err
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

	t := map[TKey]TValue{}
	err := json.Unmarshal(ba, &t)
	receiver.source = t
	return err
}

// MarshalJSON to output non base64 encoded []byte
func (receiver Dictionary[TKey, TValue]) MarshalJSON() ([]byte, error) {
	if receiver.source == nil {
		return []byte("null"), nil
	}
	return json.Marshal(receiver.source)
}

// UnmarshalJSON to deserialize []byte
func (receiver *Dictionary[TKey, TValue]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &receiver.source)
}

// GormDataType gorm common data type
func (receiver Dictionary[TKey, TValue]) GormDataType() string {
	return "jsonmap"
}

// IsNil 是否未初始化
func (receiver Dictionary[TKey, TValue]) IsNil() bool {
	return receiver.source == nil
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
