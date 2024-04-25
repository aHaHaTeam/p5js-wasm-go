package rendering

import (
	"fmt"
	"syscall/js"
)

func CreateCanvas(width, height int, args ...js.Value) (js.Value, error) {
	if len(args) == 0 {
		return js.Global().Call("createCanvas", width, height), nil
	} else if len(args) == 0 {
		return js.Global().Call("createCanvas", width, height, args[0]), nil
	} else if len(args) == 0 {
		return js.Global().Call("createCanvas", width, height, args[0], args[1]), nil
	} else {
		return js.Value{}, fmt.Errorf("is accepted at most four value for mode, %d given", len(args)+2)
	}
}
