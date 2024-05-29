#!/bin/sh

set -e

echo "LUA:"
fib-lua

echo "WASM:"
fib-wasm /build/component.wasm
