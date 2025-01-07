package main

/*
#cgo LDFLAGS: -L. -llib3
#include "example.h"
*/
import "C"
import "fmt"

func main() {
	// 调用 C++ 的 gcd 函数
	result := C._wrap_gcd_example_6a00cb6ce0344e18(48, 18)
	fmt.Printf("The GCD of 48 and 18 is: %d\n", result)
}
