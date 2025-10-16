package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 10
	wg.Add(numGoroutines)

	increment := func() {

		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}
	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("The final counter value is : %d\n\n", counter)
}
