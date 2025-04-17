package main

import (
	"fmt"
	"time"
)

// A job contains input data
type job struct {
	id    int
	value int
}

// A result contains output data and maybe metadata
type result struct {
	job    job
	square int
	worker int
}

func worker(id int, jobs <-chan job, results chan<- result) {
	for j := range jobs {
		// Simulate processing time
		time.Sleep(100 * time.Millisecond)
		results <- result{
			job:    j,
			square: j.value * j.value,
			worker: id,
		}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 2

	jobs := make(chan job, numJobs) // buffered so we don't block sending jobs
	results := make(chan result, numJobs)

	// Start N worker goroutines
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send all jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- job{id: i, value: i}
	}
	close(jobs) // Important! Otherwise workers block forever

	// Collect results
	for i := 0; i < numJobs; i++ {
		res := <-results
		fmt.Printf("Worker %d processed job %d: square=%d\n",
			res.worker, res.job.id, res.square)
	}
}
