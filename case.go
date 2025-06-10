package main

import "fmt"

func main() {

	fruit := "apple"
	day := "Monday"
	switch fruit {

	case "apple":
		fmt.Println("It is an Apple..")
	case "bannana":
		fmt.Println("it is a Bananna...")
	default:
		fmt.Println("Unknown Fruit.....")
	}

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is a weekday.....")
	case "Sunday", "Satuarday":
		fmt.Println("It is a weeend.....")
	default:
		fmt.Println("Unkown Day......")

	}
	number := 15
	switch {
	case number < 10:
		fmt.Println("It is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number greater than 20.")

	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Not 2.")

	}
	checkType(10)
	checkType(3.14)
	checkType("hello")
	checkType(true)
}

func checkType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("It is an integer.")
	case float64:
		fmt.Println("It is a float.")
	case string:
		fmt.Println("It is a string.")
	default:
		fmt.Println("Unkown type")

	}
}
