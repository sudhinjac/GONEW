package main

import (
	"fmt"
	"time"
)

func main() {

	var err error
	fmt.Println("Begining Program.")
	go sayHello()
	fmt.Println("After SayHello Functions")
	go func() {
		err = doWork()
	}()
	go primeNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work completed successfully")
	}

}

func sayHello() {

	time.Sleep(1 * time.Second)
	fmt.Println("Hello from go Routine")
}

func primeNumbers() {

	for i := 0; i < 10; i++ {
		fmt.Println("Number: ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcdefghijk" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {

	time.Sleep(1 * time.Second)

	return fmt.Errorf(" I am the error")
}
