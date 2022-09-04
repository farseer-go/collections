package collections

import (
	"encoding/json"
	"github.com/farseer-go/fs/parse"
)

type Collection[T any] struct {
	source *[]T
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
	for i := 0; i < len(*receiver.source); i++ {
		if parse.IsEqual((*receiver.source)[i], item) {
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

// MarshalJSON to output non base64 encoded []byte
func (receiver *Collection[T]) MarshalJSON() ([]byte, error) {
	if receiver.source == nil {
		return []byte("null"), nil
	}
	return json.Marshal(receiver.source)
}

// UnmarshalJSON to deserialize []byte
func (receiver *Collection[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &receiver.source)
}
