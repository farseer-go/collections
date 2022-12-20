package collections

import "reflect"

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

// MapToPageList 类型转换，如PageList[PO]转PageList[DO]
func (receiver PageList[TData]) MapToPageList(pageList any) {
	pageListValue := reflect.ValueOf(pageList).Elem()
	newValue := reflect.New(pageListValue.Type()).Elem()
	// 设置总记录数
	newValue.FieldByName("RecordCount").SetInt(receiver.RecordCount)

	// 得到目标List的any对象
	newList := newValue.FieldByName("List").Interface()
	// 将原r.List转换到目标List对象
	receiver.List.MapToList(&newList)
	// 转换后将newList赋值给目标对象的List
	newValue.FieldByName("List").Set(reflect.ValueOf(newList))
	pageListValue.Set(newValue)
}
