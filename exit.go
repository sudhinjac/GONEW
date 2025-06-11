package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Defered the statement: ")
	fmt.Println("Starting the main function: ")

	os.Exit(1)
	fmt.Println("End of main function")

}
