package events

import (
	. "p5js-wasm-go/utils"
	"syscall/js"
)

// MovedX returns x-delta of the mouse.
//
// MovedX tracks how many pixels the mouse moves left or right between frames.
// MovedX will have a negative value if the mouse moves left between frames and a positive value if it moves right.
// MovedX can be calculated as MouseX - PmouseX.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/movedX
func MovedX() js.Value {
	return GetGlobalValue("movedX")
}

// MovedY returns y-delta of the mouse.
//
// MovedY tracks how many pixels the mouse moves up or down between frames.
// MovedY will have a negative value if the mouse moves up between frames and a positive value if it moves down.
// MovedY can be calculated as MouseY - PmouseY.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/movedY
func MovedY() js.Value {
	return GetGlobalValue("movedY")
}

// MouseX returns x-coordinate of the mouse.
//
// In 2D mode, MouseX keeps track of the mouse's position relative to the top-left corner of the canvas.
// In WebGL mode, MouseX keeps track of the mouse's position relative to the center of the canvas.
// If touch is used instead of the mouse, then MouseX will hold the x-coordinate of the most recent touch point.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/mouseX
func MouseX() js.Value {
	return GetGlobalValue("mouseX")
}

// MouseY returns y-coordinate of the mouse.
//
// In 2D mode, MouseY keeps track of the mouse's position relative to the top-left corner of the canvas.
// In WebGL mode, MouseY keeps track of the mouse's position relative to the center of the canvas.
// If touch is used instead of the mouse, then MouseY will hold the y-coordinate of the most recent touch point.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/mouseY
func MouseY() js.Value {
	return GetGlobalValue("mouseY")
}

// PmouseX returns x-coordinate of the mouse from the previous frame.
//
// In 2D mode, PmouseX keeps track of the mouse's position relative to the top-left corner of the canvas.
// In WebGL mode, PmouseX keeps track of the mouse's position relative to the center of the canvas.
// If touch is used instead of the mouse, then PmouseX will hold the x-coordinate of the last touch point.
//
// Note: PmouseX is reset to the current MouseX value at the start of each touch event.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/pmouseX
func PmouseX() js.Value {
	return GetGlobalValue("pmouseX")
}

// PmouseY returns y-coordinate of the mouse from the previous frame.
//
// In 2D mode, PmouseY keeps track of the mouse's position relative to the top-left corner of the canvas.
// In WebGL mode, PmouseY keeps track of the mouse's position relative to the center of the canvas.
// If touch is used instead of the mouse, then PmouseY will hold the y-coordinate of the last touch point.
//
// Note: PmouseY is reset to the current MouseY value at the start of each touch event.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/pmouseY
func PmouseY() js.Value {
	return GetGlobalValue("pmouseY")
}

// WinMouseX returns x-coordinate of the mouse within the browser.
//
// WinMouseX keeps track of the mouse's position relative to the top-left corner of the browser window.
// On a touchscreen device, WinMouseX will hold the x-coordinate of the most recent touch point.
//
// Note: Use MouseX to track the mouse’s x-coordinate within the canvas.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/winMouseX
func WinMouseX() js.Value {
	return GetGlobalValue("winMouseX")
}

// WinMouseY returns y-coordinate of the mouse within the browser.
//
// WinMouseY keeps track of the mouse's position relative to the top-left corner of the browser window.
// On a touchscreen device, WinMouseY will hold the x-coordinate of the most recent touch point.
//
// Note: Use MouseY to track the mouse’s y-coordinate within the canvas.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/winMouseY
func WinMouseY() js.Value {
	return GetGlobalValue("winMouseY")
}

// PwinMouseX returns x-coordinate of the mouse within the browser from the previous frame.
//
// PwinMouseX keeps track of the mouse's position relative to the top-left corner of the browser window.
// On a touchscreen device, PwinMouseX will hold the x-coordinate of the most recent touch point.
//
// Note: PwinMouseX is reset to the current WinMouseX value at the start of each touch event.
// Note: Use PmouseX to track the mouse’s previous x-coordinate within the canvas.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/pwinMouseX
func PwinMouseX() js.Value {
	return GetGlobalValue("pwinMouseX")
}

// PwinMouseY returns y-coordinate of the mouse within the browser from the previous frame.
//
// PwinMouseY keeps track of the mouse's position relative to the top-left corner of the browser window.
// On a touchscreen device, PwinMouseY will hold the x-coordinate of the most recent touch point.
//
// Note: PwinMouseY is reset to the current WinMouseY value at the start of each touch event.
// Note: Use PmouseY to track the mouse’s previous y-coordinate within the canvas.
//
// Returns js.Value (Number)
//
// Binding -> https://p5js.org/reference/#/p5/pwinMouseY
func PwinMouseY() js.Value {
	return GetGlobalValue("pwinMouseY")
}

// MouseButton returns the last mouse button pressed.
//
// The MouseButton variable is either "left", "right", or "center", depending on which button was pressed last.
//
// Note: Different browsers may track MouseButton differently.
// See MDN for more information -> https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/buttons
//
// Returns js.Value (String).
//
// Binding -> https://p5js.org/reference/#/p5/mouseButton
func MouseButton() js.Value {
	return GetGlobalValue("mouseButton")
}

// MouseIsPressed returns if the mouse is pressed on this frame.
//
// Return value is true if the mouse is pressed and false if not.
//
// Returns js.Value (Bool)
//
// Binding -> https://p5js.org/reference/#/p5/mouseIsPressed
func MouseIsPressed() js.Value {
	return GetGlobalValue("mouseIsPressed")
}

// MouseMoved creates a function, that is called when the mouse moves AND no mouse button is pressed.
//
// MouseMoved sets a callback to run automatically when the user moves the mouse without clicking any mouse buttons.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MouseMoved is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mouseMoved
func MouseMoved(callback func(...js.Value) interface{}) {
	BindJsFunc("mouseMoved", callback)
}

// MouseDragged creates a function, that is called when the mouse moves AND any mouse button is pressed.
//
// MouseDragged sets a callback to run automatically when the user moves the mouse with clicking mouse buttons.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MouseDragged is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mouseDragged
func MouseDragged(callback func(...js.Value) interface{}) {
	BindJsFunc("mouseDragged", callback)
}

// MousePressed creates a function, that is called once when the mouse button is pressed.
//
// MousePressed sets a callback to run automatically when the user presses a mouse button.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MousePressed is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mousePressed
func MousePressed(callback func(...js.Value) interface{}) {
	BindJsFunc("mousePressed", callback)
}

// MouseReleased creates a function, that is called once when the mouse button is released.
//
// MouseReleased sets a callback to run automatically when the user releases a mouse button.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MouseReleased is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mouseReleased
func MouseReleased(callback func(...js.Value) interface{}) {
	BindJsFunc("mouseReleased", callback)
}

// MouseClicked creates a function, that is called once when the mouse button is pressed and released.
//
// MouseClicked sets a callback to run automatically when the user presses and releases a mouse button.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MouseClicked is called.
//
// Takes 1 argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mouseClicked
func MouseClicked(callback func(...js.Value) interface{}) {
	BindJsFunc("mouseClicked", callback)
}

// DoubleClicked creates a function, that is called once when the mouse button is pressed and released twice quickly.
//
// DoubleClicked sets a callback to run automatically when the user presses and releases a mouse button twice.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when DoubleClicked is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/doubleClicked
func DoubleClicked(callback func(...js.Value) interface{}) {
	BindJsFunc("doubleClicked", callback)
}

// MouseWheel creates a function, that is called when the mouse wheel moves.
//
// MouseWheel sets a callback to run automatically when the user rotates mouse wheel.
//
// Note: the mouse-related values, such as MouseX and MouseY,
// will be updated with their most recent value when MouseWheel is called.
//
// Takes one argument "callback" - type func(...js.Value) interface{}.
// callback can optionally take argument as js.Value.
// callback will be run with one argument "event" - type js.Value.
//
// Binding -> https://p5js.org/reference/#/p5/mouseWheel
func MouseWheel(callback func(...js.Value) interface{}) {
	BindJsFunc("mouseWheel", callback)
}

// RequestPointerLock locks the mouse to its current position and makes it invisible.
//
// RequestPointerLock() allows the mouse to move forever without leaving the screen.
//
// Note: Calling RequestPointerLock() locks the values of MouseX, MouseY, PmouseX, and PmouseY.
// MovedX and MovedY
// continue updating and can be used to get the distance the mouse moved since the last frame was drawn.
//
// Note:  Most browsers require an input, such as a click, before calling RequestPointerLock().
// It’s recommended to call requestPointerLock() in an event function such as DoubleClicked().
//
// Binding -> https://p5js.org/reference/#/p5/requestPointerLock
func RequestPointerLock() {
	_ = AnyFuncReturnErr("requestPointerLock", []int{0})
}

// ExitPointerLock exits a pointer lock started with RequestPointerLock.
//
// Calling exitPointerLock() resumes updating the mouse-related variables, such as MouseX and MouseY.
//
// Binding -> https://p5js.org/reference/#/p5/exitPointerLock
func ExitPointerLock() {
	_ = AnyFuncReturnErr("exitPointerLock", []int{0})
}
