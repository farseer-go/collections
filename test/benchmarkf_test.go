package test

import (
	"testing"

	"github.com/farseer-go/collections"
)

// BenchmarkNewListCap-10           1976743               597.2 ns/op          8240 B/op          3 allocs/op
func BenchmarkNewListCap(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		collections.NewListCap[int](1000)
	}
}
