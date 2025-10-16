package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 5 {
		fmt.Println(i)
		fmt.Println(time.Now())
		fmt.Println("\n")
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {

	for _, letter := range "ABCDE" {
		fmt.Println(string(letter))
		fmt.Println(time.Now())
		fmt.Println("\n")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	go printNumbers()
	go printLetters()
	time.Sleep(3 * time.Second)
}
