package main

import (
	"fmt"
)

func main() {
	qs := make(chan bool)
	go func() {
		fmt.Println("Hello from a new goroutine")
		qs <- false
	}()
	fmt.Println("Hello from main")
	v := <-qs
	fmt.Println(v)
}

func SayHelloFromGoroutine(qs chan bool) {
	fmt.Println("Hello from a new goroutine")
	qs <- false

}
