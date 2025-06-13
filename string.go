package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello \n Go"
	message1 := "Hello \t Go"
	message2 := "Hello \r Go"
	rawMessage := `Hello \n Go`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message is: ", len(message))
	fmt.Println("Length of message is: ", len(message1))
	fmt.Println("Length of message is: ", len(message2))
	fmt.Println("Length of message is: ", len(rawMessage))
	fmt.Println("The first character in message var is: ", message[0])

	greeting := "Hello "
	name := "Alice"

	fmt.Println(greeting + name)

	for i, v := range message {
		fmt.Printf("The index is %d and  value is %c", i, v)
		fmt.Printf("%x\n", v)
		fmt.Printf("%v\n", v)
	}
	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

}
