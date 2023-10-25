package main

import "testing"

/*
Run: BenchmarkFibRecursion10

	go test -bench= for benchmark

Result:

	goos: linux
	goarch: amd64
	pkg: golangcore/benchmark
	cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
	BenchmarkFibRecursion10-8        5183991               229.8 ns/op

Explain:

			BenchmarkFibRecursion10-8        5183991               229.8 ns/op
			BenchmarkFibRecursion10: name of benchmark to run
			8: Number of cores/CPUs the benchmark is running on
			5183991: the number of iterations performed by the benchmark
		    229.8 ns/op: the time taken for each iteration or subtest. Specifically, "228.9 ns/op" indicates that each iteration took approximately 228.9 nanoseconds (ns).

	   This measurement show how long it takes for the "FibRecursion" function to execute when given an input of 10 under the specified configuration (8 cores or CPUs).
*/

var result int

func BenchmarkFibRecursion50(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = fibRecur(50)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkFibRecursionMemo50(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int)
		r = fibRecurMemo(50, memo)
	}
	result = r
}
