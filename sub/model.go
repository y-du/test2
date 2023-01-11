package sub

type Test struct {
	A string
}

func (t Test) PrintA() {
	println(t.A)
}
