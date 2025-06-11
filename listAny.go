package collections

type ListAny struct {
	List[any]
}

func NewListAny(source ...any) ListAny {
	var lst List[any]
	if source == nil {
		lst = NewList[any]()
	} else {
		lst = NewList(source...)
	}
	return ListAny{
		List: lst,
	}
}
