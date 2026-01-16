package test

import (
	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAny(t *testing.T) {
	lst := collections.NewListAny(1, 2, 3)
	assert.Equal(t, 3, lst.Count())
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
	assert.Equal(t, 3, lst.Index(2))
}
