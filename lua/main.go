package main

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	if err := lua.DoFile(l, "fib.lua"); err != nil {
		panic(err)
	}

	l.Global("fib")
	l.PushInteger(38)
	l.Call(1, 1)
	result, ok := l.ToInteger(-1)
	if !ok {
		panic("failed to get return value")
	}

	fmt.Printf("result: %d\n", result)
}
