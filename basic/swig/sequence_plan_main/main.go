package main

import (
	"SWIG-GO-TEST/basic/swig/sequence_plan"
	"fmt"
)

func main() {
	// 调用 C++ 中的 sequence_plan 函数
	path := "/home/bf/CLionProjects/sample_data/starwars.parquet"
	result := sequence_plan.Sequence_plan(path)

	if result == 0 {
		fmt.Println("Success")
	} else {
		fmt.Printf("Failed with error code %d\n", result)
	}
}
