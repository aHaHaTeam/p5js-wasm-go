package utils

import (
	"fmt"
	"slices"
	"syscall/js"
)

func AnyFuncReturnErr(name string, possibleArgsNumbers []int, args ...any) error {
	if !slices.Contains(possibleArgsNumbers, len(args)) {
		return fmt.Errorf("%d arguments is not supported by function %s", len(args), name)
	}
	s := make([]interface{}, len(args))
	for i, arg := range args {
		s[i] = js.ValueOf(arg)
	}
	js.Global().Call(name, s...)

	return nil
}

func AnyFuncReturnJsValue(name string, possibleArgsNumbers []int, args ...any) (js.Value, error) {
	if !slices.Contains(possibleArgsNumbers, len(args)) {
		return js.Null(), fmt.Errorf("%d arguments is not supported by function %s", len(args), name)
	}
	s := make([]interface{}, len(args))
	for i, arg := range args {
		s[i] = js.ValueOf(arg)
	}
	return js.Global().Call(name, s...), nil
}

func GetGlobalValue(name string) js.Value {
	return js.Global().Get(name)
}

func BingJsFunc(name string, goFunc func(...js.Value) any) {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		return goFunc(args...)
	})
	js.Global().Set(name, jsFunc)
}

func AnyFuncReturnInt(name string, possibleArgsNumbers []int, args ...any) (int, error) {
	res, err := AnyFuncReturnJsValue(name, possibleArgsNumbers, args...)
	if err != nil {
		return 0, err
	}
	return res.Int(), nil
}
