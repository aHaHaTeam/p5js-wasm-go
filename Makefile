build:
	GOOS=js GOARCH=wasm go build -o ./examples/main.wasm ./examples/*.go

clean:
	rm ./examples/main.wasm
