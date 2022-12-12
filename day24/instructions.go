package day24

import "fmt"

type value struct {
	id          int
	operation   string
	left, right *value
	number      int
	min, max    int
	fwd         *value
}

func (v *value) Init() string {
	switch v.operation {
	case "num":
		return fmt.Sprint(v.number)
	case "inp":
		return fmt.Sprint("w", v.number) //, "[", v.min, ", ", v.max, "]")
	default:
		return fmt.Sprintf("(%v %v %v)", v.left.Name(), v.operation, v.right.Name())
	}
}

func (v *value) Name() string {
	return fmt.Sprintf("t%v", v.id)
}

func (v *value) String() string {
	return fmt.Sprintf("%v = % v", v.Name(), v.Init())
}
