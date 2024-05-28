package main

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go/v21"
)

func main() {
	filePath := "component/component.wasm"

	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	linker := wasmtime.NewLinker(engine)
	err := linker.DefineWasi()
	check(err)

	module, err := wasmtime.NewModuleFromFile(engine, filePath)
	check(err)

	instance, err := linker.Instantiate(store, module)
	check(err)

	fibFunc := instance.GetFunc(store, "fib:demo/fibonacci#fib")
	if fibFunc == nil {
		panic("float function isn't exported")
	}

	result, err := fibFunc.Call(store, 35)
	check(err)

	fmt.Println("result:", result.(int64))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
