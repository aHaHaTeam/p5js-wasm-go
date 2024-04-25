package rendering

import (
	. "p5js-wasm-go/utils"
	"syscall/js"
)

func CreateCanvas(width, height int, args ...any) (js.Value, error) {
	res, err := AnyFuncReturnJsValue("createCanvas", []int{2, 3, 4}, args...)
	if err != nil {
		return res, err
	}
	return res, nil

}
