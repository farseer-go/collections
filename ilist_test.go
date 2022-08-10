package collections

import (
	"testing"
)

func Test_list_Index(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	if lst.Index(4) != 5 {
		t.Error()
	}
}

func Test_list_IndexOf(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	if lst.IndexOf(4) != 3 {
		t.Error()
	}
}

func Test_list_Insert(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	lst.Insert(1, 8)
	if lst.Count() != 4 {
		t.Error()
	}
	if lst.Index(0) != 1 || lst.Index(1) != 8 || lst.Index(2) != 2 || lst.Index(3) != 3 {
		t.Error()
	}
}

func Test_list_RemoveAt(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst.RemoveAt(3)
	if lst.Count() != 5 {
		t.Error()
	}
	if lst.Contains(4) {
		t.Error()
	}
	if lst.Index(0) != 1 || lst.Index(1) != 2 || lst.Index(2) != 3 || lst.Index(3) != 5 || lst.Index(4) != 6 {
		t.Error()
	}
}
