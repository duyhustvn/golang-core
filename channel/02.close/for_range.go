package main

import "fmt"

/*
 * forRangeV1
 * DEADLOCK
 * Because the for range loop cannot know when the channel cInt terminate
 */
func forRangeV1() {
	cInt := make(chan int)

	go func() {
		cInt <- 1
		cInt <- 2
	}()

	for i := range cInt {
		fmt.Println(i)
	}
}

/*
 * forRangeV2
 * NO DEADLOCK
 * close() is a way to inform that the channel (cInt) will not receivce any data,
 * so the for range can know when the channel (cInt) terminate
 */
func forRangeV2() {
	cInt := make(chan int)

	go func() {
		cInt <- 1
		cInt <- 2
		close(cInt)
	}()

	for i := range cInt {
		fmt.Println(i)
	}
}

func main() {
	forRangeV2()
}
