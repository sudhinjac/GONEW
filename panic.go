package main

import "fmt"

func main() {
	process(10)
	process(-1)

}

func process(input int) {

	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")

	if input < 0 {
		panic("input must be non negative number")

	}
	fmt.Println("Processing input: ", input)
}
