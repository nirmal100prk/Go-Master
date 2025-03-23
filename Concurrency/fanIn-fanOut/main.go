package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %v processed job %v \n", id, job)
		result <- job
	}

	fmt.Printf("Worker %v leaving for the day ", id)
}

func main() {
	var wg sync.WaitGroup
	var jobs = make(chan int, 10)
	var results = make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result received: ", res)
	}
}
