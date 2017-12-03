package assertion

type Tester interface {
	Errorf(format string, args ...interface{})
}

type Assert struct {
	t Tester
}

func NewAssert(t Tester) *Assert {
	return &Assert{t}
}

func (a *Assert) Equals(expected, actual interface{}) {
	if expected != actual {
		a.t.Errorf("Expected value: %v, found %v", expected, actual)
	}
}
