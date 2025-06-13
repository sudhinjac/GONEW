package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p := Person{
		firstName: "Sudhin",
		lastName:  "Jacob",
		age:       48,
	}

	fmt.Println("First Name:", p.firstName)
	fmt.Println("First Name:", p.lastName)
	fmt.Println("First Name:", p.age)
}
