package pkg

import (
	"fmt"
	"log"
	"sync"
)

func AllocateJobs(noOfJobs int) {
	for i := 0; i <= noOfJobs; i++ {
		Jobs <- Job{i + 1}
	}
	close(Jobs)
}

func worker(wg *sync.WaitGroup) {
	for job := range Jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		Results <- *result
	}
	wg.Done()
}

func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(Results)
}

func GetResults(done chan bool) {
	for result := range Results {
		if result.Num != 0 {
			fmt.Printf("Retrieving issue #%d\n", result.Num)
			ResultCollection = append(ResultCollection, result)
		}
	}
	done <- true
}
