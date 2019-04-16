//+build wireinject

package injector

import (
	"github.com/lyric-demo/wire-demo/bar"
	"github.com/lyric-demo/wire-demo/foo"

	"github.com/google/wire"
)

func InitBarer() bar.Barer {
	wire.Build(bar.NewBarer, foo.NewFooer)
	return nil
}
