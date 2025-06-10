package main

import "fmt"

func main() {

	process(10)

}

func process(i int) {

	defer fmt.Println("Deferred i value: ", i)
	defer fmt.Println("First defer statement executed.")
	defer fmt.Println("second defer statement executed.")
	defer fmt.Println("third defer statement executed.")
	i++
	fmt.Println("normal execution statement.")
	fmt.Println("Value of i: ", i)

}
