package test

import (
	"encoding/json"
	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_collection_Count(t *testing.T) {
	lst := collections.NewList[int]()
	lst.Add(1, 2, 3) // lst = 1, 2, 3
	lst.Insert(1, 8) // lst = 1, 8, 2, 3
	assert.Equal(t, 4, lst.Count())
}

func Test_collection_Add(t *testing.T) {
	lst := collections.NewList[int](1, 2)
	lst.Add(3)
	assert.Equal(t, 3, lst.Count())
	assert.Equal(t, 3, lst.Index(2))
}

func Test_collection_AddRange(t *testing.T) {
	lst := collections.NewList[int]()
	lst2 := collections.NewList[int](1, 2, 3)
	lst.AddRange(lst2.AsEnumerable())

	assert.Equal(t, 3, lst.Count())
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
	assert.Equal(t, 3, lst.Index(2))
}

func Test_collection_Clear(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	assert.Equal(t, 3, lst.Count())
	lst.Clear()
	assert.Equal(t, 0, lst.Count())
}

func Test_collection_RemoveAll(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3, 6)
	lst.RemoveAll(func(item int) bool {
		return item >= 3
	})
	assert.Equal(t, 2, lst.Count())
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
}

func TestCollection_MarshalJSON(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3, 6)
	strjson, _ := lst.MarshalJSON()
	retjson, _ := json.Marshal(lst)
	assert.Equal(t, retjson, strjson)
}

func TestCollection_UnmarshalJSON(t *testing.T) {
	lst := collections.NewList[string]()
	jsonData := []byte(`["sam","18"]`)
	err := lst.UnmarshalJSON(jsonData)
	maps := lst.ToArray()
	assert.Equal(t, nil, err)
	assert.Equal(t, "sam", maps[0])
	assert.Equal(t, "18", maps[1])
}
