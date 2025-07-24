package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {

			data <- "Hello" + string('0'+i)
			time.Sleep(200 * time.Millisecond)
		}
		close(data)
	}()

	for value := range data {

		fmt.Println("Recieved value: ", value, ":", time.Now())
	}
}
