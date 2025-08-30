package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	select {

	case msg := <-ch:
		fmt.Println("Recieved: ", msg)
	default:
		fmt.Println("No message available....")
	}

	select {
	case ch <- 1:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is not ready to recieve...")
	}
}
