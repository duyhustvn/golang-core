package main

import "time"

func notForeverLoop() {
	done := make(chan int)

	go func(don chan int) {
		time.Sleep(5 * time.Second)
		close(don)
	}(done)

	for {
		select {
		case <-done:
			return
		default:
		}
	}
}

func foreverLoop() {
	done := make(chan int)

	for {
		select {
		case <-done:
			return
		default:
		}
	}
}

func main() {
	notForeverLoop()
}
