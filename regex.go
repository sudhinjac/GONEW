package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said \"I am Great !\"")

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_email"

	fmt.Println("Email :", re.MatchString(email1))
	fmt.Println("Email :", re.MatchString(email2))
	re = regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)
	date := "2025-06-25"
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	text := "Today's date is 2025-06-21."

	match := re.FindStringSubmatch(text)
	if len(match) > 0 {
		fmt.Println("Full Match:", match[0]) // 2025-06-21
		fmt.Println("Year:", match[1])       // 2025
		fmt.Println("Month:", match[2])      // 06
		fmt.Println("Day:", match[3])        // 21
	} else {
		fmt.Println("No match found.")
	}

}
