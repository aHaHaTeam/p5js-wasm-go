package color

import (
	. "p5js-wasm-go/utils"
	"syscall/js"
)

func Background(args ...js.Value) error {
	return AnyFuncReturnErr("background", []int{1, 2, 3, 4}, args...)
}

func Clear(args ...js.Value) error {
	return AnyFuncReturnErr("clear", []int{0, 1, 2, 3, 4}, args...)
}

func ColorMode(args ...js.Value) error {
	return AnyFuncReturnErr("colorMode", []int{1, 2, 4, 5}, args...)
}

func Fill(args ...js.Value) error {
	return AnyFuncReturnErr("fill", []int{1, 2, 3, 4}, args...)
}

func NoFill(args ...js.Value) error {
	return AnyFuncReturnErr("noFill", []int{0}, args...)
}

func NoStroke(args ...js.Value) error {
	return AnyFuncReturnErr("noStroke", []int{0}, args...)
}

func Erase(args ...js.Value) error {
	return AnyFuncReturnErr("erase", []int{0, 1, 2}, args...)
}

func NoErase(args ...js.Value) error {
	return AnyFuncReturnErr("noErase", []int{0}, args...)
}
