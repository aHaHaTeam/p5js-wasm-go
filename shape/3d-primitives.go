package shape

import (
	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
	"syscall/js"
)

func BeginGeometry(args ...any) error {
	return AnyFuncReturnErr("beginGeometry", []int{0}, args...)
}

func EndGeometry(args ...any) error {
	return AnyFuncReturnErr("endGeometry", []int{0}, args...)
}

func BuildGeometry(callback func(...js.Value) any) (js.Value, error) {
	return FuncWithCallbacksReturnJsValue("buildGeometry", []int{1}, []any{}, callback)
}

func FreeGeometry(args ...any) error {
	return AnyFuncReturnErr("freeGeometry", []int{1}, args...)
}

func Plane(args ...any) error {
	return AnyFuncReturnErr("plane", []int{0, 1, 2, 3, 4}, args...)
}

func Box(args ...any) error {
	return AnyFuncReturnErr("box", []int{0, 1, 2, 3, 4, 5}, args...)
}

func Sphere(args ...any) error { return AnyFuncReturnErr("sphere", []int{0, 1, 2, 3}, args...) }

func Cylinder(args ...any) error {
	return AnyFuncReturnErr("cylinder", []int{0, 1, 2, 3, 4, 5, 6}, args...)
}

func Cone(args ...any) error {
	return AnyFuncReturnErr("cone", []int{0, 1, 2, 3, 4, 5}, args...)
}

func Ellipsoid(args ...any) error {
	return AnyFuncReturnErr("ellipsoid", []int{0, 1, 2, 3, 4, 5}, args...)
}

func Torus(args ...any) error {
	return AnyFuncReturnErr("torus", []int{0, 1, 2, 3, 4}, args...)
}

func P5Geometry(args []any, callback ...func(...js.Value) any) (js.Value, error) {
	return FuncWithCallbacksReturnJsValue("p5.Geometry", []int{0, 1}, args, callback...)
}
