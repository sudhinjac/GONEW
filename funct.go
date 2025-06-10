package main

import "fmt"

func main() {

	sum := add(1, 2)

	fmt.Println("The sum is: ", sum)

	greet := func() {

		fmt.Println("Hello Anonymous Function")
	}

	greet()
}

func add(a, b int) int {

	return a + b
}
