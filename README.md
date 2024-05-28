# Benchmark Lua & Wasm for plugin application

This simple example program calculates 38th fibonacci sequence number using embedded Lua VM
and Wasm engine.

```bash
# Build the benchmark image
$ docker build -t fib-bench .

# Run it
$ docker run --rm -it fib-bench
```
