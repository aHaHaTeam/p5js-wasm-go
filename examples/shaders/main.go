package main

import (
	"syscall/js"

	. "github.com/aHaHaTeam/p5js-wasm-go/all"
)

func main() {
	var texcoordShader js.Value

	Preload(func(...js.Value) interface{} {
		texcoordShader, _ = LoadShader("shader.vert", "shader.frag")
		return nil
	})

	Setup(func(...js.Value) interface{} {
		CreateCanvas(WindowWidth(), WindowHeight(), "webgl")
		NoStroke()
		return nil
	})

	Draw(func(...js.Value) interface{} {
		Shader(texcoordShader)

		Rect(0, 0, Width(), Height())

		return nil
	})

	select {}
}
