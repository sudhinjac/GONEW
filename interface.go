package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	width, height float32
}
type rect1 struct {
	width, height float32
}
type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.height * r.width
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float32 {

	return 2 * (r.height + r.width)
}

func (c circle) perim() float32 {

	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float32 {

	return 2 * c.radius
}

func (r rect1) area() float32 {

	return r.height * r.width
}

func (r rect1) perim() float32 {

	return 2 * (r.height + r.width)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func printType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("Type : int")
	case string:
		fmt.Println("Type : String")
	default:
		fmt.Println("Unknown Type")

	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	fmt.Println(c.diameter())
	r1 := rect1{width: 7, height: 9}
	measure(r1)
	myPrinter(1, "john", 45.8, true)
	printType(1)
	printType("John")
	printType(true)
}
