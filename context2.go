package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled: ", ctx.Err())
			return
		default:
			fmt.Println("Working....")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "asdf123456gf")
	go doWork(ctx)
	time.Sleep(3 * time.Second)
	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("RequestID found: ", requestID)
	} else {
		fmt.Println("No requestIF found")
	}

}
