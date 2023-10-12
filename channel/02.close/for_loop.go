package main

import "fmt"

/*
 * forLoopV1
 * NO DEADLOCK
 */
func forLoopV1() {
	cint := make(chan int)

	go func() {
		defer close(cint)
		for i := 0; i < 5; i++ {
			cint <- i
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", <-cint)
	}
}

/*
 * forLoopV2
 * DEADLOCK
 */
func forLoopV2() {
	cint := make(chan int)

	go func() {
		// defer close(cint)
		for i := 0; i < 5; i++ {
			cint <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", <-cint)
	}
}

/*
 * forLoopV3
 * DEADLOCK
 */
func forLoopV3() {
	cint := make(chan int)

	go func() {
		defer close(cint)
		for i := 0; i < 5; i++ {
			fmt.Printf("%d ", <-cint)

		}
	}()

	for i := 0; i < 10; i++ {
		cint <- i
	}
}

/*
 * forLoopV4
 * NO DEADLOCK
 */
func forLoopV4() {
	cint := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			cint <- i
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			cint <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", <-cint)
	}
	close(cint)
}
