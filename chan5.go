package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {

		fmt.Println("Working....")
		ch <- 9
		time.Sleep(2 * time.Second)
		fmt.Println("sent value")
	}()

	value := <-ch
	fmt.Println(value)
	fmt.Println("finished...")
}
