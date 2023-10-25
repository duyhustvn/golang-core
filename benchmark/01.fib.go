package main

func fibRecur(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecur(n-2) + fibRecur(n-1)
}

func fibRecurMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = fibRecurMemo(n-2, memo) + fibRecurMemo(n-1, memo)
	return memo[n]
}

// func main() {
// 	for i := 0; i <= 45; i++ {
// 		memo := make(map[int]int)
// 		diff := fibRecur(i) - fibRecurMemo(i, memo)
// 		fmt.Printf("%d: %d\n", i, diff)
// 	}
// }
