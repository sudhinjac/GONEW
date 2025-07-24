package main

import (
	"fmt"
	"time"
)

func main() {
	numGoRoutines := 3
	done := make(chan int, 3)

	for i := range numGoRoutines {

		go func(id int) {
			fmt.Printf("Go Routine %d working .... \n", id)
			time.Sleep(1 * time.Second)
			done <- id
		}(i)
	}

	for range numGoRoutines {
		<-done

	}

	fmt.Println("All go routines are finsihed")
}
