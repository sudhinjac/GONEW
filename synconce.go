package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func ini() {

	fmt.Println("This will not repeated no matter ho many times you call this function once do")
}
func main() {

	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Routine # ", i)
			once.Do(ini)
		}()
	}
	wg.Wait()
}
