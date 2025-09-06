package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)

	for tick := range ticker.C {
		fmt.Println("Tick at: ", tick)
	}
}
