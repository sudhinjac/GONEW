package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	addr      Address
}

type Address struct {
	city    string
	country string
}

func (p Person) fullname() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incage() {
	p.age++
}

func main() {
	p1 := Person{
		firstName: "Sudhin",
		lastName:  "Jacob",
		age:       47,
		addr: Address{
			city:    "Bangalore",
			country: "India",
		},
	}

	fmt.Println("First Name:", p1.firstName)
	fmt.Println("Last Name:", p1.lastName)
	fmt.Println("Age:", p1.age)
	fmt.Println("City:", p1.addr.city)
	fmt.Println("Country:", p1.addr.country)
	fmt.Println("Full Name:", p1.fullname())

	p1.incage()
	fmt.Println("After Increment of Age:", p1.age)
}
