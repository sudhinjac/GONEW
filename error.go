package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("Math: Square root of negatieve number")
	}
	result := math.Sqrt(x)
	return result, nil
}

func main() {

	//	result, err := sqrt(-1)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(result)

	err := eprocess()

	if err != nil {
		fmt.Println(err)
	}

}

type myError struct {
	message string
}

func (m *myError) Error() string {

	return fmt.Sprintf("Error: %s .", m.message)
}

func eprocess() error {

	return &myError{"Customer error message"}
}
