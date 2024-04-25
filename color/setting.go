package color

import (
	"fmt"
	"syscall/js"
)

func Background(args ...js.Value) error {
	if len(args) <= 4 {
		js.Global().Call("background", args[0], args[1], args[2])
	} else {
		return fmt.Errorf("is accepted at most four value for mode, %d given", len(args))
	}
	return nil
}
