package _3d

import (
	"syscall/js"

	"github.com/aHaHaTeam/p5js-wasm-go/utils"
)

func LoadShader(vertPath, fragPath any, callbacks ...func(...js.Value) any) (js.Value, error) {
	return utils.FuncWithCallbacksReturnJsValue("loadShader", []int{2, 3, 4}, []any{vertPath, fragPath}, callbacks...)
}

func CreateShader(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("createShader", []int{2}, args...)
}

func CreateFilterShader(args ...any) (js.Value, error) {
	return utils.AnyFuncReturnJsValue("createFilterShader", []int{1}, args...)
}

func Shader(args ...any) error {
	return utils.AnyFuncReturnErr("shader", []int{1}, args...)
}

func ResetShader() error {
	return utils.AnyFuncReturnErr("resetShader", []int{0})
}

func Texture(args ...any) error {
	return utils.AnyFuncReturnErr("texture", []int{1}, args...)
}

func TextureMode(args ...any) error {
	return utils.AnyFuncReturnErr("textureMode", []int{1}, args...)
}

func TextureWrap(args ...any) error {
	return utils.AnyFuncReturnErr("textureWrap", []int{1, 2}, args...)
}

func NormalMaterial() error {
	return utils.AnyFuncReturnErr("normalMaterial", []int{0})
}

func AmbientMaterial(args ...any) error {
	return utils.AnyFuncReturnErr("ambientMaterial", []int{1, 3}, args...)
}

func EmissiveMaterial(args ...any) error {
	return utils.AnyFuncReturnErr("emissiveMaterial", []int{1, 3, 4}, args...)
}

func SpecularMaterial(args ...any) error {
	return utils.AnyFuncReturnErr("specularMaterial", []int{1, 2, 3, 4}, args...)
}

func Shininess(args ...any) error {
	return utils.AnyFuncReturnErr("shininess", []int{1}, args...)
}

func Metalness(args ...any) error {
	return utils.AnyFuncReturnErr("metalness", []int{1}, args...)
}
