package test

import (
	"math/rand"
	"testing"

	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
)

func Test_enumerable_OrderBy(t *testing.T) {
	lst := collections.NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderBy(func(item int) any {
		return item
	}).ToArray()

	assert.Equal(t, item[0], 1)
	assert.Equal(t, item[1], 2)
	assert.Equal(t, item[2], 3)
	assert.Equal(t, item[3], 4)
	assert.Equal(t, item[4], 5)
	assert.Equal(t, item[5], 6)
	assert.Equal(t, item[6], 7)
	assert.Equal(t, item[7], 8)
}

func Test_enumerable_OrderByThen(t *testing.T) {
	lst := collections.NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByThen(func(leftItem, rightItem int) bool {
		return leftItem > rightItem
	}).ToArray()

	assert.Equal(t, item[0], 1)
	assert.Equal(t, item[1], 2)
	assert.Equal(t, item[2], 3)
	assert.Equal(t, item[3], 4)
	assert.Equal(t, item[4], 5)
	assert.Equal(t, item[5], 6)
	assert.Equal(t, item[6], 7)
	assert.Equal(t, item[7], 8)
}

func Test_enumerable_OrderByItem(t *testing.T) {
	lst := collections.NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByItem().ToArray()

	assert.Equal(t, item[0], 1)
	assert.Equal(t, item[1], 2)
	assert.Equal(t, item[2], 3)
	assert.Equal(t, item[3], 4)
	assert.Equal(t, item[4], 5)
	assert.Equal(t, item[5], 6)
	assert.Equal(t, item[6], 7)
	assert.Equal(t, item[7], 8)
}

func Test_enumerable_OrderByDescending(t *testing.T) {
	lst := collections.NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescending(func(item int) any {
		return item
	}).ToArray()

	assert.Equal(t, item[0], 8)
	assert.Equal(t, item[1], 7)
	assert.Equal(t, item[2], 6)
	assert.Equal(t, item[3], 5)
	assert.Equal(t, item[4], 4)
	assert.Equal(t, item[5], 3)
	assert.Equal(t, item[6], 2)
	assert.Equal(t, item[7], 1)
}

func Test_enumerable_OrderByDescendingItem(t *testing.T) {
	lst := collections.NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescendingItem().ToArray()

	assert.Equal(t, item[0], 8)
	assert.Equal(t, item[1], 7)
	assert.Equal(t, item[2], 6)
	assert.Equal(t, item[3], 5)
	assert.Equal(t, item[4], 4)
	assert.Equal(t, item[5], 3)
	assert.Equal(t, item[6], 2)
	assert.Equal(t, item[7], 1)
}

type family struct {
	Name string
	Age  int
}

// BenchmarkOrderBy-10            2         659323792 ns/op          245808 B/op          3 allocs/op
// BenchmarkOrderBy-10          334           3615433 ns/op          245905 B/op          6 allocs/op
func BenchmarkOrderBy(b *testing.B) {
	lst := collections.NewList[family]()
	for i := 0; i < 10000; i++ {
		lst.Add(family{"", rand.Intn(100)})
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lst.OrderBy(func(item family) any {
			return item.Age
		})
	}
}
