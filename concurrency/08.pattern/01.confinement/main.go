package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)

		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()

		return results
	}

	res := chanOwner()

	for r := range res {
		fmt.Println(r)
	}
}
