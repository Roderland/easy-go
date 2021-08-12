package easy_go

type Interface interface {
	Less(e interface{}) bool
}

type Integer struct {
	value int
}

func NewInteger(value int) Integer {
	return Integer{value: value}
}

func (i Integer) Less(e interface{}) bool {
	return i.value < e.(Integer).value
}
