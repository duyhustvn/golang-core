package worker

import (
	"fmt"
	"sync"
)

func processJob(workerId int, value int, wg *sync.WaitGroup, frequencies []int) {
	defer wg.Done()
	frequencies[workerId]++
	// fmt.Printf("wokerId: %d value: %d\n", workerId, value)
}

func worker(workerId int, c chan int, wg *sync.WaitGroup, frequencies []int) {
	for v := range c {
		processJob(workerId, v, wg, frequencies)
	}

}

func createWorker(numWokers int, numOfJobs int) error {
	var wg sync.WaitGroup
	intChan := make(chan int)

	frequencies := make([]int, numWokers)
	for i := 0; i < numWokers; i++ {
		go worker(i, intChan, &wg, frequencies)
	}

	wg.Add(numOfJobs)
	for i := 0; i < numOfJobs; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Wait()

	for k, v := range frequencies {
		fmt.Printf("workerId: %d  frequency: %d\n", k, v)
	}

	return nil
}

/*
10 worker, unbuffered channel, 100.000.000 job
workerId: 0  frequency:  9817070
workerId: 1  frequency: 10155585
workerId: 2  frequency:  9953559
workerId: 3  frequency: 10278190
workerId: 4  frequency:  9948626
workerId: 5  frequency: 10084709
workerId: 6  frequency:  9905748
workerId: 7  frequency: 10026454
workerId: 8  frequency:  9883850
workerId: 9  frequency:  9946209
go run main.go  20,61s user 0,76s system 147% cpu 14,453 total


10 worker, buffered channel with size 10, 100.000.000 job
workerId: 0  frequency:  9814069
workerId: 1  frequency:  9988539
workerId: 2  frequency: 10031461
workerId: 3  frequency: 10287142
workerId: 4  frequency: 10109403
workerId: 5  frequency: 10115846
workerId: 6  frequency: 10020775
workerId: 7  frequency:  9792082
workerId: 8  frequency:  9805143
workerId: 9  frequency: 10035540
go run main.go  10,35s user 0,65s system 154% cpu 7,104 total

10 worker, buffered channel with size 100, 100.000.000 job
workerId: 0  frequency: 10.086.218
workerId: 1  frequency: 10045216
workerId: 2  frequency: 9910931
workerId: 3  frequency: 9956536
workerId: 4  frequency: 10027422
workerId: 5  frequency: 9971981
workerId: 6  frequency: 10032825
workerId: 7  frequency: 10117397
workerId: 8  frequency: 9971630
workerId: 9  frequency: 9879844
go run main.go  8,85s user 0,48s system 183% cpu 5,075 total


100 worker, buffered channel with size 10, 100.000.000 job
workerId: 0  frequency: 1.039.172
workerId: 1  frequency: 980863
workerId: 2  frequency: 966531
workerId: 3  frequency: 975226
workerId: 4  frequency: 1002683
workerId: 5  frequency: 979842
workerId: 6  frequency: 983004
workerId: 7  frequency: 1010658
workerId: 8  frequency: 993530
workerId: 9  frequency: 1010710
workerId: 10  frequency: 998731
workerId: 11  frequency: 954143
workerId: 12  frequency: 1014178
workerId: 13  frequency: 1019892
workerId: 14  frequency: 998698
workerId: 15  frequency: 996905
workerId: 16  frequency: 1004317
workerId: 17  frequency: 969051
workerId: 18  frequency: 979402
workerId: 19  frequency: 996852
workerId: 20  frequency: 1014156
workerId: 21  frequency: 980944
workerId: 22  frequency: 1032738
workerId: 23  frequency: 1043311
workerId: 24  frequency: 1022820
workerId: 25  frequency: 1038586
workerId: 26  frequency: 1021887
workerId: 27  frequency: 995941
workerId: 28  frequency: 1012997
workerId: 29  frequency: 1010146
workerId: 30  frequency: 1004951
workerId: 31  frequency: 998003
workerId: 32  frequency: 1020558
workerId: 33  frequency: 1019906
workerId: 34  frequency: 989792
workerId: 35  frequency: 1011393
workerId: 36  frequency: 1002596
workerId: 37  frequency: 1013902
workerId: 38  frequency: 1038368
workerId: 39  frequency: 967685
workerId: 40  frequency: 1014852
workerId: 41  frequency: 996948
workerId: 42  frequency: 988934
workerId: 43  frequency: 990245
workerId: 44  frequency: 998909
workerId: 45  frequency: 1020434
workerId: 46  frequency: 975307
workerId: 47  frequency: 1003738
workerId: 48  frequency: 1013453
workerId: 49  frequency: 969519
workerId: 50  frequency: 973610
workerId: 51  frequency: 1017739
workerId: 52  frequency: 1016404
workerId: 53  frequency: 1015427
workerId: 54  frequency: 992312
workerId: 55  frequency: 1010718
workerId: 56  frequency: 958632
workerId: 57  frequency: 991889
workerId: 58  frequency: 1013961
workerId: 59  frequency: 1003665
workerId: 60  frequency: 981169
workerId: 61  frequency: 1011285
workerId: 62  frequency: 983619
workerId: 63  frequency: 1015137
workerId: 64  frequency: 980467
workerId: 65  frequency: 993323
workerId: 66  frequency: 978777
workerId: 67  frequency: 1003225
workerId: 68  frequency: 964351
workerId: 69  frequency: 983304
workerId: 70  frequency: 1030239
workerId: 71  frequency: 1048744
workerId: 72  frequency: 981216
workerId: 73  frequency: 1020992
workerId: 74  frequency: 978427
workerId: 75  frequency: 967237
workerId: 76  frequency: 1001898
workerId: 77  frequency: 1016089
workerId: 78  frequency: 981305
workerId: 79  frequency: 972937
workerId: 80  frequency: 956138
workerId: 81  frequency: 1008289
workerId: 82  frequency: 997935
workerId: 83  frequency: 995056
workerId: 84  frequency: 1028347
workerId: 85  frequency: 1007954
workerId: 86  frequency: 996018
workerId: 87  frequency: 978297
workerId: 88  frequency: 1067270
workerId: 89  frequency: 980726
workerId: 90  frequency: 1003087
workerId: 91  frequency: 970099
workerId: 92  frequency: 1024692
workerId: 93  frequency: 994400
workerId: 94  frequency: 966396
workerId: 95  frequency: 1002473
workerId: 96  frequency: 995577
workerId: 97  frequency: 1035159
workerId: 98  frequency: 1012930
workerId: 99  frequency: 1003652
go run main.go  100,24s user 5,25s system 460% cpu 22,892 total
*/
