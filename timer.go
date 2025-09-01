package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting app....")
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("waiting for timer.c")
	<-timer.C //blocking
	fmt.Println("Timer Expired")
}
