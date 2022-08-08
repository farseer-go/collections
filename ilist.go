package collections

import "reflect"

type IList struct {
	source []any
}

func NewIList(source ...any) IList {
	return IList{
		source: source,
	}
}

// Add 添加元素
func (receiver *IList) Add(item ...any) {
	receiver.source = append(receiver.source, item...)
}

// Count 集合大小
func (receiver *IList) Count() int {
	return len(receiver.source)
}

// ToArray 转成数组
func (receiver *IList) ToArray() []any {
	return receiver.source
}

// Foreach 遍历集合
func (receiver *IList) Foreach(fn func(int, *any)) {
	for i, a := range receiver.source {
		fn(i, &a)
	}
}

// IsEmpty 集合是为空的
func (receiver *IList) IsEmpty() bool {
	return receiver.source == nil || len(receiver.source) == 0
}

// Index 获取第index索引位置的元素
func (receiver *IList) Index(index int) any {
	return receiver.source[index]
}

// Contains 是否包含元素
func (receiver IList) Contains(item any) bool {
	itemValue := reflect.ValueOf(item)
	for _, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return true
		}
	}
	return false
}

// IndexOf 元素在集合的索引位置
func (receiver IList) IndexOf(item any) int {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return index
		}
	}
	return -1
}

// Remove 移除元素
func (receiver IList) Remove(item any) {
	itemValue := reflect.ValueOf(item)
	for index, t := range receiver.source {
		if reflect.ValueOf(t) == itemValue {
			receiver.RemoveAt(index)
		}
	}
}

// RemoveAt 移除指定索引的元素
func (receiver *IList) RemoveAt(index int) {
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
func (receiver *IList) Insert(index int, item any) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		receiver.source = append([]any{item}, receiver.source...)
	} else {
		receiver.source = append(receiver.source[:index], append([]any{item}, receiver.source[index:]...)...)
	}
}

// Clear 清空集合
func (receiver *IList) Clear() {
	receiver.source = []any{}
}
