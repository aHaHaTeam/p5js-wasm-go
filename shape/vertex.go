package shape

import (
	. "github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func BeginContour(args ...any) error {
	return AnyFuncReturnErr("beginContour", []int{0}, args...)
}

func BeginShape(args ...any) error {
	return AnyFuncReturnErr("beginShape", []int{0, 1}, args...)
}

func BezierVertex(args ...any) error {
	return AnyFuncReturnErr("bezierVertex", []int{6, 9}, args...)
}

func CurveVertex(args ...any) error {
	return AnyFuncReturnErr("curveVertex", []int{2, 3}, args...)
}

func EndContour(args ...any) error { return AnyFuncReturnErr("endContour", []int{0}, args...) }

func EndShape(args ...any) error {
	return AnyFuncReturnErr("endShape", []int{0, 1, 2}, args...)
}

func QuadraticVertex(args ...any) error {
	return AnyFuncReturnErr("quadraticVertex", []int{4, 6}, args...)
}

func Vertex(args ...any) error {
	return AnyFuncReturnErr("vertex", []int{2, 3, 4, 5}, args...)
}

func Normal(args ...any) error {
	return AnyFuncReturnErr("normal", []int{1, 3}, args...)
}
