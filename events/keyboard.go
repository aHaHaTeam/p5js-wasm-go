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
	BingJsFunc("keyPressed", callback)
}

func KeyReleased(callback func(...js.Value) interface{}) {
	BingJsFunc("keyReleased", callback)
}

func KeyTyped(callback func(...js.Value) interface{}) {
	BingJsFunc("keyTyped", callback)
}

func KeyIsDown(code js.Value) (js.Value, error) {
	return AnyFuncReturnJsValue("keyIsDown", []int{1}, code)
}
