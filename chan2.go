package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		// time.Sleep(2 * time.Second)
		fmt.Println("Recieved: ", <-ch)
	}()
	fmt.Println("Blocking Starts:")
	ch <- 3
	fmt.Println("Blocking Ends:")
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("Buffered Channels")
}
