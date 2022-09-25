package collections

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_collection_Count(t *testing.T) {
	lst := NewList[int]()
	lst.Add(1, 2, 3) // lst = 1, 2, 3
	lst.Insert(1, 8) // lst = 1, 8, 2, 3
	assert.Equal(t, lst.Count(), 4)
}

func Test_collection_Add(t *testing.T) {
	lst := NewList[int](1, 2)
	lst.Add(3)
	assert.Equal(t, lst.Count(), 3)
	assert.Equal(t, lst.Index(2), 3)
}

func Test_collection_Clear(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	assert.Equal(t, lst.Count(), 3)
	lst.Clear()
	assert.Equal(t, lst.Count(), 0)
}

func Test_collection_RemoveAll(t *testing.T) {
	lst := NewList[int](1, 2, 3, 6)
	lst.RemoveAll(func(item int) bool {
		return item >= 3
	})
	assert.Equal(t, lst.Count(), 2)
	assert.Equal(t, lst.Index(0), 1)
	assert.Equal(t, lst.Index(1), 2)
}

func Test_collection_MarshalJSON(t *testing.T) {
	lst := NewList[int](1, 2, 3, 6)
	strjson, _ := lst.MarshalJSON()
	retjson, _ := json.Marshal(lst.source)
	assert.Equal(t, strjson, retjson)
	lst2 := NewList[interface{}](nil)
	strjson, err := lst2.MarshalJSON()
	assert.Equal(t, string(strjson), "[null]")
	assert.Equal(t, err, nil)
}
