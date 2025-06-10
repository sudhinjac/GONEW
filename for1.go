package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue //continue the loop skip rest of lines
		}
		fmt.Println("ODD Number is: ", i)
		if i == 7 {
			break
		}
	}

	rows := 5

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 10 {

		fmt.Println(10 - i)
	}
}
