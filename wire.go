package main

import (
	"fmt"

	"github.com/lyric-demo/wire-demo/injector"
)

func main() {
	bar := injector.InitBarer()
	fmt.Println(bar.Bar())
}
