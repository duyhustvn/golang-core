package main

import (
	"fmt"
	"time"
)

func leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	/*
	 the strings channel never get a value
	 so function never exit and close(channel) will never run
	*/
	doWork(nil)
	fmt.Println("Done")
}

func prevenLeak() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("dowork")
			defer fmt.Println("exit")
			defer close(terminated)
			for {
				select {
				case <-done:
					return
				case s := <-strings:
					fmt.Println(s)
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Terminating doWork goroutine")
		close(done)
	}()

	<-terminated
}

func main() {
	prevenLeak()
}
