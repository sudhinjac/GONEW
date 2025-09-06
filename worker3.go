package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker ID %d started %s \n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s \n", w.ID, w.Task)

}

func main() {

	var wg sync.WaitGroup
	tasks := []string{"digging", "Laying bricks", "painting"}

	for i, task := range tasks {

		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)

	}

	wg.Wait()

	fmt.Println("Construction is finished")

}
