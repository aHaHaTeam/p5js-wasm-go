package p5js

import (
	"p5js-wasm-go/utils"
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

func Preload(preload func(...js.Value) any) {
	utils.BindJsFunc("preload", preload)
}

// Setup executes the given function. It must be called only once, usually at
// the start of the program.
//
// It should be used to define initial environment properties such as screen
// size and background color.
func Setup(setup func(...js.Value) any) {
	utils.BindJsFunc("setup", setup)
}

// Draw continuously executes the given function until the program is stopped.
// It must be called only once, after Setup.
func Draw(draw func(...js.Value) any) {
	utils.BindJsFunc("draw", draw)
}

func Remove() error {
	return utils.AnyFuncReturnErr("remove", []int{0})
}

func NoLoop() error {
	return utils.AnyFuncReturnErr("noLoop", []int{0})
}

func Loop() error {
	return utils.AnyFuncReturnErr("loop", []int{0})
}

func IsLooping() (js.Value, error) {
	return utils.AnyFuncReturnJsValue("isLooping", []int{0})
}

func Push() error {
	return utils.AnyFuncReturnErr("push", []int{0})
}

func Pop() error {
	return utils.AnyFuncReturnErr("pop", []int{0})
}

func Redraw(args ...any) error {
	return utils.AnyFuncReturnErr("redraw", []int{0, 1}, args...)
}
