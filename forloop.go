package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {

		fmt.Println(i)
	}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
