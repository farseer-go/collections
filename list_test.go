package collections

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestList_Source(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst.Add(7)
	lst.Where(func(item int) bool { return item >= 3 }).
		Where(func(item int) bool { return item >= 5 }).
		Distinct().Skip(1).Take(3).Contains(6)

	flog.Info(lst.source)
	flog.Info(lst.Enumerable.source)
	flog.Info(lst.Collection.source)
	flog.Info(lst.IList.source)

}

func TestList_AsEnumerable(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	enumerable := lst.AsEnumerable()
	lst.Add(4)

	assert.Equal(t, 3, enumerable.Count())
}

func TestReflectItemType(t *testing.T) {
	lst := NewList[int](1, 2, 3)
	lstType := reflect.TypeOf(lst)
	assert.Equal(t, reflect.TypeOf(1), ReflectItemType(lstType))
	assert.NotEqual(t, reflect.TypeOf(true), ReflectItemType(lstType))
}

func TestReflectIsList(t *testing.T) {
	lst := NewList[int](1, 2, 3)
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
	lst := NewList[int](1, 2, 3)
	lstValue := reflect.ValueOf(lst)
	arr := ReflectToArray(lstValue)
	assert.Equal(t, 3, len(arr))
	assert.Equal(t, 1, arr[0])
	assert.Equal(t, 2, arr[1])
	assert.Equal(t, 3, arr[2])

}
