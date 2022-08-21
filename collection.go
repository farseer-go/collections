package collections

import "reflect"

type Collection[T any] struct {
	source *[]T
}

// Count 集合大小
func (receiver *Collection[T]) Count() int {
	return len(*receiver.source)
}

// Add 添加元素
func (receiver *Collection[T]) Add(item ...T) {
	if item == nil || len(item) == 0 {
		return
	}
	*receiver.source = append(*receiver.source, item...)
}

// Clear 清空集合
func (receiver *Collection[T]) Clear() {
	*receiver.source = []T{}
}

// Remove 移除元素
func (receiver *Collection[T]) Remove(item T) {
	itemValue := reflect.ValueOf(item)
	for i := 0; i < len(*receiver.source); i++ {
		if reflect.ValueOf((*receiver.source)[i]) == itemValue {
			*receiver.source = append((*receiver.source)[:i], (*receiver.source)[i+1:]...)
			i--
		}
	}
}

// RemoveAll 移除条件=true的元素
func (receiver *Collection[T]) RemoveAll(fn func(item T) bool) {
	for i := 0; i < len(*receiver.source); i++ {
		if fn((*receiver.source)[i]) {
			*receiver.source = append((*receiver.source)[:i], (*receiver.source)[i+1:]...)
			i--
		}
	}
}
