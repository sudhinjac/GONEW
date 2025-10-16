package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) increment() {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) getValue() int {

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {

	var wg sync.WaitGroup

	counter := &counter{}

	numGoroutines := 10

	// wg.Add(numGoroutines)

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()

			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value is: %d\n\n", counter.getValue())
}
