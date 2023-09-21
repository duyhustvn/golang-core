package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

/*
 * This program demonstrates the performance difference between using a regular mutex and a read-write mutex (RWMutex) in concurrent programming.
 * It does so by running a test scenario where multiple readers and one writer operate on a shared resource protected by either a regular mutex or an RWMutex.
 * The program measures the time it takes to complete the test for different numbers of readers
 */
func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		// producer simulate a writer
		// It repeatly locks and unlocks the provided locker
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(10)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		// observer simulate a reader
		defer wg.Done()
		l.Lock()
		l.Unlock()
	}

	test := func(count int, mutex, rwmutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)

		beginTestTime := time.Now()
		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwmutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 20, 30, 1, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "STT\tReaders\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%d\t%v\t%v\n", i+1, count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}
}
