package main

import (
	"sync"
	"testing"
)

/* BenchmarkContextSwitch bench the amount of time on switch between 2 goroutines
 * go test -bench=. -cpu=1 ./concurrency/02.bench_context_switch/bench_test.go
 */
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}

	sender := func() {
		defer wg.Done()
		<-begin // block until received data from begin channel

		for i := 0; i < b.N; i++ {
			c <- token
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin // block until received data from begin channel

		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin) // when a channel is closed in Go,
	// it unblocks all Goroutines waiting for data from the channel.
	// both sender and receiver Goroutines unblock simultaneously

	wg.Wait()
}
