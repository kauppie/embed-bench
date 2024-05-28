// Generated by `wit-bindgen` 0.24.0. DO NOT EDIT!
#include "fibdemo.h"
#include <stdlib.h>

// Exported Functions from `fib:demo/fibonacci`


// Canonical ABI intrinsics

__attribute__((__weak__, __export_name__("cabi_realloc")))
void *cabi_realloc(void *ptr, size_t old_size, size_t align, size_t new_size) {
  (void) old_size;
  if (new_size == 0) return (void*) align;
  void *ret = realloc(ptr, new_size);
  if (!ret) abort();
  return ret;
}

// Component Adapters

__attribute__((__export_name__("fib:demo/fibonacci#fib")))
int64_t __wasm_export_exports_fib_demo_fibonacci_fib(int64_t arg) {
  int64_t ret = exports_fib_demo_fibonacci_fib(arg);
  return ret;
}

// Ensure that the *_component_type.o object is linked in

extern void __component_type_object_force_link_fibdemo(void);
void __component_type_object_force_link_fibdemo_public_use_in_this_compilation_unit(void) {
  __component_type_object_force_link_fibdemo();
}
