package test

import (
	"github.com/farseer-go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

type po struct {
	Name string
	Age  int
}
type do struct {
	Name string
	Age  int
}

func TestPageList_MapToPageList(t *testing.T) {
	pageList := collections.NewPageList[po](collections.NewList(po{Name: "steden", Age: 18}), 10)
	assert.Equal(t, int64(10), pageList.RecordCount)
	//
	//	var newPageList collections.PageList[do]
	//	pageList.MapToPageList(&newPageList)
	//
	//	assert.Equal(t, pageList.RecordCount, newPageList.RecordCount)
	//	assert.Equal(t, pageList.List.Count(), newPageList.List.Count())
	//
	//	for i := 0; i < pageList.List.Count(); i++ {
	//		assert.Equal(t, pageList.List.Index(i).Name, newPageList.List.Index(i).Name)
	//		assert.Equal(t, pageList.List.Index(i).Age, newPageList.List.Index(i).Age)
	//	}
}
