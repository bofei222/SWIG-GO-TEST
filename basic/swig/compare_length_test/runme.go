package main

import (
	"SWIG-GO-TEST/basic/swig/compare_length"
	"fmt"
)

func main() {
	l1 := compare_length.NewVecInt()
	l2 := compare_length.NewVecInt()
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)

	l2.Add(1)
	l2.Add(2)
	l2.Add(4)
	ret := compare_length.Compare(l1, l2)
	fmt.Println(ret)
}
