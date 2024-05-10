package environment

import (
	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
	"syscall/js"
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

func GetTargetFrameRate(args ...any) error {
	return AnyFuncReturnErr("getTargetFrameRate", []int{0}, args...)
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

func WindowResized(args ...any) error {
	return AnyFuncReturnErr("windowResized", []int{0, 1}, args...)
}

func Width() js.Value {
	return GetGlobalValue("width")
}

func Height() js.Value {
	return GetGlobalValue("height")
}

func FullScreen(args ...any) error {
	return AnyFuncReturnErr("fullscreen", []int{0, 1}, args...)
}

func PixelDensity(args ...any) error {
	return AnyFuncReturnErr("pixelDensity", []int{0, 1}, args...)
}

func DisplayDensity(args ...any) error {
	return AnyFuncReturnErr("displayDensity", []int{0}, args...)
}

func GetURL(args ...any) error {
	return AnyFuncReturnErr("getURL", []int{0}, args...)
}

func GetURLPath(args ...any) error {
	return AnyFuncReturnErr("getURLPath", []int{0}, args...)
}

func GetURLParams(args ...any) error {
	return AnyFuncReturnErr("getURLParams", []int{0}, args...)
}
