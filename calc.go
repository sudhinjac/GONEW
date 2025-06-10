package main

import "fmt"

func main() {

	var a, b int = 10, 5
	var result int

	result = a + b
	fmt.Println("Addition: ", result)
	result = a - b
	fmt.Println("Subtraction is: ", result)
	result = a * b
	fmt.Println("Multiplication: ", result)
	result = a / b
	fmt.Println("Division is: ", result)
}
