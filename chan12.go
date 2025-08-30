package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel Closed...")
				return

			}
			fmt.Println("Recieved: ", msg)
		}
	}
}
