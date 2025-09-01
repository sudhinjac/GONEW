package main

import (
	"context"
	"fmt"
)

func main() {

	todoContext := context.TODO()
	contextbkg := context.Background()
	ctx := context.WithValue(todoContext, "name", "John")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))
	ctx1 := context.WithValue(contextbkg, "City", "New York")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("City"))

}
