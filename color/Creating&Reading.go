package color

import (
	"p5js-wasm-go/utils"
	"syscall/js"
)

func Alpha(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("alpha", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Blue(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("blue", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Brightness(args ...any) (int, error) {
	res, err := utils.AnyFuncReturnJsValue("brightness", []int{1}, args...)
	if err != nil {
		return -1, err
	}
	return res.Int(), nil
}

func Color(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("color", []int{1, 2, 3, 4}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Green(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("green", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Hue(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("hue", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func LerpColor(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("lerp_color", []int{3}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Lightness(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("lightness", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Red(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("red", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}

func Saturation(args ...any) (js.Value, error) {
	res, err := utils.AnyFuncReturnJsValue("saturation", []int{1}, args...)
	if err != nil {
		return js.Null(), err
	}
	return res, nil
}
