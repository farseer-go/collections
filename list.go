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

func (receiver *List[T]) Add(item ...T) {
	receiver.source = append(receiver.source, item...)
}

func (receiver *List[T]) Count() int {
	return len(receiver.source)
}

func (receiver *List[T]) ToArray() []T {
	return receiver.source
}

func (receiver *List[T]) IsEmpty() bool {
	return receiver.source == nil || len(receiver.source) == 0
}

func (receiver *List[T]) Index(index int) T {
	return receiver.source[index]
}

func (receiver List[T]) Contains(item T) bool {
	itemValue := reflect.ValueOf(item)
	for _, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return true
		}
	}
	return false
}

func (receiver List[T]) IndexOf(item T) int {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return index
		}
	}
	return -1
}

func (receiver List[T]) Remove(item T) {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			receiver.RemoveAt(index)
		}
	}
}

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

func (receiver *List[T]) Clear() {
	receiver.source = []T{}
}
