package main

import (
	"fmt"
	"time"
)

func longRunop() {

	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunop()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation timeout")
	case <-done:
		fmt.Println("Operation completed")
	}

}
