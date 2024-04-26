package typography

import (
	"p5js-wasm-go/utils"
	"syscall/js"
)

func LoadFont(path any, callbacks ...func(...js.Value) any) (js.Value, error) {
	return utils.FuncWithCallbacksReturnJsValue("loadFont", []int{0, 1, 2}, path, callbacks...)
}

func Text(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("text", []int{3, 4, 5}, args)
}

func TextFont(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textFont", []int{0, 1, 2}, args)
}
