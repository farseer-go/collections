package collections

type IList struct {
	source []any
}

func NewIList(source ...any) IList {
	return IList{
		source: source,
	}
}

func (receiver *IList) Add(item ...any) {
	receiver.source = append(receiver.source, item...)
}

func (receiver *IList) Count() int {
	return len(receiver.source)
}

func (receiver *IList) ToArray() []any {
	return receiver.source
}

func (receiver *IList) Foreach(fn func(int, *any)) {
	for i, a := range receiver.source {
		fn(i, &a)
	}
}

func (receiver *IList) IsEmpty() bool {
	return receiver.source == nil || len(receiver.source) == 0
}

func (receiver *IList) Index(index int) any {
	return receiver.source[index]
}

func (receiver *IList) RemoveAt(index int) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		receiver.source = receiver.source[1:]
	} else {
		receiver.source = append(receiver.source[:index], receiver.source[index+1:]...)
	}
}

func (receiver *IList) Insert(index int, item any) {
	if index < 0 {
		panic("index值不能小于0")
	}
	if index >= len(receiver.source) {
		panic("index值不能超出集合的长度")
	}

	if index == 0 {
		receiver.source = append([]any{item}, receiver.source...)
	} else {
		receiver.source = append(receiver.source[:index], append([]any{item}, receiver.source[index:]...)...)
	}
}

func (receiver *IList) Clear() {
	receiver.source = []any{}
}
