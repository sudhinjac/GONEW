package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func ticketPRocessor(requests <-chan ticketRequest, results chan<- int) {

	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of PersonID %d with total cost of %d \n", req.numTickets, req.personID, req.cost)
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {

	numRequest := 5
	price := 5

	ticketRequests := make(chan ticketRequest, numRequest)
	ticketResults := make(chan int)

	for range 3 {
		go ticketPRocessor(ticketRequests, ticketResults)
	}

	for i := range numRequest {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}

	for range numRequest {
		fmt.Printf("Ticket for PersonID %d processed successfully!\n", <-ticketResults)
	}
}
