package events

import (
	. "p5js-wasm-go/utils"
	"syscall/js"
)

func KeyIsPressed() js.Value {
	return GetGlobalValue("keyIsPressed")
}

func Key() js.Value {
	return GetGlobalValue("key")
}

func KeyCode() js.Value {
	return GetGlobalValue("keyCode")
}

func KeyPressed(callback func(...js.Value) interface{}) {
	BindJsFunc("keyPressed", callback)
}

func KeyReleased(callback func(...js.Value) interface{}) {
	BindJsFunc("keyReleased", callback)
}

func KeyTyped(callback func(...js.Value) interface{}) {
	BindJsFunc("keyTyped", callback)
}

func KeyIsDown(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("keyIsDown", []int{1}, args...)
}
