package collections

import (
	"testing"
)

func TestList_Add(t *testing.T) {
	list := NewList[int]()
	list.Add(1)
	if list.Count() != 1 {
		t.Error()
	}
}

func TestList_ToArray(t *testing.T) {
	list := NewList[int]()
	list.Add(1)
	arr := list.ToArray()
	if arr[0] != 1 || len(arr) != 1 {
		t.Error()
	}
}

func TestList_IsEmpty(t *testing.T) {
	list := NewList[int]()
	if !list.IsEmpty() {
		t.Error()
	}
	list.Add(1)
	if list.IsEmpty() {
		t.Error()
	}
}

func TestList_Index(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3, 4, 5)
	if list.Index(3) != 4 {
		t.Error()
	}
}

func TestList_Contains(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3, 4, 5)
	if !list.Contains(4) {
		t.Error()
	}
	if list.Contains(0) {
		t.Error()
	}
}

func TestList_IndexOf(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3, 4, 5)
	if list.IndexOf(3) != 2 {
		t.Error()
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3, 4, 3, 5)
	list.Remove(3)
	if list.Count() != 4 {
		t.Error()
	}
	if list.Contains(3) {
		t.Error()
	}
	if list.Index(0) != 1 || list.Index(1) != 2 || list.Index(2) != 4 || list.Index(3) != 5 {
		t.Error()
	}
}

func TestList_RemoveAt(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3, 4, 5)
	list.RemoveAt(3)
	if list.Count() != 4 {
		t.Error()
	}
	if list.Contains(4) {
		t.Error()
	}
	if list.Index(0) != 1 || list.Index(1) != 2 || list.Index(2) != 3 || list.Index(3) != 5 {
		t.Error()
	}
}

func TestList_Insert(t *testing.T) {
	list := NewList[int]()
	list.Add(1, 2, 3)
	list.Insert(0, 8)
	if list.Count() != 4 {
		t.Error()
	}
	if list.Index(0) != 8 || list.Index(1) != 1 || list.Index(2) != 2 || list.Index(3) != 3 {
		t.Error()
	}
}

func TestList_Clear(t *testing.T) {

}
