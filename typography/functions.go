package typography

import (
	"p5js-wasm-go/utils"
	"syscall/js"
)

func LoadFont(path any, callbacks ...func(...js.Value) any) (js.Value, error) {
	args := make([]interface{}, len(callbacks))
	args[0] = js.ValueOf(path)
	for i, callback := range callbacks {
		args[i+1] = js.FuncOf(func(this js.Value, args []js.Value) any {
			return callback(args...)
		})
	}
	return js.Global().Call("loadFont", args...), nil
}

func Text(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("text", []int{3, 4, 5}, args)
}

func TextFont(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("textFont", []int{0, 1, 2}, args)
}
