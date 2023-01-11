package sub

type Test struct {
	A string
	B int
}

func (t Test) PrintA() {
	println(t.A)
}
