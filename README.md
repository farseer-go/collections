# List集合
> 包：`"github.com/farseer-go/collections"`

- `Document`
    - [English](https://farseer-go.gitee.io/en-us/)
    - [中文](https://farseer-go.gitee.io/)
    - [English](https://farseer-go.github.io/doc/en-us/)
- Source
    - [github](https://github.com/farseer-go/fs)

![](https://img.shields.io/github/stars/farseer-go?style=social)
![](https://img.shields.io/github/license/farseer-go/collections)
![](https://img.shields.io/github/go-mod/go-version/farseer-go/collections)
![](https://img.shields.io/github/v/release/farseer-go/collections)
![Codecov](https://img.shields.io/codecov/c/github/farseer-go/collections)
![](https://img.shields.io/github/languages/code-size/farseer-go/collections)
![](https://img.shields.io/github/directory-file-count/farseer-go/collections)
![](https://goreportcard.com/badge/github.com/farseer-go/collections)

go内置的集合只有数组和切片，有时候我们需要更多高级的数据处理功能，比如：`排序、去重、交集、并集、聚合`操作时，需要编写非常多的代码才能实现。

# List
go内置的集合只有数组和切片，有时候我们需要更多高级的数据处理功能，比如：`排序、去重、交集、并集、聚合`操作时，需要编写非常多的代码才能实现。

List集合提供了`50多个数据操作`功能，可以完全代替go原生的数组类型。

# Dictionary 字典
Dictionary是字典的意思，是go内置类型map的代替者，是一种更加高级的数据操作方案

# pageList分页
PageList实质就是在List的基础上，增加了一个`RecordCount`字段，在`分页数据`的场景下一般都会使用`PageList类型`来返回

# ListAny集合
`collections.ListAny`实际就是`collections.List[any]`