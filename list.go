package collections

import (
	"reflect"
	"strings"
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
	return receiver.Enumerable
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

// ReflectNew 动态创建一个新的List
func ReflectNew(lstType reflect.Type) reflect.Value {
	lstValue := reflect.New(lstType)
	lstValue.MethodByName("New").Call(nil)
	return lstValue
}

// ReflectAdd 动态添加元素
func ReflectAdd(lstValue *reflect.Value, item any) {
	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.MethodByName("Add").CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.MethodByName("Add").Call([]reflect.Value{itemValue})
	}
}

// ReflectItemArrayType 获取List的原始数组
func ReflectItemArrayType(lstType reflect.Type) reflect.Type {
	sourceField, _ := lstType.FieldByName("source")
	return sourceField.Type.Elem()
}

// ReflectItemType 获取List的元素Type
func ReflectItemType(lstType reflect.Type) reflect.Type {
	sourceField, _ := lstType.FieldByName("source")
	return sourceField.Type.Elem().Elem()
}

// ReflectIsList 判断类型是否为List
func ReflectIsList(lstType reflect.Type) bool {
	if lstType.Kind() == reflect.Ptr {
		lstType = lstType.Elem()
	}
	return strings.HasPrefix(lstType.String(), "collections.List[")
}

func ReflectToArray(lstValue reflect.Value) []any {
	arrValue := lstValue.MethodByName("ToArray").Call(nil)[0]
	var items []any
	for i := 0; i < arrValue.Len(); i++ {
		item := arrValue.Index(i).Interface()
		items = append(items, item)
	}
	return items
}
