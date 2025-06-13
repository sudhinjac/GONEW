package main

import "fmt"

type Rectangle struct {
	length float32
	width  float32
}

type Shape struct {
	Rectangle
}

func (r Rectangle) Area() float32 {

	return r.length * r.width
}

func (r *Rectangle) scale(factor float32) {
	r.length *= factor
	r.width *= factor
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of the Rectangle is : ", area)
	rect.scale(2)
	area = rect.Area()
	fmt.Println("Area of the Rectangle increased by factor 2 is : ", area)
	s := Shape{Rectangle{length: 23, width: 34}}
	fmt.Println("Area is using emedded struct is: ", s.Area())
}
