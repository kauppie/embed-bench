#!/bin/sh

echo "LUA:"
fib-lua

echo "WASM:"
fib-wasm /build/component.wasm
