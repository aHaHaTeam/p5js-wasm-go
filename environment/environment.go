package environment

import (
	"syscall/js"

	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func Describe(args ...any) error {
	return AnyFuncReturnErr("describe", []int{1, 2}, args...)
}

func DescribeElement(args ...any) error {
	return AnyFuncReturnErr("describeElement", []int{2, 3}, args...)
}

func TextOutput(args ...any) error {
	return AnyFuncReturnErr("textOutput", []int{0, 1}, args...)
}

func GridOutput(args ...any) error {
	return AnyFuncReturnErr("gridOutput", []int{0, 1}, args...)
}

func Print(args ...any) error {
	return AnyFuncReturnErr("print", []int{1}, args...)
}

func FrameCount() js.Value {
	return GetGlobalValue("frameCount")
}

func DeltaTime() js.Value {
	return GetGlobalValue("deltaTime")
}

func Focused() js.Value {
	return GetGlobalValue("focused")
}

func Cursor(args ...any) error {
	return AnyFuncReturnErr("cursor", []int{1, 3}, args...)
}

func FrameRate(args ...any) error {
	return AnyFuncReturnErr("frameRate", []int{0, 1}, args...)
}

func GetTargetFrameRate(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("getTargetFrameRate", []int{0}, args...)
}

func NoCursor(args ...any) error {
	return AnyFuncReturnErr("noCursor", []int{0}, args...)
}

func WebglVersion() js.Value {
	return GetGlobalValue("webglVersion")
}

func DisplayWidth() js.Value {
	return GetGlobalValue("displayWidth")
}

func DisplayHeight() js.Value {
	return GetGlobalValue("displayHeight")
}

func WindowWidth() js.Value {
	return GetGlobalValue("windowWidth")
}

func WindowHeight() js.Value {
	return GetGlobalValue("windowHeight")
}

func WindowResized(callback func(...js.Value) interface{}) {
	BindJsFunc("windowResized", callback)
}

func Width() js.Value {
	return GetGlobalValue("width")
}

func Height() js.Value {
	return GetGlobalValue("height")
}

func FullScreen(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("fullscreen", []int{0, 1}, args...)
}

func PixelDensity(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("pixelDensity", []int{0, 1}, args...)
}

func DisplayDensity(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("displayDensity", []int{0}, args...)
}

func GetURL(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("getURL", []int{0}, args...)
}

func GetURLPath(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("getURLPath", []int{0}, args...)
}

func GetURLParams(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("getURLParams", []int{0}, args...)
}
