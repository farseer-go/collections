package collections

import (
	"github.com/farseer-go/fs/parse"
	"sync"
)

type Collection[T any] struct {
	source *[]T
	lock   *sync.RWMutex
}

// Add 添加元素
func (receiver *Collection[T]) Add(item ...T) {
	if len(item) > 0 {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()

		*receiver.source = append(*receiver.source, item...)
	}
}

// AddRange 添加元素
func (receiver *Collection[T]) AddRange(lst Enumerable[T]) {
	if lst.Count() > 0 {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()

		*receiver.source = append(*receiver.source, lst.ToArray()...)
	}
}

// Clear 清空集合
func (receiver *Collection[T]) Clear() {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	*receiver.source = []T{}
}

// Remove 移除元素
func (receiver *Collection[T]) Remove(item T) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	for i := 0; i < len(*receiver.source); i++ {
		if parse.IsEqual((*receiver.source)[i], item) {
			*receiver.source = append((*receiver.source)[:i], (*receiver.source)[i+1:]...)
			i--
		}
	}
}

// RemoveAll 移除条件=true的元素
func (receiver *Collection[T]) RemoveAll(fn func(item T) bool) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	for i := 0; i < len(*receiver.source); i++ {
		if fn((*receiver.source)[i]) {
			*receiver.source = append((*receiver.source)[:i], (*receiver.source)[i+1:]...)
			i--
		}
	}
}
