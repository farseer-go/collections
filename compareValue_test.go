package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLeftGreaterThanRight(t *testing.T) {
	assert.False(t, CompareLeftGreaterThanRight(int(1), int(5)))
	assert.True(t, CompareLeftGreaterThanRight(int(5), int(1)))

	assert.False(t, CompareLeftGreaterThanRight(int8(1), int8(5)))
	assert.True(t, CompareLeftGreaterThanRight(int8(5), int8(1)))

	assert.False(t, CompareLeftGreaterThanRight(int16(1), int16(5)))
	assert.True(t, CompareLeftGreaterThanRight(int16(5), int16(1)))

	assert.False(t, CompareLeftGreaterThanRight(int32(1), int32(5)))
	assert.True(t, CompareLeftGreaterThanRight(int32(5), int32(1)))

	assert.False(t, CompareLeftGreaterThanRight(int64(1), int64(5)))
	assert.True(t, CompareLeftGreaterThanRight(int64(5), int64(1)))

	assert.False(t, CompareLeftGreaterThanRight(uint(1), uint(5)))
	assert.True(t, CompareLeftGreaterThanRight(uint(5), uint(1)))

	assert.False(t, CompareLeftGreaterThanRight(uint8(1), uint8(5)))
	assert.True(t, CompareLeftGreaterThanRight(uint8(5), uint8(1)))

	assert.False(t, CompareLeftGreaterThanRight(uint16(1), uint16(5)))
	assert.True(t, CompareLeftGreaterThanRight(uint16(5), uint16(1)))

	assert.False(t, CompareLeftGreaterThanRight(uint32(1), uint32(5)))
	assert.True(t, CompareLeftGreaterThanRight(uint32(5), uint32(1)))

	assert.False(t, CompareLeftGreaterThanRight(uint64(1), uint64(5)))
	assert.True(t, CompareLeftGreaterThanRight(uint64(5), uint64(1)))

	assert.False(t, CompareLeftGreaterThanRight(float32(1), float32(5)))
	assert.True(t, CompareLeftGreaterThanRight(float32(5), float32(1)))

	assert.False(t, CompareLeftGreaterThanRight(float64(1), float64(5)))
	assert.True(t, CompareLeftGreaterThanRight(float64(5), float64(1)))
}
