package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("End of Program")
}
