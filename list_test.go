package collections

import (
	"fmt"
	"testing"
)

func TestList_Source(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst.Add(7)
	lst.Where(func(item int) bool { return item >= 3 }).
		Where(func(item int) bool { return item >= 5 }).
		Distinct().Skip(1).Take(3).Contains(6)

	fmt.Println(lst.source)
	fmt.Println(lst.Enumerable.source)
	fmt.Println(lst.Collection.source)
	fmt.Println(lst.IList.source)

}
