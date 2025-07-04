package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello Go"
	fmt.Println(len(str))
	str1 := "hello"
	str2 := "wolrd."
	result := str1 + "  " + str2
	fmt.Println(result)
	fmt.Println(str[0])
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))
	fruits := "apple,orange,banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)
	countries := []string{"Germany", "Britian", "Alaska"}
	joined := strings.Join(countries, ",")
	fmt.Println(joined)
	fmt.Println(strings.Contains(str, "Go?"))
	fmt.Println(strings.Repeat("Foo ", 5))
	fmt.Println(strings.ToUpper(str1))
	str5 := "hello, 123 Go!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString("..")
	builder.WriteString("wolrd!!")
	builder.WriteString(".  .")
	result1 := builder.String()
	fmt.Println(result1)
}
