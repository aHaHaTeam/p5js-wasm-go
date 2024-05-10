package rendering

import (
	"syscall/js"

	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func CreateCanvas(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("createCanvas", []int{2, 3, 4}, args...)
}

func ResizeCanvas(args ...any) error {
	return AnyFuncReturnErr("resizeCanvas", []int{2, 3}, args...)
}

func NoCanvas(args ...any) error {
	return AnyFuncReturnErr("noCanvas", []int{0}, args...)
}

func CreateGraphics(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("CreateGraphics", []int{2, 3, 4}, args...)
}

func CreateFramebuffer(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("CreateFramebuffer", []int{0, 1}, args...)
}

func ClearDepth(args ...any) error {
	return AnyFuncReturnErr("clearDepth", []int{0, 1}, args...)
}

func BlendMode(args ...any) error {
	return AnyFuncReturnErr("blendMode", []int{1}, args...)
}

func SetAttributes(args ...any) error {
	return AnyFuncReturnErr("setAttributes", []int{1, 2}, args...)
}
