package color

import (
	"fmt"
	"slices"
	"syscall/js"
)

func anyFuncReturnErr(name string, possibleArgsNumber []int, args ...js.Value) error {
	if !slices.Contains(possibleArgsNumber, len(args)) {
		return fmt.Errorf("%d arguments is not supported by function %s", len(args), name)
	}
	s := make([]interface{}, len(args))
	for i, arg := range args {
		s[i] = arg
	}
	js.Global().Call(name, s...)

	return nil
}

func Background(args ...js.Value) error {
	return anyFuncReturnErr("background", []int{1, 2, 3, 4}, args...)
}

func Clear(args ...js.Value) error {
	return anyFuncReturnErr("clear", []int{0, 1, 2, 3, 4}, args...)
}

func ColorMode(args ...js.Value) error {
	return anyFuncReturnErr("colorMode", []int{1, 2, 4, 5}, args...)
}

func Fill(args ...js.Value) error {
	return anyFuncReturnErr("fill", []int{1, 2, 3, 4}, args...)
}

func NoFill(args ...js.Value) error {
	return anyFuncReturnErr("noFill", []int{0}, args...)
}

func NoStroke(args ...js.Value) error {
	return anyFuncReturnErr("noStroke", []int{0}, args...)
}

func Erase(args ...js.Value) error {
	return anyFuncReturnErr("erase", []int{0, 1, 2}, args...)
}

func NoErase(args ...js.Value) error {
	return anyFuncReturnErr("noErase", []int{0}, args...)
}
