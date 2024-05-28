package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/Shopify/go-lua"
)

//go:embed fib.lua
var fibFile string

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	if err := lua.DoString(l, fibFile); err != nil {
		panic(err)
	}

	l.Global("fib")
	l.PushInteger(38)

	now := time.Now()

	l.Call(1, 1)

	elapsed := time.Since(now)

	result, ok := l.ToInteger(-1)
	if !ok {
		panic("failed to get return value")
	}

	fmt.Printf("result: %d\n", result)
	fmt.Println("time took:", elapsed)
}
