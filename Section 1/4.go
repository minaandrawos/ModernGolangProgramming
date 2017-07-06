package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	jason := person{
		Name:    "Json S.",
		Age:     38,
		Address: "Germany",
	}
	fmt.Println(jason.Name)

}
