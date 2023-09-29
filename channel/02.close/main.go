package main

import "fmt"

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

func forLoopV3() {
	cint := make(chan int)

	go func() {
		defer close(cint)
		for i := 0; i < 5; i++ {
			cint <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", <-cint)
	}
}

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

func main() {
	forLoopV4()
}
