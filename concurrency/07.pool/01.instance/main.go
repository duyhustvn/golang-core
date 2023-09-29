package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int

	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024) // 1KB
			return &mem
		},
	}

	// seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWokers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWokers)

	for i := numWokers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			// time.Sleep(100)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}
