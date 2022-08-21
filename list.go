package collections

// List 集合
type List[T any] struct {
	source        *[]T // 集合
	IList[T]           // 对集合做修改操作
	Enumerable[T]      // 对集合做读操作
}

// NewList 创建集合
func NewList[T any](source ...T) List[T] {
	var lst = List[T]{}
	lst.New()
	lst.Add(source...)
	return lst
}

// AsEnumerable 返回enumerable类型
func (receiver *List[T]) AsEnumerable() Enumerable[T] {
	return receiver.Enumerable
}

// New 返回enumerable类型
func (receiver *List[T]) New() {
	if receiver.source == nil {
		source := &[]T{}
		receiver.source = source
		receiver.IList.source = source
		receiver.IList.Collection.source = source
		receiver.Enumerable.source = source
	}
}
