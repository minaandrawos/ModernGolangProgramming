package main

import (
	"fmt"
	"time"
)

func main() {
	ic := make(chan int)
	go periodicSend(ic)
	for i := range ic {
		fmt.Println(i)
	}
	_, ok := <-ic
	fmt.Println(ok)
}

func periodicSend(ic chan int) {
	i := 0
	for i <= 3 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
