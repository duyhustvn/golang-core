package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("init make chan")
	c := make(chan int, 1)
	//c := make(chan int, 3)
	go func() {
		otherRoutine(c)
		otherRoutine(c)
		otherRoutine(c)
		otherRoutine(c)
	}()

	time.Sleep(1 * time.Second)
	x := <-c
	<-c
	wg.Done()
	wg.Wait()
	fmt.Println(x)
}

func otherRoutine(c chan int) {
	fmt.Println("receiveChan")
	c <- 1
}
