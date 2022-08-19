package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_list_Index(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, lst.Index(4), 5)
	lst.Set(4, 9)
	assert.Equal(t, lst.Index(4), 9)
}

func Test_list_IndexOf(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, lst.IndexOf(4), 3)
}

func Test_list_Insert(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	lst.Insert(1, 8)
	assert.Equal(t, lst.Count(), 4)
	assert.Equal(t, lst.Index(0), 1)
	assert.Equal(t, lst.Index(1), 8)
	assert.Equal(t, lst.Index(2), 2)
	assert.Equal(t, lst.Index(3), 3)
}

func Test_list_RemoveAt(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst.RemoveAt(3)
	assert.Equal(t, lst.Count(), 5)
	assert.False(t, lst.Contains(4))
	assert.Equal(t, lst.Index(0), 1)
	assert.Equal(t, lst.Index(1), 2)
	assert.Equal(t, lst.Index(2), 3)
	assert.Equal(t, lst.Index(3), 5)
	assert.Equal(t, lst.Index(4), 6)
}
