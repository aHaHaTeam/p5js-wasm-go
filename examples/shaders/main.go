package main

import (
	"syscall/js"

	. "github.com/aHaHaTeam/p5js-wasm-go/3d"
	. "github.com/aHaHaTeam/p5js-wasm-go/color"
	. "github.com/aHaHaTeam/p5js-wasm-go/environment"
	. "github.com/aHaHaTeam/p5js-wasm-go/p5js"
	. "github.com/aHaHaTeam/p5js-wasm-go/rendering"
	. "github.com/aHaHaTeam/p5js-wasm-go/shape"
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
