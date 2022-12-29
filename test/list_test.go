package test

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestList_AsEnumerable(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	enumerable := lst.AsEnumerable()
	lst.Add(4)

	assert.Equal(t, 3, enumerable.Count())
}

func TestReflectItemType(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	lstType := reflect.TypeOf(lst)
	assert.Equal(t, reflect.TypeOf(1), collections.ReflectItemType(lstType))
	assert.NotEqual(t, reflect.TypeOf(true), collections.ReflectItemType(lstType))
}

func TestReflectIsList(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	type testStruct struct {
	}
	var test testStruct
	_, isList := types.IsList(reflect.ValueOf(lst))
	assert.True(t, isList)

	_, isList = types.IsList(reflect.ValueOf(&lst))
	assert.True(t, isList)

	_, isList = types.IsList(reflect.ValueOf(1))
	assert.False(t, isList)

	_, isList = types.IsList(reflect.ValueOf(test))
	assert.False(t, isList)
}

func TestReflectToArray(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	lstValue := reflect.ValueOf(lst)
	arr := collections.ReflectToArray(lstValue)
	assert.Equal(t, 3, len(arr))
	assert.Equal(t, 1, arr[0])
	assert.Equal(t, 2, arr[1])
	assert.Equal(t, 3, arr[2])

}

func TestList_Value(t *testing.T) {
	lst := collections.NewList[int](1, 2, 3)
	vals, err := lst.Value()
	assert.Equal(t, err, nil)
	assert.Equal(t, vals, `[1,2,3]`)
	lst = collections.List[int]{}
	vals, err = lst.Value()
	assert.Equal(t, err, nil)
	assert.Equal(t, vals, `{}`)
}

func TestScan(t *testing.T) {
	lst := collections.NewList[int]()
	_ = lst.Scan("1,2,3")
	assert.Equal(t, 3, lst.Count())
	lst = lst.OrderByItem().ToList()
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
	assert.Equal(t, 3, lst.Index(2))

	_ = lst.Scan(nil)
	assert.Equal(t, 0, lst.Count())

	_ = lst.Scan([]byte("1,2,3"))
	assert.Equal(t, 3, lst.Count())
	lst = lst.OrderByItem().ToList()
	assert.Equal(t, 1, lst.Index(0))
	assert.Equal(t, 2, lst.Index(1))
	assert.Equal(t, 3, lst.Index(2))

	assert.NotNil(t, lst.Scan(0))
}

func TestNil(t *testing.T) {
	var lst collections.List[int]
	assert.True(t, lst.IsNil())

	lst = collections.NewList[int]()
	assert.False(t, lst.IsNil())
}
