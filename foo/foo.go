package foo

type Fooer interface {
	Foo() string
}

type foo struct {
}

func (foo) Foo() string {
	return "foo"
}

func NewFooer() Fooer {
	return foo{}
}
