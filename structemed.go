package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person
	emId   string
	salary float32
}

func (p person) introduce() {

	fmt.Printf("I am %v and age is %v.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I am %s employee ID is : %s I earn %.2f.\n", e.name, e.emId, e.salary)
}

func main() {

	emp := Employee{
		person: person{name: "John", age: 30},
		emId:   "E1001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Emp ID: ", emp.emId)
	fmt.Println("Salary: ", emp.salary)
	emp.introduce() // override the first method
}
