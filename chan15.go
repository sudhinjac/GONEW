package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	close(ch)
	val, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed")
		return
	}

	fmt.Println(val)
	time.Sleep(time.Second * 1)
}
