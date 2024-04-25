package main

import (
	. "p5js-wasm-go/color"
	. "p5js-wasm-go/p5js"
	. "p5js-wasm-go/rendering"
	"syscall/js"
)

func main() {
	Setup(func() interface{} {
		_, _ = CreateCanvas(400, 400)
		return nil
	})

	Draw(func() interface{} {
		_ = Background(js.ValueOf(0), js.ValueOf(200), js.ValueOf(200))
		return nil
	})

	select {}
}
