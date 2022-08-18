package collections

import (
	"testing"
)

func Test_enumerable_Any(t *testing.T) {
	lst := NewList[int]()
	if lst.Any() {
		t.Error()
	}
	lst.Add(1)
	if !lst.Any() {
		t.Error()
	}
}

func Test_enumerable_IsEmpty(t *testing.T) {
	lst := NewList[int]()
	if !lst.IsEmpty() {
		t.Error()
	}
	lst.Add(1)
	if lst.IsEmpty() {
		t.Error()
	}
}

func Test_enumerable_First(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	if lst.First() != 1 {
		t.Error()
	}
}

func Test_enumerable_Last(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	if lst.Last() != 6 {
		t.Error()
	}
}

func Test_enumerable_Contains(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if !lst.Contains(4) {
		t.Error()
	}
	if lst.Contains(0) {
		t.Error()
	}
}

func Test_enumerable_Where(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5, 6)
	lst = lst.Where(func(item int) bool {
		return item >= 3
	}).Where(func(item int) bool {
		return item >= 5
	}).ToList()

	if lst.Count() != 2 {
		t.Error()
	}
	if lst.Index(0) != 5 || lst.Index(1) != 6 {
		t.Error()
	}
}

func Test_enumerable_All(t *testing.T) {
	lst := NewList[int](5, 6)
	result := lst.All(func(item int) bool {
		return item == 5 || item == 6
	})
	if !result {
		t.Error()
	}
}
func Test_enumerable_Take(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	lst = lst.Take(3).ToList()
	if lst.Count() != 3 {
		t.Error()
	}
	array := lst.ToArray()
	if array[0] != 1 || array[1] != 2 || array[2] != 3 {
		t.Error()
	}
}

func Test_enumerable_Skip(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	lst = lst.Skip(2).ToList()
	if lst.Count() != 3 {
		t.Error()
	}
	array := lst.ToArray()
	if array[0] != 3 || array[1] != 4 || array[2] != 5 {
		t.Error()
	}
}

func Test_enumerable_Sum(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.Sum(func(item int) any {
		return item - 1
	}) != 10 {
		t.Error()
	}
}

func Test_enumerable_SumItem(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.SumItem() != 15 {
		t.Error()
	}
}

func Test_enumerable_Average(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.Average(func(item int) any {
		return item - 1
	}) != 2 {
		t.Error()
	}
}

func Test_enumerable_AverageItem(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.AverageItem() != 3 {
		t.Error()
	}
}

func Test_enumerable_Min(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.Min(func(item int) any {
		return item - 1
	}) != 0 {
		t.Error()
	}
}

func Test_enumerable_MinItem(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.MinItem() != 1 {
		t.Error()
	}
}

func Test_enumerable_Max(t *testing.T) {
	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.Max(func(item int) any {
		return item - 1
	}) != 4 {
		t.Error()
	}
}

func Test_enumerable_MaxItem(t *testing.T) {

	lst := NewList[int](1, 2, 3, 4, 5)
	if lst.MaxItem() != 5 {
		t.Error()
	}
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

	if len(lstMap) != 2 {
		t.Error()
	}

	if len(lstMap["steden"]) != 2 {
		t.Error()
	}

	if len(lstMap["steden2"]) != 1 {
		t.Error()
	}

	if lstMap["steden"][0].age != 36 {
		t.Error()
	}
	if lstMap["steden"][1].age != 18 {
		t.Error()
	}
	if lstMap["steden2"][0].age != 40 {
		t.Error()
	}
}

func Test_enumerable_OrderBy(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderBy(func(item int) any {
		return item
	}).ToArray()

	if item[0] != 1 || item[1] != 2 || item[2] != 3 || item[3] != 4 || item[4] != 5 || item[5] != 6 || item[6] != 7 || item[7] != 8 {
		t.Error()
	}
}

func Test_enumerable_OrderByItem(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByItem().ToArray()

	if item[0] != 1 || item[1] != 2 || item[2] != 3 || item[3] != 4 || item[4] != 5 || item[5] != 6 || item[6] != 7 || item[7] != 8 {
		t.Error()
	}
}

func Test_enumerable_OrderByDescending(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescending(func(item int) any {
		return item
	}).ToArray()
	if item[0] != 8 || item[1] != 7 || item[2] != 6 || item[3] != 5 || item[4] != 4 || item[5] != 3 || item[6] != 2 || item[7] != 1 {
		t.Error()
	}
}

func Test_enumerable_OrderByDescendingItem(t *testing.T) {
	lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
	item := lst.OrderByDescendingItem().ToArray()
	if item[0] != 8 || item[1] != 7 || item[2] != 6 || item[3] != 5 || item[4] != 4 || item[5] != 3 || item[6] != 2 || item[7] != 1 {
		t.Error()
	}
}

func Test_enumerable_Select(t *testing.T) {
	lst := NewList("1", "", "2")
	var arr []string
	lst.Select(&arr, func(item string) any {
		return "go:" + item
	})

	if len(arr) != lst.Count() {
		t.Error("数量不致")
	}

	for index, yaml := range lst.ToArray() {
		if arr[index] != "go:"+yaml {
			t.Error()
		}
	}
}

func Test_enumerable_SelectMany(t *testing.T) {
	lst := NewList([]string{"1", "2"}, []string{"3", "4"})
	var arr []string
	lst.SelectMany(&arr, func(item []string) any {
		return item
	})

	if len(arr) != 4 {
		t.Error("数量不致")
	}

	if arr[0] != "1" || arr[1] != "2" || arr[2] != "3" || arr[3] != "4" {
		t.Error("数据不正确")
	}
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

	if len(lstMap) != 2 {
		t.Error()
	}

	if len(lstMap["steden"]) != 2 {
		t.Error()
	}

	if len(lstMap["steden2"]) != 1 {
		t.Error()
	}

	if lstMap["steden"][0] != 36 {
		t.Error()
	}
	if lstMap["steden"][1] != 18 {
		t.Error()
	}
	if lstMap["steden2"][0] != 40 {
		t.Error()
	}
}

func Test_enumerable_ToPageList(t *testing.T) {
	lst := NewList(1, 2, 3, 4, 5, 6, 7)
	item := lst.ToPageList(3, 2)
	if item.RecordCount != int64(lst.Count()) {
		t.Error()
	}
	array := item.List.ToArray()
	if array[0] != 4 || array[1] != 5 || array[2] != 6 {
		t.Error()
	}
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

	if lstDO.Count() != 1 {
		t.Error()
	}

	if lstDO.First().Name != "steden" || lstDO.First().Age != 37 {
		t.Error()
	}
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

	if len(lstDO) != 1 {
		t.Error()
	}

	if lstDO[0].Name != "steden" || lstDO[0].Age != 37 {
		t.Error()
	}
}

func Test_enumerable_Intersect(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Intersect(lst2)
	if lst3.Count() != 1 {
		t.Error()
	}
	if lst3.First() != 3 {
		t.Error()
	}
}

func Test_enumerable_Concat(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Concat(lst2)
	if lst3.Count() != 6 {
		t.Error()
	}
	array := lst3.ToArray()
	if array[0] != 1 || array[1] != 2 || array[2] != 3 || array[3] != 3 || array[4] != 4 || array[5] != 5 {
		t.Error()
	}
}

func Test_enumerable_Union(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Union(lst2)
	if lst3.Count() != 5 {
		t.Error()
	}
	array := lst3.ToArray()
	if array[0] != 1 || array[1] != 2 || array[2] != 3 || array[3] != 4 || array[4] != 5 {
		t.Error()
	}
}

func Test_enumerable_Distinct(t *testing.T) {
	lst1 := NewList(1, 2, 3, 3, 4, 5)
	lst2 := lst1.Distinct()

	if lst2.Count() != 5 {
		t.Error()
	}
	array := lst2.ToArray()
	if array[0] != 1 || array[1] != 2 || array[2] != 3 || array[3] != 4 || array[4] != 5 {
		t.Error()
	}
}

func Test_enumerable_Except(t *testing.T) {
	lst1 := NewList(1, 2, 3)
	lst2 := NewList(3, 4, 5)
	lst3 := lst1.Except(lst2)
	if lst3.Count() != 2 {
		t.Error()
	}
	array := lst3.ToArray()
	if array[0] != 1 || array[1] != 2 {
		t.Error()
	}
}

func Test_enumerable_MapToListAny(t *testing.T) {
	type po struct {
		Name string
		Age  int
	}
	lst := NewList(po{Name: "steden", Age: 36}, po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 40})
	listAny := lst.MapToListAny()

	if lst.Count() != listAny.Count() {
		t.Error()
	}

	for i := 0; i < lst.Count(); i++ {
		item := listAny.Index(i).(po)
		if lst.Index(i).Name != item.Name || lst.Index(i).Age != item.Age {
			t.Error()
		}
	}

}
