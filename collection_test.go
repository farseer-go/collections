package collections

import "testing"

func Test_collection_Count(t *testing.T) {
	lst := NewList[int]()
	lst.Add(1, 2, 3) // lst = 1, 2, 3
	lst.Insert(1, 8) // lst = 1, 8, 2, 3
	if lst.Count() != 4 {
		t.Error()
	}
}

func Test_collection_Add(t *testing.T) {
	lst := NewList[int](1, 2)
	lst.Add(3)
	if lst.Count() != 3 {
		t.Error()
	}
	if lst.Index(2) != 3 {
		t.Error()
	}
}

func Test_collection_Clear(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	if lst.Count() != 3 {
		t.Error()
	}
	lst.Clear()
	if lst.Count() != 0 {
		t.Error()
	}
}

func Test_collection_RemoveAll(t *testing.T) {
	lst := NewList[int](1, 2, 3, 6)
	lst.RemoveAll(func(item int) bool {
		return item >= 3
	})
	if lst.Count() != 2 {
		t.Error()
	}
	if lst.Index(0) != 1 && lst.Index(1) != 2 {
		t.Error()
	}
}
