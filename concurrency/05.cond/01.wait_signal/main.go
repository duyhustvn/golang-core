// references: https://hpc-tutorials.llnl.gov/posix/waiting_and_signaling/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		fmt.Println("Send signal to wait group, and wake the wait group that is suspend")
		/*
		 * Signal() is used to signal (or wake up) another routine which is waiting on the condition variable.
		 * Signal() should be called after the mutex is locked
		 */
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("Queue Len is 2. Wait for signal")
			/*
			 * Wait() blocks its routine and waits until the routine receives a signal
			 * Wait() should be called while the mutex is locked, and it will automatically release the mutex while it waits.
			 * After the signal is received and the routine is awakened, mutex will be automatically locked by the routine.
			 * Routine will continue to work from where it is blocked.
			 * Programmer is then responsible for unlocking the mutex when the routine is finished with it
			 */
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
