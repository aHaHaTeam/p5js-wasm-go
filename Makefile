
EXAMPLES_MAIN_GO = $(wildcard ./examples/*/main.go)
EXAMPLES_MAIN = $(patsubst %.go, %, $(EXAMPLES_MAIN_GO))
EXAMPLES_MAIN_WASM = $(patsubst %.go, %.wasm, $(EXAMPLES_MAIN_GO))

build: combine $(EXAMPLES_MAIN)

clean:
	echo $(EXAMPLES_MAIN)
	rm ./examples/*/main.wasm

combine:
	go run ./combine.go

$(EXAMPLES_MAIN):
	GOOS=js GOARCH=wasm go build -o ./$@.wasm ./$@.go
