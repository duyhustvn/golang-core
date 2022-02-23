package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numberOfUrl := 100000
	numberOfWorker := 5

	queue := make(chan int, numberOfUrl)
	quit := make(chan int, numberOfWorker)

	wg := new(sync.WaitGroup)

	go func() {
		for i := 1; i <= numberOfUrl; i++ {
			time.Sleep(1 * time.Millisecond)
			queue <- i
		}
		for j := 0; j < numberOfWorker; j++ {
			quit <- 1
		}
	}()

	for i := 1; i <= numberOfWorker; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			startWorker(queue, quit, fmt.Sprintf("worker %d", k))
		}(i)
	}

	wg.Wait()
	fmt.Println("Finish all worker")
}

func startWorker(queue chan int, quit chan int, name string) {
	//for i := range queue {
	//	fmt.Printf("Worker %s is crawling URL %d\n", name, i)
	//	time.Sleep(1 * time.Millisecond)
	//}
	//
	for {
		select {
		case i := <-queue:
			fmt.Printf("Worker %s is crawling URL %d\n", name, i)

		case <-quit:
			fmt.Println("QUIT")
			return
		}
	}
}
