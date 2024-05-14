package shape

import . "github.com/aHaHaTeam/p5js-wasm-go/utils"

func NoSmooth(args ...any) error {
	return AnyFuncReturnErr("noSmooth", []int{0}, args...)
}

func EllipseMode(args ...any) error {
	return AnyFuncReturnErr("ellipseMode", []int{1}, args...)
}

func RectMode(args ...any) error {
	return AnyFuncReturnErr("rectMode", []int{1}, args...)
}

func Smooth(args ...any) error {
	return AnyFuncReturnErr("smooth", []int{0}, args...)
}

func StrokeCap(args ...any) error {
	return AnyFuncReturnErr("strokeCap", []int{1}, args...)
}

func StrokeJoin(args ...any) error {
	return AnyFuncReturnErr("strokeJoin", []int{1}, args...)
}

func StrokeWeight(args ...any) error {
	return AnyFuncReturnErr("strokeWeight", []int{1}, args...)
}
