package collections

type IList struct {
	List[any]
}

// NewIList 创建集合
func NewIList(source ...any) IList {
	return IList{}
}
