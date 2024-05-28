package color

import (
	"syscall/js"

	"github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func Alpha(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("alpha", []int{1}, args...)
}

func Blue(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("blue", []int{1}, args...)

}

func Brightness(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("brightness", []int{1}, args...)
}

func Color(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("color", []int{1, 2, 3, 4}, args...)
}

func Green(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("green", []int{1}, args...)
}

func Hue(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("hue", []int{1}, args...)
}

func LerpColor(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("lerp_color", []int{3}, args...)
}

func Lightness(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("lightness", []int{1}, args...)
}

func Red(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("red", []int{1}, args...)
}

func Saturation(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("saturation", []int{1}, args...)
}
