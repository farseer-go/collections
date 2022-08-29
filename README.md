# collections
Support for List collections and linq syntax

## What are the functions?

* collections
  * struct
    * PageList （用于分页数组，包含总记录数）
      * MapToPageList（类型转换，如PageList[PO]转PageList[DO]）
    * Dictionary（字典）
      * Values（获取字典的value）
      * Keys（获取字典的Keys）
      * Count（获取字典数量）
      * Add（添加元素）
      * AddMap（添加元素）
      * Clear（清除元素）
      * Remove（移除元素）
      * ContainsKey（是否存在KEY）
      * ContainsValue（是否存在指定的Value）
    * List（泛型集合）
      * AsEnumerable（返回enumerable类型）
    * IList
      * Index（获取第index索引位置的元素）
      * IndexOf（元素在集合的索引位置）
      * Insert（向第index索引位置插入元素）
      * RemoveAt（移除指定索引的元素）
      * Set（设置值）
    * Collection
      * Count（集合大小）
      * Add（添加元素）
      * Clear（清空集合）
      * Remove（移除元素）
      * RemoveAll（移除条件=true的元素）
    * Enumerable
      * Any（是否存在）
      * IsEmpty（集合是为空的）
      * First（查找符合条件的第一个元素）
      * Last（集合最后一个元素）
      * Count（集合大小）
      * Contains（是否包含元素）
      * Where（对数据进行筛选）
      * All（是否所有数据都满足fn条件）
      * Take（返回前多少条数据）
      * Skip（跳过前多少条记录）
      * Sum（求总和）
      * SumItem（求总和）
      * Average（求平均数）
      * AverageItem（求平均数）
      * Min（获取最小值）
      * MinItem（获取最小值）
      * Max（获取最大值）
      * MaxItem（获取最大值）
      * GroupBy（将数组进行分组后返回map）
      * OrderBy（正序排序）
      * OrderByItem（正序排序）
      * OrderByDescending（倒序排序）
      * OrderByDescendingItem（倒序排序）
      * Intersect（两个集合的交集）
      * Select（筛选子元素字段）
      * SelectMany（筛选子元素字段）
      * ToMap（转成字典）
      * ToList（返回List集合）
      * ToArray（转成数组）
      * ToPageList（数组分页）
      * MapToList（类型转换）
      * ToListAny（转成ListAny）
      * MapToArray（类型转换）
      * Concat（合并两个集合）
      * Union（合并两个集合，并去重）
      * Distinct（集合去重）
      * Empty（返回一个新的Empty集合）
      * Except（移除参数中包含的集合元素）
      * Range（获取切片范围）
      * RangeStart（获取切片开始位置起的范围）
      * Rand（返回随机元素）
    * func
      * NewList（创建集合）
      * NewPageList （数据分页列表及总数）
      * NewDictionary （创建一个字典）
      * NewDictionaryFromMap （创建一个字典）

## Getting Started
```go
lst := NewList[int](1, 2, 3, 4, 5, 6)
lst.Add(7)
lst.Where(func(item int) bool { return item >= 3 }).
    Where(func(item int) bool { return item >= 5 }).
    Distinct().Skip(1).Take(3).Contains(6)
```

## Add item
```go
lst := NewList[int]()
lst.Add(1, 2, 3) // lst = 1, 2, 3
lst.Insert(1, 8) // lst = 1, 8, 2, 3
lst.Count()      // return 4
```

## Remove
```go
lst := NewList[int](1, 2, 3, 4, 5, 6)
lst.RemoveAt(3)  // lst = 1, 2, 3, 5, 6
lst.Remove(5)  // lst = 1, 2, 3, 6
lst.RemoveAll(func(item int) bool {  // count = 2
    return item >= 3
})
list.Clear() // count = 0
```

## Index
```go
lst := NewList[int](1, 2, 3, 4, 5, 6)
lst.First()     // return 1
lst.Last()      // return 6
lst.Index(4)    // return item 5
lst.IndexOf(4)  // return index 3
```

## Where
```go
lst := NewList[int]()

// Any----------------------------------------------------->
lst.Any()       // return false

// IsEmpty----------------------------------------------------->
lst.IsEmpty()   // return true
lst.Add(1)
lst.Any()       // return true
lst.IsEmpty()   // return false

// Contains----------------------------------------------------->
lst.Contains(1) // return true
lst.Add(2, 3, 4, 5, 6)

// Where----------------------------------------------------->
lst = lst.Where(func(item int) bool {   // lst = 5, 6
    return item >= 3
}).Where(func(item int) bool {
    return item >= 5
}).ToList()

// All----------------------------------------------------->
lst.All(func(item int) bool { // return true
    return item == 5 || item == 6
})
```

## Get
```go
lst := NewList[int](1, 2, 3, 4, 5)
lst1 = lst.Take(3).ToList() // lst1 = 1, 2, 3
lst2 = lst.Skip(2).ToList() // lst2 = 3, 4, 5
```

## Aggregation
```go
lst := NewList[int](1, 2, 3, 4, 5)
// sum----------------------------------------------------->
lst.SumItem()                       // return 15
lst.Sum(func(item int) any {        // return 10
    return item - 1
})

// avg----------------------------------------------------->
lst.AverageItem()                   // return 3
lst.Average(func(item int) any {    // return 2
    return item - 1
})

// min----------------------------------------------------->
lst.Min(func(item int) any {        // return 0
    return item - 1
})
lst.MinItem()                       // return 1

// max----------------------------------------------------->
st.Max(func(item int) any {         // return 4
    return item - 1
})
lst.MaxItem()                       // return 5
```

## GroupBy
```go
type testItem struct {
    name string
    age  int
}
lst := NewList[testItem](testItem{name: "steden", age: 36}, testItem{name: "steden", age: 18}, testItem{name: "steden2", age: 40})
var lstMap map[string][]testItem
lst.GroupBy(&lstMap, func(item testItem) any {
    return item.name
})

len(lstMap)                         // return 2
len(lstMap["steden"]) != 2          // return 2
len(lstMap["steden2"])              // return 1
lstMap["steden"][0].age             // return 36
lstMap["steden"][1].age             // return 18
lstMap["steden2"][0].age            // return 40
```

## Order
```go
lst := NewList(3, 5, 6, 2, 1, 8, 7, 4)
// asc----------------------------------------------------->
lst.OrderByItem().ToArray()         // return 1, 2, 3, 4, 5, 6, 7, 8
lst.OrderBy(func(item int) any {    // return 1, 2, 3, 4, 5, 6, 7, 8
    return item
}).ToArray()

// desc----------------------------------------------------->
lst.OrderByDescendingItem().ToArray()       // return 8, 7, 6, 5, 4, 3, 2, 1
lst.OrderByDescending(func(item int) any {  // return 8, 7, 6, 5, 4, 3, 2, 1
    return item
}).ToArray()
```

## Join
```go
lst1 := NewList(1, 2, 3)
lst2 := NewList(3, 4, 5)
lst1.Intersect(lst2)                    // return NewList(3)
lst1.Concat(lst2)                       // return NewList(1, 2, 3, 3, 4, 5)
lst1.Union(lst2)                        // return NewList(1, 2, 3, 4, 5)
NewList(1, 2, 3, 3, 4, 5).Distinct()    // return NewList(1, 2, 3, 4, 5)
lst1.Except(lst2)                       // return NewList(1, 2)
```

## Select
```go
// []string----------------------------------------------------->
lst := NewList("1", "", "2")
var arr []string
lst.Select(&arr, func(item string) any {    // arr = {"go:1", "go:", "go:2"}
    return "go:" + item
})
// List[string]----------------------------------------------------->
var lstSelect List[string]
lst.Select(&lstSelect, func(item string) any {
    return "go:" + item
})
```
## SelectMany
```go
lst := NewList([]string{"1", "2"}, []string{"3", "4"})
var arr []string
lst.SelectMany(&arr, func(item []string) any {  // arr = "1", "2", "3", "4"
    return item
})

var lst2 List[string]
lst.SelectMany(&lst2, func(item []string) any {
    return item
})
```

## ToMap
```go
type testItem struct {
    name string
    age  int
}
lst := NewList[testItem](testItem{name: "steden", age: 36}, testItem{name: "steden", age: 18}, testItem{name: "steden2", age: 40})
var lstMap map[string][]int

lst.ToMap(&lstMap,    // ["steden"][0] = 36, ["steden"][1] = 18, ["steden2"][0] = 40
    func(key testItem) any {
        return key.name
    }, func(value testItem) any {
        return value.age
    })
```

## ToPageList
```go
lst := NewList(1, 2, 3, 4, 5, 6, 7)
item := lst.ToPageList(3, 2)     // item.RecordCount = 7    item.List = NewList(4, 5, 6)
```

## MapToList
```go
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
lst.MapToList(&lstDO) // lstDO.First().Name = "steden"   lstDO.First().Age = 37
```

## MapToArray
```go
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
lst.MapToArray(&lstDO) // lstDO[0].Name = "steden"   lstDO[0].Age = 37
```

## ToListAny
```go
type po struct {
    Name string
    Age  int
}
lst := NewList(po{Name: "steden", Age: 36}, po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 40})
lst.ToListAny()  // return collections.ListAny
```
## Range
```go
lst1 := NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
lst1.Range(3, 2)    // return lstCut[0] = 4, lstCut[1] = 5
lst1.RangeStart(7)  // return lstCut[0] = 8, lstCut[1] = 9, lstCut[2] = 10
```

## Rand
```go
NewList(1, 2, 3).Rand() // return 1 or 2 or 3
```

## Dictionary
```go
maps := make(map[string]string)
maps["name"] = "steden"
maps["age"] = "18"
dic := collections.NewDictionaryFromMap(maps)   // dic: ["name"] = "steden"，["age"] = "18"
dic.Keys()                                      // collections.List[string]: "name"、"age"
dic.Values()                                    // collections.List[string]: "steden"、"18"
dic.Count()                                     // return 2
dic.GetValue("name")                            // return "steden"
dic.Remove("name")                              // remove key name
dic.Add("name", "steden")                       // add key=name,value=steden
dic.ContainsKey("name")                         // return true
dic.ContainsValue("steden")                     // return true
dic.Clear()                                     // clear all item
dic.ToMap()                                     // mapTo map
```