package main

import "fmt"

func main() {
	message := "hello world"

	for i, v := range message {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}
