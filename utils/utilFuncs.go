package utils

import (
	"fmt"
	"slices"
	"syscall/js"
)

func AnyFuncReturnErr(name string, possibleArgsNumbers []int, args ...js.Value) error {
	if !slices.Contains(possibleArgsNumbers, len(args)) {
		return fmt.Errorf("%d arguments is not supported by function %s", len(args), name)
	}
	s := make([]interface{}, len(args))
	for i, arg := range args {
		s[i] = arg
	}
	js.Global().Call(name, s...)

	return nil
}

func AnyFuncReturnJsValue(name string, possibleArgsNumbers []int, args ...js.Value) (js.Value, error) {
	if !slices.Contains(possibleArgsNumbers, len(args)) {
		return js.Null(), fmt.Errorf("%d arguments is not supported by function %s", len(args), name)
	}
	s := make([]interface{}, len(args))
	for i, arg := range args {
		s[i] = arg
	}
	return js.Global().Call(name, s...), nil
}

func AnyFuncReturnInt(name string, possibleArgsNumbers []int, args ...js.Value) (int, error) {
	res, err := AnyFuncReturnJsValue(name, possibleArgsNumbers, args...)
	if err != nil {
		return 0, err
	}
	return res.Int(), nil
}
