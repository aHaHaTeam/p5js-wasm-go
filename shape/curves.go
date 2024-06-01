package shape

import (
	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
	"syscall/js"
)

func Bezier(args ...any) error {
	return AnyFuncReturnErr("bezier", []int{8, 12}, args...)
}

func BezierDetail(args ...any) error {
	return AnyFuncReturnErr("bezierDetail", []int{1}, args...)
}

func BezierPoint(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("bezierPoint", []int{5}, args...)
}

func BezierTangent(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("bezierTangent", []int{5}, args...)
}

func Curve(args ...any) error {
	return AnyFuncReturnErr("curve", []int{8, 12}, args...)
}

func CurveDetail(args ...any) error {
	return AnyFuncReturnErr("curveDetail", []int{1}, args...)
}

func CurveTightness(args ...any) error {
	return AnyFuncReturnErr("curveTightness", []int{1}, args...)
}

func CurvePoint(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("curvePoint", []int{5}, args...)
}

func CurveTangent(args ...any) (js.Value, error) {
	return AnyFuncReturnJsValue("curveTangent", []int{5}, args...)
}
