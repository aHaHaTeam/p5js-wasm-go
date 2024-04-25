package p5js

import (
	"syscall/js"
)

func init() {
	document := js.Global().Get("document")

	p5 := document.Call("createElement", "script")
	p5.Set("src", "https://cdnjs.cloudflare.com/ajax/libs/p5.js/1.4.0/p5.js")

	styles := document.Call("createElement", "style")
	styles.Set("innerHTML", `
html, body {
   margin: 0;
   padding: 0;
}
canvas {
   display: block;
}`)

	document.Get("head").Call("appendChild", p5)
	document.Get("head").Call("appendChild", styles)
}

// Setup executes the given function. It must be called only once, usually at
// the start of the program.
//
// It should be used to define initial environment properties such as screen
// size and background color.
func Setup(setup func() interface{}) {
	setupFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		initConstants()
		return setup()
	})
	js.Global().Set("setup", setupFunc)
}

// Draw continuously executes the given function until the program is stopped.
// It must be called only once, after Setup.
func Draw(draw func() interface{}) {
	drawFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return draw()
	})
	js.Global().Set("draw", drawFunc)
}
