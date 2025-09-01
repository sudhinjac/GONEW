package main

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {

	select {
	case <-ctx.Done():
		return "operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even ", num)
		} else {
			return fmt.Sprintf("%d is odd ", num)
		}
	}

}

func main() {

	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with  context.TODO(): ", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	result = checkEvenOdd(ctx, 10)
	time.Sleep(2 * time.Second)
	fmt.Println("Result from timeout context: ", result)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result After timeout context: ", result)

}
