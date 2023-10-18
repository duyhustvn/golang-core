package worker

import "testing"

var err error

// go test -bench=.
func BenchmarkWorker(b *testing.B) {
	var e error
	numWorkers := 100
	numJobs := 100000000
	for i := 0; i < b.N; i++ {
		e = createWorker(numWorkers, numJobs)
	}
	err = e
}
