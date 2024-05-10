package shape

import . "github.com/aHaHaTeam/p5js-wasm-go/utils"

func Arc(args ...any) error {
	return AnyFuncReturnErr("arc", []int{6, 7, 8}, args...)
}

func Ellipse(args ...any) error {
	return AnyFuncReturnErr("ellipse", []int{3, 4, 5}, args...)
}

func Circle(args ...any) error {
	return AnyFuncReturnErr("circle", []int{3}, args...)
}

func Line(args ...any) error {
	return AnyFuncReturnErr("line", []int{4, 6}, args...)
}

func Point(args ...any) error {
	return AnyFuncReturnErr("point", []int{1, 2, 3}, args...)
}

func Quad(args ...any) error {
	return AnyFuncReturnErr("quad", []int{8, 9, 10, 12, 13, 14}, args...)
}

func Rect(args ...any) error {
	return AnyFuncReturnErr("rect", []int{3, 4, 5, 6, 7, 8}, args...)
}

func Square(args ...any) error {
	return AnyFuncReturnErr("square", []int{3, 4, 5, 6, 7}, args...)
}

func Triangle(args ...any) error {
	return AnyFuncReturnErr("triangle", []int{6}, args...)
}
