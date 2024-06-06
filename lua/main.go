package main

import (
	_ "embed"
	"fmt"
	"time"

	lua "github.com/yuin/gopher-lua"
)

//go:embed fib.lua
var fibFile string

func main() {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoString(fibFile); err != nil {
		panic(err)
	}

	now := time.Now()

	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("fib"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(38)); err != nil {
		panic(err)
	}

	elapsed := time.Since(now)

	luaResult := l.Get(-1)
	result := int64(lua.LVAsNumber(luaResult))

	fmt.Printf("result: %v\n", result)
	fmt.Println("time:", elapsed)
}
