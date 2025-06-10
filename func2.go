package main

import (
	"errors"
	"fmt"
)

func main() {

	q, r := divide(10, 3)
	fmt.Println("Quotinet is: ", q, "Reminder is: ", r)

	result, err := compare(3, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Return value is: ", result)
	}

}

func divide(a, b int) (int, int) {

	quot := a / b
	rem := a % b

	return quot, rem
}

func compare(a, b int) (string, error) {

	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return " ", errors.New("Unable to compare which is greater")
	}
}
