build:
	GOOS=js GOARCH=wasm go build -o ./examples/events/main.wasm ./examples/events/*.go

clean:
	rm ./examples/*/main.wasm

example/%:
	GOOS=js GOARCH=wasm go build -o ./examples/$*/main.wasm ./examples/$*/*.go
