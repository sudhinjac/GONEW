package main

import (
	"errors"
	"fmt"
)

func main() {

	err := doSomething1()

	if err != nil {
		fmt.Println(err)
		return
	}

}

type customError struct {
	code    int
	message string
	err     error
}

func (e *customError) Error() string {

	return fmt.Sprintf("Error: %d: %s %v\n", e.code, e.message, e.err)
}

func doSomething() error {

	return &customError{
		code:    500,
		message: "Something Went Wrong!!!",
	}
}

func doSomething1() error {

	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something Went Wrong",
			err:     err,
		}
	}
	return nil
}
func doSomethingElse() error {
	return errors.New("Internal error")
}
