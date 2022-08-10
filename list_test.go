package collections

import (
	"fmt"
	"testing"
)

func TestList_Source(t *testing.T) {
	list := NewList[int](1, 2, 3, 4, 5, 6)
	list.Add(7)
	fmt.Println(list.source)
	fmt.Println(list.enumerable.source)
	fmt.Println(list.collection.source)
	fmt.Println(list.list.source)
	// {"1", "2"}, {"3", "4"}

}

/*
func TestList_Add(t *testing.T) {
	list := NewList[int]()
	list.Add(1)
	if list.Count() != 1 {
		t.Error()
	}
}

func TestList_Take(t *testing.T) {
	list := NewList[int](1, 2, 3, 4, 5)
	list.Take(3)
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
	list := NewList[int]()
	list.Add(1, 2, 3)
	list.Clear()
	if list.Count() != 0 {
		t.Error()
	}
}

func TestList_MapToList(t *testing.T) {
	type a struct {
		Name string
	}

	type b struct {
		Name string
	}
	list := NewList[a]()
	list.Add(a{"steden1"}, a{"steden2"})

	var list2 List[b]
	list.MapToList(&list2)

	if list2.Count() != 2 {
		t.Error()
	}

	if list.Index(0).Name != list2.Index(0).Name {
		t.Error()
	}

	if list.Index(1).Name != list2.Index(1).Name {
		t.Error()
	}
}

func TestList_MapToArray(t *testing.T) {
	type a struct {
		Name string
	}

	type b struct {
		Name string
	}
	list := NewList[a]()
	list.Add(a{"steden1"}, a{"steden2"})

	var list2 []b
	list.MapToArray(&list2)

	if len(list2) != 2 {
		t.Error()
	}

	if list.Index(0).Name != list2[0].Name {
		t.Error()
	}

	if list.Index(1).Name != list2[1].Name {
		t.Error()
	}
}
*/
