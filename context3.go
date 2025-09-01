package main

import (
	"context"
	"fmt"
	"log"
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
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestID", "asdf123456gf")
	go doWork(ctx)
	time.Sleep(3 * time.Second)
	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("RequestID found: ", requestID)
	} else {
		fmt.Println("No requestIF found")
	}
	logwithContext(ctx, "This is a test log message")
}

func logwithContext(ctx context.Context, message string) {

	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID:  %v - %v", requestIDVal, message)
}
