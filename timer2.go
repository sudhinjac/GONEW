package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Deleyed Operation executed")
	}()

	fmt.Println("Waiting....")
	time.Sleep(3 * time.Second)
	fmt.Println("End of Program")
}
