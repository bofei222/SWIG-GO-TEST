package main

import (
	"SWIG-GO-TEST/basic/swig/example"
	"fmt"
)

func main() {

	// 调用 gcd 函数
	gcd := example.Gcd(48, 18)
	fmt.Println("GCD of 48 and 18 is:", gcd)
}
