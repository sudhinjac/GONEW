package main

import (
	"fmt"
	"time"
)

func periodicTask() {
	fmt.Println("Performing Periodic Task at: ", time.Now())

}

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			periodicTask()

		}
	}

}
