package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()
	time.Sleep(2 * time.Second)
	for range 2 {
		select {

		case msg := <-ch1:
			fmt.Println("Recieved from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2: ", msg)
		default:
			fmt.Println("No channels ready....")
		}
	}
	fmt.Println("End of Pogram: ")
}
