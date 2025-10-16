package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex
	go func() {
		mu1.Lock()
		fmt.Println(" Go routine 1  locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Go routine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Go routine 1 finished")
	}()

	go func() {
		mu1.Lock()
		fmt.Println(" Go routine 2  locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Go routine 2 locked mu2")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Go routine 2 completed")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
	// select {}
}
