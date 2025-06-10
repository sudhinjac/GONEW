package main

import "fmt"

//function as first class citizens it can be used in any ways.
func main() {

	result := applyOper(5, 3, add)
	fmt.Println("5+3 = ", result)

	muliply := createMultiplier(2)
	fmt.Println("6 * 2 = ", muliply(6))

}

func add(a, b int) int {

	return a + b
}

func applyOper(x int, y int, operation func(int, int) int) int {

	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
