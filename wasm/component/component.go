package main

import (
	wit "component.com/gen"
)

func main() {}

func init() {
	var fib Fibonacci
	wit.SetExportsFibDemoFibonacci(&fib)
}

type Fibonacci struct{}

func (f *Fibonacci) Fib(x int64) int64 {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return f.Fib(x-1) + f.Fib(x-2)
}
