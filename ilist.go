package collections

import (
	"fmt"

	"github.com/farseer-go/fs/parse"
)

type IList[T any] struct {
	source *[]T
	Collection[T]
}

// Index 获取第index索引位置的元素
func (receiver *IList[T]) Index(index int) T {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return (*receiver.source)[index]
}

// Set 设置值
func (receiver *IList[T]) Set(index int, item T) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	(*receiver.source)[index] = item
}

// IndexOf 元素在集合的索引位置
func (receiver *IList[T]) IndexOf(item T) int {
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for index, t := range *receiver.source {
		if parse.IsEqual(t, item) {
			return index
		}
	}
	return -1
}

// Insert 向第index索引位置插入元素
func (receiver *IList[T]) Insert(index int, item T) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	if index < 0 {
		panic("list.Insert index值不能小于0")
	}
	if index > len(*receiver.source) {
		panic(fmt.Sprintf("list.Insert index值:%v 不能超出集合的长度:%v", index, len(*receiver.source)))
	}

	if index == 0 {
		*receiver.source = append([]T{item}, *receiver.source...)
	} else {
		*receiver.source = append((*receiver.source)[:index], append([]T{item}, (*receiver.source)[index:]...)...)
	}
}

// RemoveAt 移除指定索引的元素
func (receiver *IList[T]) RemoveAt(index int) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	if index < 0 {
		panic("list.RemoveAt index值不能小于0")
	}
	if index >= len(*receiver.source) {
		panic(fmt.Sprintf("list.RemoveAt index值:%v 不能超出集合的长度:%v", index, len(*receiver.source)))
	}

	if index == 0 {
		*receiver.source = (*receiver.source)[1:]
	} else {
		*receiver.source = append((*receiver.source)[:index], (*receiver.source)[index+1:]...)
	}
}
