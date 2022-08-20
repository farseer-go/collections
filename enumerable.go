package collections

import (
	"github.com/devfeel/mapper"
	"github.com/farseer-go/fs/parse"
	"reflect"
	"strings"
)

type enumerable[T any] struct {
	source *[]T
}

// Any 是否存在
func (receiver enumerable[T]) Any() bool {
	return len(*receiver.source) > 0
}

// IsEmpty 集合是为空的
func (receiver enumerable[T]) IsEmpty() bool {
	return receiver.source == nil || len(*receiver.source) == 0
}

// First 查找符合条件的第一个元素
func (receiver enumerable[T]) First() T {
	if len(*receiver.source) > 0 {
		return (*receiver.source)[0]
	}
	var t T
	return t
}

// Last 集合最后一个元素
func (receiver enumerable[T]) Last() T {
	if len(*receiver.source) > 0 {
		return (*receiver.source)[len(*receiver.source)-1]
	}
	var t T
	return t
}

// Count 集合大小
func (receiver enumerable[T]) Count() int {
	return len(*receiver.source)
}

// Contains 是否包含元素
func (receiver enumerable[T]) Contains(item T) bool {
	itemValue := reflect.ValueOf(item)
	for _, t := range *receiver.source {
		if reflect.ValueOf(t) == itemValue {
			return true
		}
	}
	return false
}

// Where 对数据进行筛选
func (receiver enumerable[T]) Where(fn func(item T) bool) enumerable[T] {
	var lst []T
	for _, item := range *receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	return enumerable[T]{source: &lst}
}

// All 是否所有数据都满足fn条件
func (receiver enumerable[T]) All(fn func(item T) bool) bool {
	for _, item := range *receiver.source {
		if !fn(item) {
			return false
		}
	}
	return true
}

// Take 返回前多少条数据
func (receiver enumerable[T]) Take(count int) enumerable[T] {
	recordCount := len(*receiver.source)
	// 总长度比count小，则直接返回全部数据
	if recordCount < count {
		return receiver
	}
	arr := (*receiver.source)[0:count]
	return enumerable[T]{source: &arr}
}

// Skip 跳过前多少条记录
func (receiver enumerable[T]) Skip(count int) enumerable[T] {
	recordCount := len(*receiver.source)
	// 总长度比count小，则返回空数据
	if recordCount < count {
		return enumerable[T]{source: new([]T)}
	}
	arr := (*receiver.source)[count:]
	return enumerable[T]{source: &arr}
}

// Sum 求总和
func (receiver enumerable[T]) Sum(fn func(item T) any) any {
	lst := *receiver.source
	var sum any
	for index := 0; index < len(lst); index++ {
		sum = Addition(sum, fn(lst[index]))
	}
	return sum
}

// SumItem 求总和
func (receiver enumerable[T]) SumItem() T {
	lst := *receiver.source
	var sum T
	for index := 0; index < len(lst); index++ {
		sum = Addition(sum, lst[index]).(T)
	}
	return sum
}

// Average 求平均数
func (receiver enumerable[T]) Average(fn func(item T) any) float64 {
	sum := receiver.Sum(fn)
	count := len(*receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}

// AverageItem 求平均数
func (receiver enumerable[T]) AverageItem() float64 {
	sum := receiver.Sum(func(item T) any { return item })
	count := len(*receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}

// Min 获取最小值
func (receiver enumerable[T]) Min(fn func(item T) any) any {
	lst := *receiver.source

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
func (receiver enumerable[T]) MinItem() T {
	lst := *receiver.source

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
func (receiver enumerable[T]) Max(fn func(item T) any) any {
	lst := *receiver.source

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
func (receiver enumerable[T]) MaxItem() T {
	lst := *receiver.source

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
func (receiver enumerable[T]) GroupBy(mapSlice any, getMapKeyFunc func(item T) any) {
	mapSliceVal := reflect.ValueOf(mapSlice).Elem()
	if mapSliceVal.Kind() != reflect.Map {
		panic("mapSlice入参必须为map类型")
	}

	// make....
	mapSliceVal.Set(reflect.MakeMap(mapSliceVal.Type()))

	for _, item := range *receiver.source {
		// 生成key
		key := reflect.ValueOf(getMapKeyFunc(item))
		// 根据key，找到map的value
		findMapValue := mapSliceVal.MapIndex(key)
		// nil说明map不存在这个key
		if findMapValue == reflect.ValueOf(nil) {
			findMapValue = reflect.MakeSlice(mapSliceVal.Type().Elem(), 0, 0)
			mapSliceVal.SetMapIndex(key, findMapValue)
		}
		mapValue := reflect.Append(findMapValue, reflect.ValueOf(item))
		mapSliceVal.SetMapIndex(key, mapValue)
	}
}

// OrderBy 正序排序，fn 返回的是要排序的字段的值
func (receiver enumerable[T]) OrderBy(fn func(item T) any) enumerable[T] {
	lst := *receiver.source

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

	return enumerable[T]{source: &lst}
}

// OrderByItem 正序排序，fn 返回的是要排序的字段的值
func (receiver enumerable[T]) OrderByItem() enumerable[T] {
	lst := *receiver.source

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

	return enumerable[T]{source: &lst}
}

// OrderByDescending 倒序排序，fn 返回的是要排序的字段的值
func (receiver enumerable[T]) OrderByDescending(fn func(item T) any) enumerable[T] {
	lst := *receiver.source

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

	return enumerable[T]{source: &lst}
}

// OrderByDescendingItem 倒序排序，fn 返回的是要排序的字段的值
func (receiver enumerable[T]) OrderByDescendingItem() enumerable[T] {
	lst := *receiver.source

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
	return enumerable[T]{source: &lst}
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
func (receiver enumerable[T]) Select(sliceOrList any, fn func(item T) any) {
	sliceOrListVal := reflect.ValueOf(sliceOrList).Elem()
	// 切片类型
	if sliceOrListVal.Kind() == reflect.Slice {
		var lst = make([]reflect.Value, 0)
		for _, item := range *receiver.source {
			lst = append(lst, reflect.ValueOf(fn(item)))
		}

		value := reflect.Append(sliceOrListVal, lst...)
		sliceOrListVal.Set(value)
		return
	}
	if strings.HasPrefix(sliceOrListVal.Type().String(), "collections.List[") {
		// 初始化
		value := reflect.New(sliceOrListVal.Type())
		value.MethodByName("New").Call(nil)

		for _, item := range *receiver.source {
			itemValue := reflect.ValueOf(fn(item))
			value.MethodByName("Add").Call([]reflect.Value{itemValue})
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
func (receiver enumerable[T]) SelectMany(sliceOrList any, fn func(item T) any) {
	sliceOrListVal := reflect.ValueOf(sliceOrList).Elem()

	// 切片类型
	if sliceOrListVal.Kind() == reflect.Slice {
		value := reflect.MakeSlice(sliceOrListVal.Type(), 0, 0)
		for _, item := range *receiver.source {
			itemValue := reflect.ValueOf(fn(item))
			if itemValue.Type() != sliceOrListVal.Type() {
				panic("arrSlice入参类型必须与fn返回的类型一致")
			}
			value = reflect.AppendSlice(value, itemValue)
		}
		sliceOrListVal.Set(value)
		return
	}

	if strings.HasPrefix(sliceOrListVal.Type().String(), "collections.List[") {
		// 初始化
		value := reflect.New(sliceOrListVal.Type())
		value.MethodByName("New").Call(nil)

		for _, item := range *receiver.source {
			itemValue := reflect.ValueOf(fn(item))
			value.MethodByName("Add").CallSlice([]reflect.Value{itemValue})
		}
		sliceOrListVal.Set(value.Elem())
		return
	}

	panic("sliceOrList入参必须为切片或collections.List类型")
}

// ToMap 转成字典
func (receiver enumerable[T]) ToMap(mapSlice any, getMapKeyFunc func(item T) any, getMapValueFunc func(item T) any) {
	mapSliceVal := reflect.ValueOf(mapSlice).Elem()
	if mapSliceVal.Kind() != reflect.Map {
		panic("mapSlice入参必须为map类型")
	}

	// make....
	mapSliceVal.Set(reflect.MakeMap(mapSliceVal.Type()))

	for _, item := range *receiver.source {
		// 生成key
		key := reflect.ValueOf(getMapKeyFunc(item))
		// 根据key，找到map的value
		findMapValue := mapSliceVal.MapIndex(key)
		// nil说明map不存在这个key
		if findMapValue == reflect.ValueOf(nil) {
			findMapValue = reflect.MakeSlice(mapSliceVal.Type().Elem(), 0, 0)
			mapSliceVal.SetMapIndex(key, findMapValue)
		}
		mapValue := reflect.Append(findMapValue, reflect.ValueOf(getMapValueFunc(item)))
		mapSliceVal.SetMapIndex(key, mapValue)
	}
}

// ToList 返回List集合
func (receiver enumerable[T]) ToList() List[T] {
	return NewList(*receiver.source...)
}

// ToArray 转成数组
func (receiver enumerable[T]) ToArray() []T {
	return *receiver.source
}

// ToPageList 数组分页
func (receiver enumerable[T]) ToPageList(pageSize int, pageIndex int) PageList[T] {
	pageList := PageList[T]{
		RecordCount: int64(len(*receiver.source)),
	}

	if pageSize < 1 {
		pageSize = 10
	}

	// 计算总页数
	var allCurrentPage int64 = 0
	// 没有设置pageIndex，则按take返回
	if pageIndex < 1 {
		take := receiver.Take(pageSize)
		pageList.List = take.ToList()
		return pageList
	}

	allCurrentPage = pageList.RecordCount / int64(pageSize)
	if pageList.RecordCount%int64(pageSize) != 0 {
		allCurrentPage++
	}
	if allCurrentPage == 0 {
		allCurrentPage = 1
	}

	if int64(pageIndex) > allCurrentPage {
		pageIndex = int(allCurrentPage)
	}
	skipCount := pageSize * (pageIndex - 1)
	lst := (*receiver.source)[skipCount : skipCount+pageSize]
	pageList.List = NewList(lst...)
	return pageList
}

// MapToList 类型转换，比如List[PO] 转换为 List[DO]
// toList：必须为List类型
func (receiver enumerable[T]) MapToList(toList any) {
	toValue := reflect.ValueOf(toList).Elem()
	// 传进来的，可能不是struct，而是通过反射创建的any
	if toValue.Kind() == reflect.Ptr || toValue.Kind() == reflect.Interface {
		toValue = toValue.Elem()
	}
	if !strings.HasPrefix(toValue.Type().String(), "collections.List[") {
		panic("要转换的类型，必须也是collections.List集合")
	}

	// 拿到数组类型后，先mapper到数组
	destToArrayType := toValue.MethodByName("ToArray").Type().Out(0)
	destArr := reflect.New(destToArrayType).Interface()
	_ = mapper.MapperSlice(receiver.ToArray(), destArr)

	newValue := reflect.New(toValue.Type())
	// 初始化集合
	newValue.MethodByName("New").Call(nil)
	// 将数组添加到集合
	newValue.MethodByName("Add").CallSlice([]reflect.Value{reflect.ValueOf(destArr).Elem()})
	reflect.ValueOf(toList).Elem().Set(newValue.Elem())
}

// ToListAny 转成ListAny
func (receiver enumerable[T]) ToListAny() ListAny {
	array := receiver.ToArray()
	lst := NewListAny()
	for _, item := range array {
		lst.Add(item)
	}
	return lst
}

// MapToArray 类型转换，比如List[PO] 转换为 []DO
// toSlice：必须为切片类型
func (receiver enumerable[T]) MapToArray(toSlice any) {
	toValue := reflect.ValueOf(toSlice)
	toType := toValue.Type()
	if toType.Elem().Kind() != reflect.Slice {
		panic("要转换的类型，必须是切片类型")
	}
	destArr := reflect.New(toType.Elem()).Interface()
	_ = mapper.MapperSlice(receiver.ToArray(), destArr)

	toValue.Elem().Set(reflect.ValueOf(destArr).Elem())
}

// Empty 返回一个新的Empty集合
func (receiver enumerable[T]) Empty() enumerable[T] {
	return enumerable[T]{source: &[]T{}}
}

// Intersect 两个集合的交集（共同存在的元素）
func (receiver enumerable[T]) Intersect(lstRight List[T]) enumerable[T] {
	var lst []T
	for _, item := range *receiver.source {
		if lstRight.Contains(item) {
			lst = append(lst, item)
		}
	}
	return enumerable[T]{source: &lst}
}

// Concat 合并两个集合
func (receiver enumerable[T]) Concat(lstRight List[T]) enumerable[T] {
	lst := append(*receiver.source, *lstRight.source...)
	return enumerable[T]{source: &lst}
}

// Union 合并两个集合，并去重
func (receiver enumerable[T]) Union(lstRight List[T]) enumerable[T] {
	union := receiver.Concat(lstRight)
	return union.Distinct()
}

// Distinct 集合去重
func (receiver enumerable[T]) Distinct() enumerable[T] {
	lst := NewList[T]()
	for _, item := range *receiver.source {
		if !lst.Contains(item) {
			lst.Add(item)
		}
	}
	return lst.enumerable
}

// Except 移除参数中包含的集合元素
func (receiver enumerable[T]) Except(lstRight List[T]) enumerable[T] {
	lst := receiver.ToList()
	for _, item := range *lstRight.source {
		lst.Remove(item)
	}
	return lst.enumerable
}
