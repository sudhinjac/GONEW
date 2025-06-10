package main

import "fmt"

func main() {

	fmt.Println("Sum of 1,2,3,4,5: ", sum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}

	tot := sum(numbers...)
	fmt.Println("The total of slice is ", tot)

}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
