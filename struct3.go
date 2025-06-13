package main

import (
	"fmt"
)

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	var n int
	fmt.Print("Enter the number of rectangles: ")
	fmt.Scanln(&n)

	rectangles := make([]Rectangle, n)

	for i := 0; i < n; i++ {
		var l, w float32
		fmt.Printf("Enter length and width for rectangle %d: ", i+1)
		fmt.Scanln(&l, &w)
		rectangles[i] = Rectangle{length: l, width: w}
	}

	fmt.Println("\nAreas of the rectangles:")
	for i, rect := range rectangles {
		fmt.Printf("Rectangle %d -> Length: %.2f, Width: %.2f, Area: %.2f\n",
			i+1, rect.length, rect.width, rect.Area())
	}
}
