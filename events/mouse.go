package events

import (
	. "p5js-wasm-go/utils"
	"syscall/js"
)

func MovedX() js.Value {
	return GetGlobalValue("movedX")
}

func MovedY() js.Value {
	return GetGlobalValue("movedY")
}

func MouseX() js.Value {
	return GetGlobalValue("mouseX")
}

func MouseY() js.Value {
	return GetGlobalValue("mouseY")
}

func PmouseX() js.Value {
	return GetGlobalValue("pmouseX")
}

func PmouseY() js.Value {
	return GetGlobalValue("pmouseY")
}

func WinMouseX() js.Value {
	return GetGlobalValue("winMouseX")
}

func WinMouseY() js.Value {
	return GetGlobalValue("winMouseY")
}

func PwinMouseX() js.Value {
	return GetGlobalValue("pwinMouseX")
}

func PwinMouseY() js.Value {
	return GetGlobalValue("pwinMouseY")
}

func MouseButton() js.Value {
	return GetGlobalValue("mouseButton")
}

func MouseIsPressed() js.Value {
	return GetGlobalValue("mouseIsPressed")
}

func MouseMoved(callback func(...js.Value) interface{}) {
	BingJsFunc("mouseMoved", callback)
}

func MouseDragged(callback func(...js.Value) interface{}) {
	BingJsFunc("mouseDragged", callback)
}

func MousePressed(callback func(...js.Value) interface{}) {
	BingJsFunc("mousePressed", callback)
}

func MouseReleased(callback func(...js.Value) interface{}) {
	BingJsFunc("mouseReleased", callback)
}

func MouseClicked(callback func(...js.Value) interface{}) {
	BingJsFunc("mouseClicked", callback)
}

func DoubleClicked(callback func(...js.Value) interface{}) {
	BingJsFunc("doubleClicked", callback)
}

func MouseWheel(callback func(...js.Value) interface{}) {
	BingJsFunc("mouseWheel", callback)
}

func RequestPointerLock(callback func(...js.Value) interface{}) {
	BingJsFunc("requestPointerLock", callback)
}

func ExitPointerLock(callback func(...js.Value) interface{}) {
	BingJsFunc("exitPointerLock", callback)
}
