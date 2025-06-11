package main

import "fmt"

func main() {

	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sub := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(sub(1))
	fmt.Println(sub(2))
}

func adder() func() int {

	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
