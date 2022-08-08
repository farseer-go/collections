package collections

import "reflect"

// List 集合
type List[T any] struct {
	source []T
}

func NewList[T any](source ...T) List[T] {
	return List[T]{
		source: source,
	}
}

// Add 添加元素
func (receiver *List[T]) Add(item ...T) {
	receiver.source = append(receiver.source, item...)
}

// Count 集合大小
func (receiver *List[T]) Count() int {
	return len(receiver.source)
}

// ToArray 转成数组
func (receiver *List[T]) ToArray() []T {
	return receiver.source
}

// IsEmpty 集合是为空的
func (receiver *List[T]) IsEmpty() bool {
	return receiver.source == nil || len(receiver.source) == 0
}

// Index 获取第index索引位置的元素
func (receiver *List[T]) Index(index int) T {
	return receiver.source[index]
}

// Contains 是否包含元素
func (receiver List[T]) Contains(item T) bool {
	itemValue := reflect.ValueOf(item)
	for _, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return true
		}
	}
	return false
}

// IndexOf 元素在集合的索引位置
func (receiver List[T]) IndexOf(item T) int {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return index
		}
	}
	return -1
}

// Remove 移除元素
func (receiver List[T]) Remove(item T) {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			receiver.RemoveAt(index)
		}
	}
}

// RemoveAt 移除指定索引的元素
func (receiver *List[T]) RemoveAt(index int) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		receiver.source = receiver.source[1:]
	} else {
		receiver.source = append(receiver.source[:index], receiver.source[index+1:]...)
	}
}

// Insert 向第index索引位置插入元素
func (receiver *List[T]) Insert(index int, item T) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		receiver.source = append([]T{item}, receiver.source...)
	} else {
		receiver.source = append(receiver.source[:index], append([]T{item}, receiver.source[index:]...)...)
	}
}

// Clear 清空集合
func (receiver *List[T]) Clear() {
	receiver.source = []T{}
}
