package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, int(3), Addition(int(1), int(2)))
	assert.Equal(t, int8(3), Addition(int8(1), int8(2)))
	assert.Equal(t, int16(3), Addition(int16(1), int16(2)))
	assert.Equal(t, int32(3), Addition(int32(1), int32(2)))
	assert.Equal(t, int64(3), Addition(int64(1), int64(2)))

	assert.Equal(t, uint(3), Addition(uint(1), uint(2)))
	assert.Equal(t, uint8(3), Addition(uint8(1), uint8(2)))
	assert.Equal(t, uint16(3), Addition(uint16(1), uint16(2)))
	assert.Equal(t, uint32(3), Addition(uint32(1), uint32(2)))
	assert.Equal(t, uint64(3), Addition(uint64(1), uint64(2)))

	assert.Equal(t, float32(3), Addition(float32(1), float32(2)))
	assert.Equal(t, float64(3), Addition(float64(1), float64(2)))

	assert.Panics(t, func() {
		Addition("1", "2")
	})
}
