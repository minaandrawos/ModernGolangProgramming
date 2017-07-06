package main

import (
	"fmt"
)

func main() {
	// key ==> value (key value pair)
	/*
		x := make(map[string]int)
		x["first"] = 1
		x["second"] = 2
	*/
	x := map[string]int{
		"first":  1,
		"second": 2,
	}
	fmt.Println(x["first"])
	if v, ok := x["second"]; ok {
		fmt.Println(v)
	}
	fmt.Println(x)
	delete(x, "first")
	fmt.Println(x)
}
