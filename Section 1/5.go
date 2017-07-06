package main

import (
	"fmt"
)

func main() {
	// var pI *int //memory address ==> of a value of type int
	var I int = 3
	increment(I)
	fmt.Println(I)
}

func increment(I int) {
	I++
}
