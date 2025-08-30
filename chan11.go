package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	select {

	case msg := <-ch:
		fmt.Println("Received: ", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: ")
	}
}
