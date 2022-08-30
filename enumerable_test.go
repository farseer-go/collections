package collections

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_enumerable_Any(t *testing.T) {
	lst := NewList[int]()
	assert.False(t, lst.Any())
	lst.Add(1)
	assert.True(t, lst.Any())
}

func Test_enumerable_IsEmpty(t *testing.T) {
	lst := NewList[int]()
	assert.True(t, lst.IsEmpty())
	lst.Add(1)
	assert.False(t, lst.IsEmpty())
}

func Test_enumerable_First(t *testing.T) {
	lst := NewList[int]()
	assert.Equal(t, 0, lst.First())

	lst = NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, 1, lst.First())
}

func Test_enumerable_Last(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	assert.Equal(t, lst.Last(), 6)

	lst = NewList[int]()
	assert.Equal(t, 0, lst.Last())
}

func Test_enumerable_Contains(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	assert.True(t, lst.Contains(4))
	assert.False(t, lst.Contains(0))
}

func Test_enumerable_Where(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst = lst.Where(func(item int) bool {
		return item >= 3
	}).Where(func(item int) bool {
		return item >= 5
	}).ToList()

	assert.Equal(t, lst.Count(), 2)
	assert.Equal(t, lst.Index(0), 5)
	assert.Equal(t, lst.Index(1), 6)
}

func Test_enumerable_All(t *testing.T) {
	lst := NewList[int](5, 6)
	result := lst.All(func(item int) bool {
		return item == 5 || item == 6
	})
	assert.True(t, result)

	result = lst.All(func(item int) bool {
		return item == 1
	})
	assert.False(t, result)
}

func Test_enumerable_Take(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	lst = lst.Take(3).ToList()
	assert.Equal(t, lst.Count(), 3)
	array := lst.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
	assert.Equal(t, array[2], 3)
}

func Test_enumerable_Skip(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	lst = lst.Skip(2).ToList()
	assert.Equal(t, 3, lst.Count())
	array := lst.ToArray()
	assert.Equal(t, 3, array[0])
	assert.Equal(t, 4, array[1])
	assert.Equal(t, 5, array[2])

	lst = lst.Skip(6).ToList()
	assert.Equal(t, 0, lst.Count())
}

func Test_enumerable_Sum(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	sum := lst.Sum(func(item int) any {
		return item - 1
	})
	assert.Equal(t, sum, 10)
}

func Test_enumerable_SumItem(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	assert.Equal(t, lst.SumItem(), 15)
}

func Test_enumerable_Average(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	avg := lst.Average(func(item int) any {
		return item - 1
	})
	assert.Equal(t, avg, float64(2))
}

func Test_enumerable_AverageItem(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	assert.Equal(t, lst.AverageItem(), float64(3))
}

func Test_enumerable_Min(t *testing.T) {
	lst := NewList[int](5, 2, 4, 1, 3)
	min := lst.Min(func(item int) any {
		return item
	})
	assert.Equal(t, min, 1)

	lst = NewList[int]()
	min = lst.Min(func(item int) any {
		return item
	})
	assert.Equal(t, min, 0)
}

func Test_enumerable_MinItem(t *testing.T) {
	lst := NewList[int](5, 2, 4, 1, 3)
	assert.Equal(t, 1, lst.MinItem())

	lst = NewList[int]()
	assert.Equal(t, 0, lst.MinItem())
}

func Test_enumerable_Max(t *testing.T) {
	lst := NewList[int](2, 4, 5, 1, 3)
	max := lst.Max(func(item int) any {
		return item - 1
	})
	assert.Equal(t, 4, max)

	lst = NewList[int]()
	max = lst.Max(func(item int) any {
		return item
	})
	assert.Equal(t, 0, max)
}

func Test_enumerable_MaxItem(t *testing.T) {
	lst := NewList[int](2, 4, 5, 1, 3)
	assert.Equal(t, 5, lst.MaxItem())

	lst = NewList[int]()
	assert.Equal(t, 0, lst.MaxItem())
}

func Test_enumerable_GroupBy(t *testing.T) {
	type testItem struct {
		name string
		age  int
	}
	lst := NewList[testItem](testItem{name: "steden", age: 36}, testItem{name: "steden", age: 18}, testItem{name: "steden2", age: 40})
	var lstMap map[string][]testItem
	lst.GroupBy(&lstMap, func(item testItem) any {
		return item.name
	})

	assert.Equal(t, len(lstMap), 2)
	assert.Equal(t, len(lstMap["steden"]), 2)
	assert.Equal(t, len(lstMap["steden2"]), 1)
	assert.Equal(t, lstMap["steden"][0].age, 36)
	assert.Equal(t, lstMap["steden"][1].age, 18)
	assert.Equal(t, lstMap["steden2"][0].age, 40)

	assert.Panics(t, func() {
		var arr []string
		lst.GroupBy(&arr, func(item testItem) any {
			return item.name
		})
	})
}

func Test_enumerable_OrderBy(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderBy(func(item int) any {
		return item
	}).ToArray()

	assert.Equal(t, item[0], 1)
	assert.Equal(t, item[1], 2)
	assert.Equal(t, item[2], 3)
	assert.Equal(t, item[3], 4)
	assert.Equal(t, item[4], 5)
	assert.Equal(t, item[5], 6)
	assert.Equal(t, item[6], 7)
	assert.Equal(t, item[7], 8)
}

func Test_enumerable_OrderByItem(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByItem().ToArray()

	assert.Equal(t, item[0], 1)
	assert.Equal(t, item[1], 2)
	assert.Equal(t, item[2], 3)
	assert.Equal(t, item[3], 4)
	assert.Equal(t, item[4], 5)
	assert.Equal(t, item[5], 6)
	assert.Equal(t, item[6], 7)
	assert.Equal(t, item[7], 8)
}

func Test_enumerable_OrderByDescending(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescending(func(item int) any {
		return item
	}).ToArray()

	assert.Equal(t, item[0], 8)
	assert.Equal(t, item[1], 7)
	assert.Equal(t, item[2], 6)
	assert.Equal(t, item[3], 5)
	assert.Equal(t, item[4], 4)
	assert.Equal(t, item[5], 3)
	assert.Equal(t, item[6], 2)
	assert.Equal(t, item[7], 1)
}

func Test_enumerable_OrderByDescendingItem(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescendingItem().ToArray()

	assert.Equal(t, item[0], 8)
	assert.Equal(t, item[1], 7)
	assert.Equal(t, item[2], 6)
	assert.Equal(t, item[3], 5)
	assert.Equal(t, item[4], 4)
	assert.Equal(t, item[5], 3)
	assert.Equal(t, item[6], 2)
	assert.Equal(t, item[7], 1)
}

func Test_enumerable_Select(t *testing.T) {
	lst := NewList("1", "", "2")
	var arr []string
	lst.Select(&arr, func(item string) any {
		return "go:" + item
	})

	assert.Equal(t, len(arr), lst.Count())
	for index, yaml := range lst.ToArray() {
		assert.Equal(t, arr[index], "go:"+yaml)
	}

	var lstSelect List[string]
	lst.Select(&lstSelect, func(item string) any {
		return "go:" + item
	})

	assert.Equal(t, lstSelect.Count(), lst.Count())
	for index, yaml := range lst.ToArray() {
		assert.Equal(t, lstSelect.Index(index), "go:"+yaml)
	}

	assert.Panics(t, func() {
		var lstSelect int
		lst.Select(&lstSelect, func(item string) any {
			return "go:" + item
		})
	})
}

func Test_enumerable_SelectMany(t *testing.T) {
	lst := NewList([]string{"1", "2"}, []string{"3", "4"})

	var arr []string
	lst.SelectMany(&arr, func(item []string) any {
		return item
	})

	assert.Equal(t, len(arr), 4)
	assert.Equal(t, arr[0], "1")
	assert.Equal(t, arr[1], "2")
	assert.Equal(t, arr[2], "3")
	assert.Equal(t, arr[3], "4")

	var lst2 List[string]
	lst.SelectMany(&lst2, func(item []string) any {
		return item
	})

	assert.Equal(t, lst2.Count(), 4)
	assert.Equal(t, lst2.Index(0), "1")
	assert.Equal(t, lst2.Index(1), "2")
	assert.Equal(t, lst2.Index(2), "3")
	assert.Equal(t, lst2.Index(3), "4")

	assert.Panics(t, func() {
		var lst2 List[int]
		lst.SelectMany(&lst2, func(item []string) any {
			return item
		})
	})

	assert.Panics(t, func() {
		var lst2 []int
		lst.SelectMany(&lst2, func(item []string) any {
			return item
		})
	})

	assert.Panics(t, func() {
		var lst2 int
		lst.SelectMany(&lst2, func(item []string) any {
			return item
		})
	})
}

func Test_enumerable_SelectManyItem(t *testing.T) {
	lst := NewList([]string{"1", "2"}, []string{"3", "4"})

	var arr []string
	lst.SelectManyItem(&arr)

	assert.Equal(t, len(arr), 4)
	assert.Equal(t, arr[0], "1")
	assert.Equal(t, arr[1], "2")
	assert.Equal(t, arr[2], "3")
	assert.Equal(t, arr[3], "4")

	var lst2 List[string]
	lst.SelectManyItem(&lst2)

	assert.Equal(t, lst2.Count(), 4)
	assert.Equal(t, lst2.Index(0), "1")
	assert.Equal(t, lst2.Index(1), "2")
	assert.Equal(t, lst2.Index(2), "3")
	assert.Equal(t, lst2.Index(3), "4")
}

func Test_enumerable_ToMap(t *testing.T) {
	type testItem struct {
		name string
		age  int
	}
	lst := NewList[testItem](testItem{name: "steden", age: 36}, testItem{name: "steden", age: 18}, testItem{name: "steden2", age: 40})
	var lstMap map[string][]int
	lst.ToMap(&lstMap,
		func(key testItem) any {
			return key.name
		}, func(value testItem) any {
			return value.age
		})

	assert.Equal(t, len(lstMap), 2)
	assert.Equal(t, len(lstMap["steden"]), 2)
	assert.Equal(t, len(lstMap["steden2"]), 1)
	assert.Equal(t, lstMap["steden"][0], 36)
	assert.Equal(t, lstMap["steden"][1], 18)
	assert.Equal(t, lstMap["steden2"][0], 40)

	assert.Panics(t, func() {
		var arr []string
		lst.ToMap(&arr,
			func(key testItem) any {
				return key.name
			}, func(value testItem) any {
				return value.age
			})
	})
}

func Test_enumerable_ToPageList(t *testing.T) {
	lst := NewList(1, 2, 3, 4, 5, 6, 7)
	item := lst.ToPageList(3, 2)
	assert.Equal(t, item.RecordCount, int64(lst.Count()))
	array := item.List.ToArray()
	assert.Equal(t, array[0], 4)
	assert.Equal(t, array[1], 5)
	assert.Equal(t, array[2], 6)

	assert.Panics(t, func() {
		lst.ToPageList(0, 0)
	})
	assert.Panics(t, func() {
		lst.ToPageList(3, 0)
	})

	item = lst.ToPageList(8, 2)
	array = item.List.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
	assert.Equal(t, array[2], 3)
	assert.Equal(t, array[3], 4)
	assert.Equal(t, array[4], 5)
	assert.Equal(t, array[5], 6)
	assert.Equal(t, array[6], 7)
}

func Test_enumerable_MapToList(t *testing.T) {
	type po struct {
		Name string
		Age  int
	}
	type do struct {
		Name string
		Age  int
	}

	lst := NewList(po{Name: "steden", Age: 37})
	var lstDO List[do]
	lst.MapToList(&lstDO)

	assert.Equal(t, 1, lstDO.Count())
	assert.Equal(t, "steden", lstDO.First().Name)
	assert.Equal(t, 37, lstDO.First().Age)

	lstAny := NewListAny(1, 2)
	var lstInt List[int]
	lstAny.MapToList(&lstInt)

	assert.Equal(t, lstAny.Count(), lstInt.Count())
	assert.Equal(t, lstAny.Index(0), lstInt.Index(0))
	assert.Equal(t, lstAny.Index(1), lstInt.Index(1))

	assert.Panics(t, func() {
		var lstInt []int
		lstAny.MapToList(&lstInt)
	})
}

func Test_enumerable_MapToArray(t *testing.T) {
	type po struct {
		Name string
		Age  int
	}
	type do struct {
		Name string
		Age  int
	}

	lst := NewList(po{Name: "steden", Age: 37})
	var lstDO []do
	lst.MapToArray(&lstDO)

	assert.Equal(t, 1, len(lstDO))
	assert.Equal(t, "steden", lstDO[0].Name)
	assert.Equal(t, 37, lstDO[0].Age)

	lstAny := NewListAny(1, 2)
	var arrInt []int
	lstAny.MapToArray(&arrInt)

	assert.Equal(t, lstAny.Count(), len(arrInt))
	assert.Equal(t, lstAny.Index(0), arrInt[0])
	assert.Equal(t, lstAny.Index(1), arrInt[1])

	assert.Panics(t, func() {
		var val int
		lstAny.MapToArray(&val)
	})
}

func Test_enumerable_Intersect(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Intersect(lst2)
	assert.Equal(t, lst3.Count(), 1)
	assert.Equal(t, lst3.First(), 3)
}

func Test_enumerable_Concat(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Concat(lst2)
	assert.Equal(t, lst3.Count(), 6)
	array := lst3.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
	assert.Equal(t, array[2], 3)
	assert.Equal(t, array[3], 3)
	assert.Equal(t, array[4], 4)
	assert.Equal(t, array[5], 5)
}

func Test_enumerable_Union(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Union(lst2)
	assert.Equal(t, lst3.Count(), 5)
	array := lst3.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
	assert.Equal(t, array[2], 3)
	assert.Equal(t, array[3], 4)
	assert.Equal(t, array[4], 5)
}

func Test_enumerable_Distinct(t *testing.T) {
	lst1 := NewList(1, 2, 3, 3, 4, 5)
	lst2 := lst1.Distinct()

	assert.Equal(t, lst2.Count(), 5)
	array := lst2.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
	assert.Equal(t, array[2], 3)
	assert.Equal(t, array[3], 4)
	assert.Equal(t, array[4], 5)
}

func Test_enumerable_Except(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Except(lst2)
	assert.Equal(t, lst3.Count(), 2)
	array := lst3.ToArray()
	assert.Equal(t, array[0], 1)
	assert.Equal(t, array[1], 2)
}

func Test_enumerable_MapToListAny(t *testing.T) {
	type po struct {
		Name string
		Age  int
	}
	lst := NewList(po{Name: "steden", Age: 36}, po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 40})
	listAny := lst.ToListAny()

	assert.Equal(t, lst.Count(), listAny.Count())

	for i := 0; i < lst.Count(); i++ {
		item := listAny.Index(i).(po)
		assert.Equal(t, lst.Index(i).Name, item.Name)
		assert.Equal(t, lst.Index(i).Age, item.Age)
	}
}

func Test_enumerable_Range(t *testing.T) {
	lst1 := NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	lstCut := lst1.Range(3, 2)
	assert.Equal(t, 10, lst1.Count())
	assert.Equal(t, 2, lstCut.Count())
	assert.Equal(t, 4, lstCut.ToArray()[0])
	assert.Equal(t, 5, lstCut.ToArray()[1])

}

func Test_enumerable_RangeStart(t *testing.T) {
	lst1 := NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	lstCut := lst1.RangeStart(7)
	assert.Equal(t, 10, lst1.Count())
	assert.Equal(t, 3, lstCut.Count())
	assert.Equal(t, 8, lstCut.ToArray()[0])
	assert.Equal(t, 9, lstCut.ToArray()[1])
	assert.Equal(t, 10, lstCut.ToArray()[2])
}

func TestEnumerable_Rand(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 100; i++ {
		val := NewList(1, 2, 3).Rand()
		if val != 1 && val != 2 && val != 3 {
			t.Error()
		}
	}

	val := NewList(1).Rand()
	assert.Equal(t, 1, val)
}

func TestEnumerable_ToString(t *testing.T) {
	assert.Equal(t, "1,2,3", NewList(1, 2, 3).ToString(","))
}

func TestEnumerable_Empty(t *testing.T) {
	lst1 := NewList(1, 2, 3).Empty()
	assert.Equal(t, 0, lst1.Count())
}
