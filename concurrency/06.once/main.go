package main

import (
	"fmt"
	"sync"
)

/*
 * Once is an object that will perform exactly one action.
 * if once.Do(f) is called multiple times, only the first call will invoke f,
 * even if f has a different value in each invocation. A new instance of
 * Once is required for each function to execute.
 */

func main() {
	var count int
	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count) // => count = 1

	for i := 0; i < 10; i++ {
		once.Do(increment)
	}

	fmt.Printf("Count is %d\n", count) // => count = 1

	once.Do(decrement)
	fmt.Printf("Count is %d\n", count) // => count = 1
}
