package main

import (
	"fmt"
	"time"
)

func main() {
	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 1):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1):
		fmt.Println(v2)
	case ic <- 23:
		fmt.Println("ic received a value")
	}
}

func waitAndSend(v, i int) chan int { //will wait for i seconds before sending value v on the return channel
	retCh := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		retCh <- v
	}()
	return retCh

}
