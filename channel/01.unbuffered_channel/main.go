package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * onlyReadInMain
 * DEADLOCK
 * define a channel in main routine and the main routine only read from that channel
 * Since this is an unbuffered channel and there is no other goroutine sending to it (cint<-),
 * the read operation (<-cint) will block indefinitely, leading to a deadlock.
 * This is because unbuffered channels require both a sender and a receiver to be ready at the same time for communication to occur.
 */
func onlyReadInMain() {
	cint := make(chan int)
	<-cint
}

/*
 * onlyReadInSubRoutineV1
 * NOT DEADLOCK
 * because the main routine will exit after 10 seconds
 */
func onlyReadInSubRoutineV1() {
	cint := make(chan int)
	go func() {
		<-cint
	}()

	time.Sleep(10 * time.Second)
}

/*
 * onlyReadInSubRoutineV2
 * DEADLOCK
 */
func onlyReadInSubRoutineV2() {
	cint := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-cint
	}()
	wg.Wait()
}

/*
 * onlyWriteInMain
 * DEADLOCK
 * Since this is an unbuffered channel and there is no other goroutine waiting to receive from it (<-cint),
 * the send operation (cint <- 24) will block indefinitely, leading to a deadlock.
 * This is because unbuffered channels require both a sender and a receiver to be ready at the same time for communication to occur.
 */
func onlyWriteInMain() {
	cint := make(chan int)
	cint <- 24
}

/*
 * onlyWriteInSubRoutineV1
 * NOT DEADLOCK
 * because the main routine will exit after 10 seconds
 */
func onlyWriteInSubRoutineV1() {
	cint := make(chan int)
	go func() {
		cint <- 24
	}()
	time.Sleep(10 * time.Second)
}

/*
 * onlyWriteInSubRoutineV2
 * DEADLOCK
 */
func onlyWriteInSubRoutineV2() {
	cint := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cint <- 24
	}()
	wg.Wait()
}

/*
 * writeBeforReadV1
 * DEADLOCK
 */
func writeBeforReadV1() {
	cint := make(chan int)
	cint <- 24
	go func() {
		fmt.Println("Subroutine")
		fmt.Println(<-cint)
	}()
}

/*
 * writeBeforReadV2
 * NOT DEADLOCK
 */
func writeBeforReadV2() {
	cint := make(chan int)
	go func() {
		cint <- 24
	}()

	fmt.Println(<-cint)
}

/*
 * readBeforeWriteV1
 * DEADLOCK
 */
func readBeforeWriteV1() {
	cint := make(chan int)
	<-cint
	go func() {
		fmt.Println("Subroutine")
		cint <- 24
	}()
}

/*
 * readBeforeWriteV2
 * NOT DEADLOCK
 */
func readBeforeWriteV2() {
	cint := make(chan int)
	go func() {
		fmt.Println(<-cint)
	}()

	cint <- 24
}
