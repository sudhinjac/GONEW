package main

import "fmt"

func main() {

	s := fmt.Sprint("hello ", " world ", 123, 456)
	fmt.Print(s)
	fmt.Println()
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s Age: %d", name, age)
	fmt.Println()
}
