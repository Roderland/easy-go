package easy_go

type Interface interface {
	// CompareTo return: <0 means less, =0 means equals, >0 means more
	CompareTo(e interface{}) int
}

type Integer struct {
	value int
}

func NewInteger(value int) Integer {
	return Integer{value: value}
}

func (i Integer) CompareTo(e interface{}) int {
	return i.value - e.(Integer).value
}
