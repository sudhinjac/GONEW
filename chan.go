package main

import (
	"fmt"
	"time"
)

func main() {

	greeting := make(chan string)
	greetString := "hello"
	go func() {
		greeting <- greetString
		greeting <- "world"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()
	go func() {
		receiver := <-greeting
		fmt.Println(receiver)
		receiver = <-greeting
		fmt.Println(receiver)
		for _ = range 5 {
			receiver = <-greeting
			fmt.Println(receiver)

		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")
}
