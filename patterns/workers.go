package main

import (
	"fmt"
	"sync"
	"time"
)

const numberOfWorkers = 3

const numberOfJobs = 2

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//make a chanel buffer to be equal numberObJobs to ensure that only 2 jobs available at the time
	ch := make(chan int, numberOfJobs)

	///Create  5 Jobs
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			ch <- i
		}
		close(ch)

	}()

	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go processingJob(i, ch)
	}

	wg.Wait()

}

func processingJob(id int, ch chan int) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Worker %d starts job %d", id, v)
		fmt.Println()
		time.Sleep(time.Second * 4)
		fmt.Printf("Worker %d done job %d", id, v)
		fmt.Println()
	}
}
