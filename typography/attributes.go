package typography

import (
	"syscall/js"

	"github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func TextAlign(args ...any) error {
	return utils.AnyFuncReturnErr("textAlign", []int{0, 1, 2}, args...)
}

func TextLeading(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textLeading", []int{0, 1}, args...)
}

func TextSize(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textSize", []int{0, 1}, args...)
}

func TextStyle(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textStyle", []int{0, 1}, args...)
}

func TextWidth(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textWidth", []int{1}, args...)
}

func TextAscent(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textAscent", []int{0}, args...)
}

func TextDescent(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textDescent", []int{0}, args...)
}

func TextWrap(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textWrap", []int{1}, args...)
}
