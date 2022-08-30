package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_list_Index(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, 5, lst.Index(4))
	lst.Set(4, 9)
	assert.Equal(t, 9, lst.Index(4))
}

func Test_list_IndexOf(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, 3, lst.IndexOf(4))
	assert.Equal(t, -1, lst.IndexOf(7))
}

func Test_list_Insert(t *testing.T) {
	lst := NewList[int](1, 2, 3)

	assert.Panics(t, func() {
		lst.Insert(-1, 10)
	})

	assert.Panics(t, func() {
		lst.Insert(10, 55)
	})

	lst.Insert(1, 8)
	assert.Equal(t, 4, lst.Count())
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 8, lst.Index(1))
	assert.Equal(t, 2, lst.Index(2))
	assert.Equal(t, 3, lst.Index(3))

	lst.Insert(0, 9)
	assert.Equal(t, 5, lst.Count())
	assert.Equal(t, 9, lst.Index(0))
	assert.Equal(t, 1, lst.Index(1))
	assert.Equal(t, 8, lst.Index(2))
	assert.Equal(t, 2, lst.Index(3))
	assert.Equal(t, 3, lst.Index(4))

}

func Test_list_RemoveAt(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)

	assert.Panics(t, func() {
		lst.RemoveAt(-1)
	})

	assert.Panics(t, func() {
		lst.RemoveAt(10)
	})

	lst.RemoveAt(3)
	assert.Equal(t, 5, lst.Count())
	assert.False(t, lst.Contains(4))
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
	assert.Equal(t, 3, lst.Index(2))
	assert.Equal(t, 5, lst.Index(3))
	assert.Equal(t, 6, lst.Index(4))

	lst.RemoveAt(0)
	assert.Equal(t, 4, lst.Count())
	assert.False(t, lst.Contains(4))
	assert.False(t, lst.Contains(1))
	assert.Equal(t, 2, lst.Index(0))
	assert.Equal(t, 3, lst.Index(1))
	assert.Equal(t, 5, lst.Index(2))
	assert.Equal(t, 6, lst.Index(3))
}
