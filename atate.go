package main

import (
	"fmt"
	"time"
)

type SatefulWorker struct {
	count int
	ch    chan int
}

func (w *SatefulWorker) start() {

	go func() {
		for {

			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count: ", w.count)
			}
		}
	}()
}

func (w *SatefulWorker) send(value int) {
	w.ch <- value

}

func main() {

	stWorker := &SatefulWorker{
		ch: make(chan int),
	}
	stWorker.start()
	for i := range 5 {
		stWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
