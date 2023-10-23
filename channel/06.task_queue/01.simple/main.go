package main

import (
	"fmt"
	"golangcore/channel/06.task_queue/01.simple/task"
	taskqueue "golangcore/channel/06.task_queue/01.simple/task_queue"
	"golangcore/channel/06.task_queue/01.simple/worker"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	taskQueue := taskqueue.NewTaskQueue()
	taskQueue.Start()

	numWorkers := 10
	go func() {
		for i := 0; i < 20; i++ {
			printTask := task.NewPrintTask(fmt.Sprintf("hello %d", i))
			taskQueue.Enqueue <- printTask
		}
	}()

	// time.Sleep(1 * time.Millisecond)
	workers := make([]*worker.Worker, 10)
	for i := 0; i < numWorkers; i++ {
		workers[i] = worker.NewWorker(i, taskQueue)
		workers[i].Start()
	}

	go func() {
		for i := 21; i < 1000; i++ {
			printTask := task.NewPrintTask(fmt.Sprintf("hello %d", i))
			taskQueue.RegisterTask(printTask)
		}
	}()

	wg.Wait()
}
