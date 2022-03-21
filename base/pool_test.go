package base

import (
	"fmt"
	"testing"
	"time"
)

func Test_Pool(t *testing.T) {
	pool()
}

// wokers

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job", j)
		result <- j * 2
	}
}
func pool() {
	numJobs := 5
	numWorker := 3
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 0; i < numWorker; i++ {
		go worker(i, jobs, result)
	}
	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for w := 1; w < numJobs; w++ {
		<-result
		// k := <-result
		// fmt.Println("result ", " = ", k)
	}
}
