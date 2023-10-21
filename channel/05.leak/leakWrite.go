package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
leakWriteChan() the producer cannot stop
*/
func leakWriteChan() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newReadStream exited")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	for i := 0; i < 3; i++ {
		fmt.Println(<-randStream)
	}
	time.Sleep(1 * time.Second)
}

func preventLeakWriteChan() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newReadStream exited")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():

				case <-done:
					return
				}

			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	for i := 0; i < 3; i++ {
		fmt.Println(<-randStream)
	}

	close(done)
	time.Sleep(1 * time.Second)
}
