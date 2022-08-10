package collections

import "reflect"

type list[T any] struct {
	source *[]T
	collection[T]
}

// Index 获取第index索引位置的元素
func (receiver *list[T]) Index(index int) T {
	return (*receiver.source)[index]
}

// IndexOf 元素在集合的索引位置
func (receiver *list[T]) IndexOf(item T) int {
	itemValue := reflect.ValueOf(item)
	for index, t := range *receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return index
		}
	}
	return -1
}

// Insert 向第index索引位置插入元素
func (receiver *list[T]) Insert(index int, item T) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(*receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		*receiver.source = append([]T{item}, *receiver.source...)
	} else {
		*receiver.source = append((*receiver.source)[:index], append([]T{item}, (*receiver.source)[index:]...)...)
	}
}

// RemoveAt 移除指定索引的元素
func (receiver *list[T]) RemoveAt(index int) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(*receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		*receiver.source = (*receiver.source)[1:]
	} else {
		*receiver.source = append((*receiver.source)[:index], (*receiver.source)[index+1:]...)
	}
}
