package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{} // declare a read-only channel of empty interfaces
	var wg sync.WaitGroup

	noop := func() {
		// noop function will indeed block while waiting to receive a value from channel c
		// since there are no other Goroutines writing to this channel
		// goroutine will never exit so that we can keep a number of them
		// in memory for measurement.
		wg.Done()
		<-c // read from channel
	}

	const numGoroutines = 1e4 // 10000
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
