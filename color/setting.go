package color

import (
	. "p5js-wasm-go/utils"
)

func Background(args ...any) error {
	return AnyFuncReturnErr("background", []int{1, 2, 3, 4}, args...)
}

func Clear(args ...any) error {
	return AnyFuncReturnErr("clear", []int{0, 1, 2, 3, 4}, args...)
}

func ColorMode(args ...any) error {
	return AnyFuncReturnErr("colorMode", []int{1, 2, 4, 5}, args...)
}

func Fill(args ...any) error {
	return AnyFuncReturnErr("fill", []int{1, 2, 3, 4}, args...)
}

func NoFill(args ...any) error {
	return AnyFuncReturnErr("noFill", []int{0}, args...)
}

func NoStroke(args ...any) error {
	return AnyFuncReturnErr("noStroke", []int{0}, args...)
}

func Erase(args ...any) error {
	return AnyFuncReturnErr("erase", []int{0, 1, 2}, args...)
}

func NoErase(args ...any) error {
	return AnyFuncReturnErr("noErase", []int{0}, args...)
}
