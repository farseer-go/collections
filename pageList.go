package collections

// PageList 用于分页数组，包含总记录数
type PageList[TData any] struct {
	// 总记录数
	RecordCount int64
	// 数据列表
	List List[TData]
}

// NewPageList 数据分页列表及总数
func NewPageList[TData any](list List[TData], recordCount int64) PageList[TData] {
	return PageList[TData]{
		List:        list,
		RecordCount: recordCount,
	}
}
