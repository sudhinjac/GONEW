package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				count++ // ❗ shared variable modified by multiple goroutines
			}
			fmt.Println("Goroutine", id, "done")
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}

// go -race run race.go
