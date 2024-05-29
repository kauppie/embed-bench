package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bytecodealliance/wasmtime-go/v21"
)

func main() {
	filePath := os.Args[1]

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
		panic("fibonacci function isn't exported")
	}

	now := time.Now()

	result, err := fibFunc.Call(store, 38)
	check(err)

	elapsed := time.Since(now)

	fmt.Println("result:", result.(int64))
	fmt.Println("time:", elapsed)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
