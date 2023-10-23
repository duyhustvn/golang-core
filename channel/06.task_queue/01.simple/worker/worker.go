package worker

import (
	taskqueue "golangcore/channel/06.task_queue/01.simple/task_queue"
	"log"
)

type Worker struct {
	id        int
	taskQueue *taskqueue.TaskQueue
}

func NewWorker(id int, taskQueue *taskqueue.TaskQueue) *Worker {
	return &Worker{id: id, taskQueue: taskQueue}
}

func (w *Worker) Start() {
	go func() {
		for {
			task := <-w.taskQueue.Dequeue
			if task == nil {
				log.Println("error task is nill")
				continue
			}
			if err := task.Execute(); err != nil {
				log.Printf("workerId: %d: error in executing task %+v", w.id, err)
			}
		}
	}()
}
