package collections

import (
	"sync"

	"github.com/farseer-go/fs/parse"
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

// AddIfNotExists 如果元素不存在，则添加元素
func (receiver *Collection[T]) AddIfNotExists(item T) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	for _, t := range *receiver.source {
		if parse.IsEqual(t, item) {
			return
		}
	}
	*receiver.source = append(*receiver.source, item)
}

// AddRange 添加元素
func (receiver *Collection[T]) AddRange(lst Enumerable[T]) {
	if lst.Count() > 0 {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()

		*receiver.source = append(*receiver.source, lst.ToArray()...)
	}
}

// AddList 添加元素
func (receiver *Collection[T]) AddList(lst List[T]) {
	if lst.Count() > 0 {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()

		*receiver.source = append(*receiver.source, lst.ToArray()...)
	}
}

// AddArray 添加元素
func (receiver *Collection[T]) AddArray(arr []T) {
	if len(arr) > 0 {
		receiver.lock.Lock()
		defer receiver.lock.Unlock()

		*receiver.source = append(*receiver.source, arr...)
	}
}

// Clear 清空集合
func (receiver *Collection[T]) Clear() {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	//*receiver.source = []T{}
	*receiver.source = (*receiver.source)[:0]
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
