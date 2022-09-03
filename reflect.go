package collections

import "reflect"

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

func ReflectToArray(lstValue reflect.Value) []any {
	arrValue := lstValue.MethodByName("ToArray").Call(nil)[0]
	var items []any
	for i := 0; i < arrValue.Len(); i++ {
		item := arrValue.Index(i).Interface()
		items = append(items, item)
	}
	return items
}
