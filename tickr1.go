package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	i := 1
	for range 5 {
		fmt.Println("Tick at: ", i)
		i *= 2
	}
}
