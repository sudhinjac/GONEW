package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan struct{})

	go func() {

		fmt.Println("Working....")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()
	<-done
	fmt.Println("finished...")
}
