package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	for task := range tasks {
		results <- task * 2
	}

	fmt.Printf("Worker %d finished.\n", id)
}

func main() {

	var wg sync.WaitGroup

	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)
	wg.Add(numWorkers)
	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)

	}
	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Results: ", result)
	}
	fmt.Println("All workers Finished.")

}
