# collections集合
> 包：`"github.com/farseer-go/collections"`

> [文档：https://farseer-go.github.io/doc/](https://farseer-go.github.io/doc/)

# List
go内置的集合只有数组和切片，有时候我们需要更多高级的数据处理功能，比如：`排序、去重、交集、并集、聚合`操作时，需要编写非常多的代码才能实现。

List集合提供了`50多个数据操作`功能，可以完全代替go原生的数组类型。

# Dictionary 字典
Dictionary是字典的意思，是go内置类型map的代替者，是一种更加高级的数据操作方案

# pageList分页
PageList实质就是在List的基础上，增加了一个`RecordCount`字段，在`分页数据`的场景下一般都会使用`PageList类型`来返回

# ListAny集合
`collections.ListAny`实际就是`collections.List[any]`