package main

import "fmt"

func main() {

	num := 1

	for num <= 10 {

		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("ODD Number is: ", num)
		num++
	}

}
