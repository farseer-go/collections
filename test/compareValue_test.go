package test

import (
	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLeftGreaterThanRight(t *testing.T) {
	assert.False(t, collections.CompareLeftGreaterThanRight(int(1), int(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(int(5), int(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(int8(1), int8(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(int8(5), int8(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(int16(1), int16(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(int16(5), int16(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(int32(1), int32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(int32(5), int32(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(int64(1), int64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(int64(5), int64(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(uint(1), uint(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(uint(5), uint(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(uint8(1), uint8(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(uint8(5), uint8(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(uint16(1), uint16(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(uint16(5), uint16(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(uint32(1), uint32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(uint32(5), uint32(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(uint64(1), uint64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(uint64(5), uint64(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(float32(1), float32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(float32(5), float32(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight(float64(1), float64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(float64(5), float64(1)))

	assert.False(t, collections.CompareLeftGreaterThanRight("a", "b"))
	assert.False(t, collections.CompareLeftGreaterThanRight("aa", "ab"))
	assert.True(t, collections.CompareLeftGreaterThanRight("", "ab"))
	assert.False(t, collections.CompareLeftGreaterThanRight("aba", "ab"))

	assert.Panics(t, func() {
		collections.CompareLeftGreaterThanRight([]int{}, []int{})
	})
}
