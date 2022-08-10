package collections

// List 集合
type List[T any] struct {
	source        *[]T // 集合
	list[T]            // 对集合做修改操作
	enumerable[T]      // 对集合做读操作
}

// NewList 创建集合
func NewList[T any](source ...T) List[T] {
	if source == nil {
		source = []T{}
	}
	var lst = List[T]{
		source: &source,
		list: list[T]{
			source: &source,
			collection: collection[T]{
				source: &source,
			},
		},
		enumerable: enumerable[T]{
			source: &source,
		},
	}
	return lst
}

// AsEnumerable 返回enumerable类型
func (receiver *List[T]) AsEnumerable() enumerable[T] {
	return receiver.enumerable
}
