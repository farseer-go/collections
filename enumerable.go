package collections

import (
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/types"
	"github.com/timandy/routine"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Enumerable[T any] struct {
	source *[]T
	lock   *sync.RWMutex
}

// Any 是否存在
func (receiver Enumerable[T]) Any() bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return len(*receiver.source) > 0
}

// IsEmpty 集合是为空的
func (receiver Enumerable[T]) IsEmpty() bool {
	if receiver.lock == nil {
		return true
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return receiver.source == nil || len(*receiver.source) == 0
}

// First 查找符合条件的第一个元素
func (receiver Enumerable[T]) First() T {
	if receiver.lock == nil {
		var t T
		return t
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	if len(*receiver.source) > 0 {
		return (*receiver.source)[0]
	}
	var t T
	return t
}

// FirstAddr 返回第一个指针元素
func (receiver Enumerable[T]) FirstAddr() *T {
	if receiver.lock == nil {
		return nil
	}
	item := &(*receiver.source)[0]
	return item
}

// Last 集合最后一个元素
func (receiver Enumerable[T]) Last() T {
	if receiver.lock == nil {
		var t T
		return t
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	if len(*receiver.source) > 0 {
		return (*receiver.source)[len(*receiver.source)-1]
	}
	var t T
	return t
}

// Count 集合大小
func (receiver Enumerable[T]) Count() int {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return len(*receiver.source)
}

// Contains 是否包含元素
func (receiver Enumerable[T]) Contains(item T) bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, t := range *receiver.source {
		if parse.IsEqual(t, item) {
			return true
		}
	}
	return false
}

// ContainsCount 获取包含几个元素
func (receiver Enumerable[T]) ContainsCount(item T) int {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	count := 0
	for _, t := range *receiver.source {
		if parse.IsEqual(t, item) {
			count++
		}
	}
	return count
}

// ContainsPrefix 元素中是否包含前缀
func (receiver Enumerable[T]) ContainsPrefix(item T) bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, t := range *receiver.source {
		if strings.HasPrefix(parse.ToString(t), parse.ToString(item)) {
			return true
		}
	}
	return false
}

// ContainsSuffix 元素中是否包含后缀
func (receiver Enumerable[T]) ContainsSuffix(item T) bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, t := range *receiver.source {
		if strings.HasSuffix(parse.ToString(t), parse.ToString(item)) {
			return true
		}
	}
	return false
}

// ContainsAny 元素中是否包含item（strings.Contains）
func (receiver Enumerable[T]) ContainsAny(item T) bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, t := range *receiver.source {
		if strings.Contains(parse.ToString(t), parse.ToString(item)) {
			return true
		}
	}
	return false
}

// Where 对数据进行筛选
func (receiver Enumerable[T]) Where(fn func(item T) bool) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var lst []T
	for _, item := range *receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// All 是否所有数据都满足fn条件
func (receiver Enumerable[T]) All(fn func(item T) bool) bool {
	if receiver.lock == nil {
		return false
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	for _, item := range *receiver.source {
		if !fn(item) {
			return false
		}
	}
	return true
}

// Take 返回前多少条数据
func (receiver Enumerable[T]) Take(count int) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	recordCount := len(*receiver.source)
	// 总长度比count小，则直接返回全部数据
	if recordCount < count {
		return receiver
	}
	arr := (*receiver.source)[0:count]
	return Enumerable[T]{source: &arr, lock: &sync.RWMutex{}}
}

// Skip 跳过前多少条记录
func (receiver Enumerable[T]) Skip(count int) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	recordCount := len(*receiver.source)
	// 总长度比count小，则返回空数据
	if recordCount < count {
		return Enumerable[T]{source: new([]T), lock: &sync.RWMutex{}}
	}
	arr := (*receiver.source)[count:]
	return Enumerable[T]{source: &arr, lock: &sync.RWMutex{}}
}

// Sum 求总和
func (receiver Enumerable[T]) Sum(fn func(item T) any) any {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source
	var sum any
	for index := 0; index < len(lst); index++ {
		sum = Addition(sum, fn(lst[index]))
	}
	return sum
}

// SumItem 求总和
func (receiver Enumerable[T]) SumItem() T {
	if receiver.lock == nil {
		var t T
		return t
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source
	var sum T
	for index := 0; index < len(lst); index++ {
		sum = Addition(sum, lst[index]).(T)
	}
	return sum
}

// Average 求平均数
func (receiver Enumerable[T]) Average(fn func(item T) any) float64 {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	sum := receiver.Sum(fn)
	count := len(*receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}

// AverageItem 求平均数
func (receiver Enumerable[T]) AverageItem() float64 {
	if receiver.lock == nil {
		return 0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	sum := receiver.Sum(func(item T) any { return item })
	count := len(*receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}

// Min 获取最小值
func (receiver Enumerable[T]) Min(fn func(item T) any) any {
	if receiver.lock == nil {
		var t T
		return fn(t) // 不能直接返回0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source

	if len(lst) == 0 {
		var t T
		return fn(t) // 不能直接返回0
	}
	minValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if CompareLeftGreaterThanRight(minValue, value) {
			minValue = value
		}
	}
	return minValue
}

// MinItem 获取最小值
func (receiver Enumerable[T]) MinItem() T {
	if receiver.lock == nil {
		var t T
		return t
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source

	if len(lst) == 0 {
		var t T
		return t
	}
	minValue := lst[0]
	for index := 1; index < len(lst); index++ {
		value := lst[index]
		if CompareLeftGreaterThanRight(minValue, value) {
			minValue = value
		}
	}
	return minValue
}

// Max 获取最大值
func (receiver Enumerable[T]) Max(fn func(item T) any) any {
	if receiver.lock == nil {
		var t T
		return fn(t) // 不能直接返回0
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source

	if len(lst) == 0 {
		var t T
		return fn(t) // 不能直接返回0
	}

	maxValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if CompareLeftGreaterThanRight(value, maxValue) {
			maxValue = value
		}
	}
	return maxValue
}

// MaxItem 获取最大值
func (receiver Enumerable[T]) MaxItem() T {
	if receiver.lock == nil {
		var t T
		return t
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := *receiver.source

	if len(lst) == 0 {
		var t T
		return t
	}
	maxValue := lst[0]
	for index := 1; index < len(lst); index++ {
		value := lst[index]
		if CompareLeftGreaterThanRight(value, maxValue) {
			maxValue = value
		}
	}
	return maxValue
}

// GroupBy 将数组进行分组后返回map
//
//	var lstMap map[string][]testItem
//	lst.GroupBy(&lstMap, func(item testItem) any {
//		return item.name
//	})
//	or
//	var lstMap map[string]collections.List[testItem]
//	lst.GroupBy(&lstMap, func(item testItem) any {
//		return item.name
//	})
func (receiver Enumerable[T]) GroupBy(mapSlice any, getMapKeyFunc func(item T) any) {
	if receiver.lock == nil {
		return
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	mapSliceVal := reflect.ValueOf(mapSlice).Elem()
	mapSliceType, isMap := types.IsMap(mapSliceVal)
	if !isMap {
		panic("mapSlice入参必须为map类型")
	}
	// mapSlice的Value是切片还是List类型
	mapValueItemType := mapSliceType.Elem()
	_, mapValueIsListType := types.IsListByType(mapValueItemType)

	// make....
	mapSliceVal.Set(reflect.MakeMap(mapSliceType))

	nilValueOf := reflect.ValueOf(nil)
	for _, item := range *receiver.source {
		// 生成key
		key := reflect.ValueOf(getMapKeyFunc(item))
		// 根据key，找到map的value
		findMapValue := mapSliceVal.MapIndex(key)
		// nil说明map不存在这个key
		if findMapValue == nilValueOf {
			// List集合
			if mapValueIsListType {
				findMapValue = types.ListNew(mapValueItemType)
			} else {
				findMapValue = reflect.MakeSlice(mapValueItemType, 0, 0)
			}
		}
		// List集合
		if mapValueIsListType {
			// 在原来的List集合，做Add操作，必须保住该集合是指针类型
			if findMapValue.Kind() != reflect.Pointer {
				arrValues := types.GetListToArrayValue(findMapValue)
				findMapValue = types.ListNew(mapValueItemType)
				types.ListAddValue(findMapValue, arrValues)
			}
			types.ListAddValue(findMapValue, reflect.ValueOf(item))
			mapSliceVal.SetMapIndex(key, findMapValue.Elem())
		} else {
			mapValue := reflect.Append(findMapValue, reflect.ValueOf(item))
			mapSliceVal.SetMapIndex(key, mapValue)
		}
	}
}

// OrderBy 正序排序，fn 返回的是要排序的字段的值
func (receiver Enumerable[T]) OrderBy(fn func(item T) any) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var lst []T
	lst = append(lst, *receiver.source...)

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if CompareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = fn(lst[leftIndex])
			}
		}
	}

	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// OrderByItem 正序排序，fn 返回的是要排序的字段的值
func (receiver Enumerable[T]) OrderByItem() Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var lst []T
	lst = append(lst, *receiver.source...)

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := lst[leftIndex]

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := lst[rightIndex] // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if CompareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = lst[leftIndex]
			}
		}
	}

	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// OrderByDescending 倒序排序，fn 返回的是要排序的字段的值
func (receiver Enumerable[T]) OrderByDescending(fn func(item T) any) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	var lst []T
	lst = append(lst, *receiver.source...)

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if !CompareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = fn(lst[leftIndex])
			}
		}
	}

	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// OrderByDescendingItem 倒序排序，fn 返回的是要排序的字段的值
func (receiver Enumerable[T]) OrderByDescendingItem() Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var lst []T
	lst = append(lst, *receiver.source...)

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := lst[leftIndex]

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := lst[rightIndex] // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if !CompareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = lst[leftIndex]
			}
		}
	}
	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// Select 筛选子元素字段
//
// arrSlice：切片数组类型
//
// eg:
//
//	lstYaml := NewList("1", "", "2")
//	var lst []string
//	lstYaml.Select(&lst, func(item string) any {
//	    return "go:" + item
//	})
//	result: lst = []string { "go:1", "go:", "go:2" }
//
//	var lstSelect List[string]
//	lstYaml.Select(&lstSelect, func(item string) any {
//		return "go:" + item
//	})
//	result: lstSelect = List[string] { "go:1", "go:", "go:2" }
func (receiver Enumerable[T]) Select(sliceOrList any, fn func(item T) any) {
	if receiver.lock == nil {
		return
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	sliceOrListVal := reflect.ValueOf(sliceOrList).Elem()
	// 切片类型
	if sliceOrListType, isSlice := types.IsSlice(sliceOrListVal); isSlice {
		value := reflect.MakeSlice(sliceOrListType, 0, 0)
		for _, item := range *receiver.source {
			value = reflect.Append(value, reflect.ValueOf(fn(item)))
		}
		sliceOrListVal.Set(value)
		return
	}
	if sliceOrListType, isList := types.IsList(sliceOrListVal); isList {
		// 初始化
		value := types.ListNew(sliceOrListType)

		for _, item := range *receiver.source {
			types.ListAdd(value, fn(item))
		}
		sliceOrListVal.Set(value.Elem())
		return
	}

	panic("sliceOrList入参必须为切片或collections.List类型")
}

// SelectMany 筛选子元素字段
//
// arrSlice：切片数组类型
//
// eg:
//
//	lstYaml := NewList([]string{"1", "2"}, []string{"3", "4"})
//	var arr []string
//	lstYaml.SelectMany(&arr, func(item []string) any {
//		return item
//	})
//	// result:	arr = []string { "1", "2", "3", "4" }
//
//	var lst2 List[string]
//	lst.SelectMany(&lst2, func(item []string) any {
//		return item
//	})
//	// result:	lst = List[string] { "1", "2", "3", "4" }
func (receiver Enumerable[T]) SelectMany(sliceOrList any, fn func(item T) any) {
	if receiver.lock == nil {
		return
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	sliceOrListVal := reflect.ValueOf(sliceOrList).Elem()

	// 切片类型
	if sliceOrListType, isSlice := types.IsSlice(sliceOrListVal); isSlice {
		value := reflect.MakeSlice(sliceOrListType, 0, 0)
		for _, item := range *receiver.source {
			itemValue := reflect.ValueOf(fn(item))
			if itemValue.Type() != sliceOrListType {
				panic("sliceOrList入参类型必须与fn返回的类型一致")
			}
			value = reflect.AppendSlice(value, itemValue)
		}
		sliceOrListVal.Set(value)
		return
	}
	if sliceOrListType, isList := types.IsList(sliceOrListVal); isList {
		// 初始化
		value := types.ListNew(sliceOrListType)

		for _, item := range *receiver.source {
			types.ListAdd(value, fn(item))
		}
		sliceOrListVal.Set(value.Elem())
		return
	}

	panic("sliceOrList入参必须为切片或collections.List类型")
}

// SelectManyItem 筛选子元素字段
//
// arrSlice：切片数组类型
//
// eg:
//
//	lstYaml := NewList([]string{"1", "2"}, []string{"3", "4"})
//	var arr []string
//	lstYaml.SelectMany(&arr)
//	// result:	arr = []string { "1", "2", "3", "4" }
//
//	var lst2 List[string]
//	lst.SelectMany(&lst2)
//	// result:	lst = List[string] { "1", "2", "3", "4" }
func (receiver Enumerable[T]) SelectManyItem(sliceOrList any) {
	if receiver.lock == nil {
		return
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	receiver.SelectMany(sliceOrList, func(item T) any {
		return item
	})
}

// ToMap 转成字典
func (receiver Enumerable[T]) ToMap(mapSlice any, getMapKeyFunc func(item T) any, getMapValueFunc func(item T) any) {
	if receiver.lock == nil {
		return
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	mapSliceVal := reflect.ValueOf(mapSlice).Elem()
	mapSliceType, isMap := types.IsMap(mapSliceVal)
	if !isMap {
		panic("mapSlice入参必须为map类型")
	}

	// make.... var m map[k]v = make(map[k]v)
	mapSliceVal.Set(reflect.MakeMap(mapSliceType))

	for _, item := range *receiver.source {
		// 生成key
		key := reflect.ValueOf(getMapKeyFunc(item))
		// 根据key，找到map的value
		findMapValue := mapSliceVal.MapIndex(key)
		// nil说明map不存在这个key
		if findMapValue == reflect.ValueOf(nil) {
			findMapValue = reflect.MakeSlice(mapSliceType.Elem(), 0, 0)
			mapSliceVal.SetMapIndex(key, findMapValue)
		}
		mapValue := reflect.Append(findMapValue, reflect.ValueOf(getMapValueFunc(item)))
		mapSliceVal.SetMapIndex(key, mapValue)
	}
}

// ToList 返回List集合
func (receiver Enumerable[T]) ToList() List[T] {
	if receiver.lock == nil {
		return NewList[T]()
	}

	return NewList(*receiver.source...)
}

// ToArray 转成数组
func (receiver Enumerable[T]) ToArray() []T {
	if receiver.lock == nil {
		return []T{}
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return *receiver.source
}

// ToPageList 数组分页
func (receiver Enumerable[T]) ToPageList(pageSize int, pageIndex int) PageList[T] {
	if receiver.lock == nil {
		return PageList[T]{}
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	pageList := PageList[T]{
		RecordCount: int64(len(*receiver.source)),
	}

	if pageList.RecordCount < 1 {
		pageList.List = NewList(*receiver.source...)
		return pageList
	}

	if pageSize < 1 {
		panic("pageSize不能小于1")
	}

	// 计算总页数
	var allCurrentPage int64
	// 没有设置pageIndex，则按take返回
	if pageIndex < 1 {
		panic("pageIndex不能小于1")
	}

	allCurrentPage = pageList.RecordCount / int64(pageSize)
	if pageList.RecordCount%int64(pageSize) != 0 {
		allCurrentPage++
	}

	if int64(pageIndex) > allCurrentPage {
		pageIndex = int(allCurrentPage)
	}
	skipCount := pageSize * (pageIndex - 1)
	endIndex := skipCount + pageSize
	if endIndex > len(*receiver.source) {
		endIndex = len(*receiver.source)
	}
	lst := (*receiver.source)[skipCount:endIndex]
	pageList.List = NewList(lst...)
	return pageList
}

// ToListAny 转成ListAny
func (receiver Enumerable[T]) ToListAny() ListAny {
	lst := NewListAny()
	if receiver.lock == nil {
		return lst
	}

	array := receiver.ToArray()
	for _, item := range array {
		lst.Add(item)
	}
	return lst
}

// Empty 返回一个新的Empty集合
func (receiver Enumerable[T]) Empty() Enumerable[T] {
	if receiver.lock == nil {
		list := NewList[T]()
		return list.AsEnumerable()
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	return Enumerable[T]{source: &[]T{}, lock: &sync.RWMutex{}}
}

// Intersect 两个集合的交集（共同存在的元素）
func (receiver Enumerable[T]) Intersect(lstRight List[T]) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var lst []T
	for _, item := range *receiver.source {
		if lstRight.Contains(item) {
			lst = append(lst, item)
		}
	}
	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// Concat 合并两个集合
func (receiver Enumerable[T]) Concat(lstRight List[T]) Enumerable[T] {
	if receiver.lock == nil {
		return lstRight.AsEnumerable()
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := append(*receiver.source, *lstRight.source...)
	return Enumerable[T]{source: &lst, lock: &sync.RWMutex{}}
}

// Union 合并两个集合，并去重
func (receiver Enumerable[T]) Union(lstRight List[T]) Enumerable[T] {
	if receiver.lock == nil {
		return lstRight.AsEnumerable()
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	union := receiver.Concat(lstRight)
	return union.Distinct()
}

// Distinct 集合去重
func (receiver Enumerable[T]) Distinct() Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := NewList[T]()

	for _, item := range *receiver.source {
		if !lst.Where(func(t T) bool { return reflect.DeepEqual(item, t) }).Any() {
			lst.Add(item)
		}
	}
	return lst.Enumerable
}

// Reverse 集合反转
func (receiver Enumerable[T]) Reverse() List[T] {
	if receiver.lock == nil {
		return receiver.ToList()
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := NewList[T]()
	for i := len(*receiver.source) - 1; i >= 0; i-- {
		lst.Add((*receiver.source)[i])
	}
	return lst
}

// Except 移除参数中包含的集合元素
func (receiver Enumerable[T]) Except(lstRight List[T]) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	lst := receiver.ToList()
	for _, item := range *lstRight.source {
		lst.Remove(item)
	}
	return lst.Enumerable
}

// Range 获取切片范围
// startIndex：起始位置
// length：从startIndex开始之后的长度
func (receiver Enumerable[T]) Range(startIndex int, length int) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	newList := (*receiver.source)[startIndex : startIndex+length]
	return Enumerable[T]{source: &newList, lock: &sync.RWMutex{}}
}

// RangeStart 获取切片范围（指定startIndex，不指定endIndex)
func (receiver Enumerable[T]) RangeStart(startIndex int) Enumerable[T] {
	if receiver.lock == nil {
		return receiver
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	newList := (*receiver.source)[startIndex:]
	return Enumerable[T]{source: &newList, lock: &sync.RWMutex{}}
}

// Rand 返回随机元素
func (receiver Enumerable[T]) Rand() *T {
	if receiver.lock == nil {
		return nil
	}

	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	if receiver.Count() < 2 {
		return &(*receiver.source)[0]
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(receiver.Count())
	return &(*receiver.source)[random]
}

// ToString 将集合转成字符串，并用split分隔
func (receiver Enumerable[T]) ToString(split string) string {
	if receiver.lock == nil {
		return ""
	}
	receiver.lock.RLock()
	defer receiver.lock.RUnlock()

	var arrStr []string
	for _, item := range *receiver.source {
		arrStr = append(arrStr, parse.Convert(item, ""))
	}
	return strings.Join(arrStr, split)
}

// For 遍历操作
func (receiver Enumerable[T]) For(itemFn func(index int, item *T)) {
	if receiver.lock == nil {
		return
	}
	for i := 0; i < len(*receiver.source); i++ {
		item := &(*receiver.source)[i]
		itemFn(i, item)
	}
}

// Foreach for range操作
func (receiver Enumerable[T]) Foreach(itemFn func(item *T)) {
	if receiver.lock == nil {
		return
	}
	for i := 0; i < len(*receiver.source); i++ {
		item := &(*receiver.source)[i]
		itemFn(item)
	}
}

// Find 查找指定条件的元素，返回批一个指针元素
func (receiver Enumerable[T]) Find(itemFn func(item *T) bool) *T {
	if receiver.lock == nil {
		return nil
	}
	for i := 0; i < len(*receiver.source); i++ {
		item := &(*receiver.source)[i]
		if itemFn(item) {
			return item
		}
	}
	return nil
}

// Parallel for range 并行操作
func (receiver Enumerable[T]) Parallel(itemFn func(item *T)) {
	if receiver.lock == nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(*receiver.source))
	for i := 0; i < len(*receiver.source); i++ {
		item := &(*receiver.source)[i]
		routine.Go(func() {
			defer wg.Done()
			itemFn(item)
		})
	}
	wg.Wait()
}
