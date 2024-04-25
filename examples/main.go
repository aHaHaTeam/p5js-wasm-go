package main

import (
	. "p5js-wasm-go/Shape"
	. "p5js-wasm-go/color"
	"p5js-wasm-go/events"
	. "p5js-wasm-go/p5js"
	. "p5js-wasm-go/rendering"
	_ "p5js-wasm-go/rendering"
	"syscall/js"
	_ "syscall/js"
)

func main() {
	r := 20.0
	c := 0
	vars := [][]any{{239, 30, 30}, {30, 239, 30}, {30, 30, 239}}
	_ = vars
	Setup(func() interface{} {
		_, err := CreateCanvas(400, 400)
		if err != nil {
			return nil
		}
		err = Background(100)
		if err != nil {
			return nil
		}
		return nil
	})

	Draw(func() interface{} {
		err := Ellipse(events.MouseX(), events.MouseY(), r, r)
		if err != nil {
			return nil
		}
		return nil
	})

	events.MouseMoved(func(args ...js.Value) interface{} {
		_ = NoStroke()
		_ = Fill(vars[c%len(vars)]...)
		return nil
	})

	events.MouseClicked(func(args ...js.Value) interface{} {
		c = c + 1
		return nil
	})

	events.MouseDragged(func(args ...js.Value) interface{} {
		_ = Stroke(255, 255, 255)
		_ = Fill(255, 255, 255, 63)
		_ = Ellipse(events.MouseX(), events.MouseY(), r, r)
		return nil
	})

	events.MouseWheel(func(args ...js.Value) interface{} {
		if args[0].Get("delta").Int() > 0 {
			r *= 1.1
		} else {
			r /= 1.1
		}
		return nil
	})
	select {}
}
