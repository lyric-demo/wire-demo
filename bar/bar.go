package bar

import "github.com/lyric-demo/wire-demo/foo"

type Barer interface {
	Bar() string
}

type bar struct {
	foo foo.Fooer
}

func (b bar) Bar() string {
	return b.foo.Foo() + "-bar"
}

func NewBarer(foo foo.Fooer) Barer {
	return bar{foo}
}
