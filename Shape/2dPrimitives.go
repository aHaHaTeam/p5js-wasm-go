package Shape

import . "p5js-wasm-go/utils"

func Arc(args ...any) error {
	err := AnyFuncReturnErr("arc", []int{6, 7, 8}, args...)
	return err
}

func Ellipse(args ...any) error {
	err := AnyFuncReturnErr("ellipse", []int{3, 4, 5}, args...)
	return err
}

func Circle(args ...any) error {
	err := AnyFuncReturnErr("circle", []int{3}, args...)
	return err
}

func Line(args ...any) error {
	err := AnyFuncReturnErr("line", []int{4, 6}, args...)
	return err
}

func Point(args ...any) error {
	err := AnyFuncReturnErr("point", []int{1, 2, 3}, args...)
	return err
}

func Quad(args ...any) error {
	err := AnyFuncReturnErr("quad", []int{8, 9, 10, 12, 13, 14}, args...)
	return err
}

func Rect(args ...any) error {
	err := AnyFuncReturnErr("rect", []int{3, 4, 5, 6, 7, 8}, args...)
	return err
}

func Square(args ...any) error {
	err := AnyFuncReturnErr("square", []int{3, 4, 5, 6, 7}, args...)
	return err
}

func Triangle(args ...any) error {
	err := AnyFuncReturnErr("triangle", []int{6}, args...)
	return err
}
