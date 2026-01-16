package test

import (
	"testing"

	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/fastReflect"
	"github.com/stretchr/testify/assert"
)

func TestCompareLeftGreaterThanRight(t *testing.T) {
	pointerMeta := fastReflect.PointerOf(int(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, int(1), int(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, int(5), int(1)))

	pointerMeta = fastReflect.PointerOf(int8(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, int8(1), int8(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, int8(5), int8(1)))

	pointerMeta = fastReflect.PointerOf(int16(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, int16(1), int16(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, int16(5), int16(1)))

	pointerMeta = fastReflect.PointerOf(int32(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, int32(1), int32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, int32(5), int32(1)))

	pointerMeta = fastReflect.PointerOf(int64(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, int64(1), int64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, int64(5), int64(1)))

	pointerMeta = fastReflect.PointerOf(uint(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint(1), uint(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint(5), uint(1)))

	pointerMeta = fastReflect.PointerOf(uint8(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint8(1), uint8(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint8(5), uint8(1)))

	pointerMeta = fastReflect.PointerOf(uint16(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint16(1), uint16(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint16(5), uint16(1)))

	pointerMeta = fastReflect.PointerOf(uint32(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint32(1), uint32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint32(5), uint32(1)))

	pointerMeta = fastReflect.PointerOf(uint64(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint64(1), uint64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, uint64(5), uint64(1)))

	pointerMeta = fastReflect.PointerOf(float32(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, float32(1), float32(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, float32(5), float32(1)))

	pointerMeta = fastReflect.PointerOf(float64(1))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, float64(1), float64(5)))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, float64(5), float64(1)))

	pointerMeta = fastReflect.PointerOf("")
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, "a", "b"))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, "aa", "ab"))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, "", "ab"))
	assert.True(t, collections.CompareLeftGreaterThanRight(pointerMeta, "aba", "ab"))
	assert.False(t, collections.CompareLeftGreaterThanRight(pointerMeta, "aa", "aaa"))

	assert.Panics(t, func() {
		pointerMeta = fastReflect.PointerOf([]int{})
		collections.CompareLeftGreaterThanRight(pointerMeta, []int{}, []int{})
	})
}
