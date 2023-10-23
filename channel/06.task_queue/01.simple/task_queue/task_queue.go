package taskqueue

import (
	"golangcore/channel/06.task_queue/01.simple/task"
)

type TaskQueue struct {
	Task    []task.Task
	Enqueue chan task.Task
	Dequeue chan task.Task
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{Task: make([]task.Task, 0), Enqueue: make(chan task.Task), Dequeue: make(chan task.Task)}
}

func (tq *TaskQueue) RegisterTask(taskFn task.Task) {
	tq.Enqueue <- taskFn
}

func (tq *TaskQueue) Start() {
	// Handle add task to queue
	go func() {
		for {
			select {
			case task := <-tq.Enqueue:
				tq.Task = append(tq.Task, task)
				// log.Println("add task")
			default:
			}
		}
	}()

	// Handle get task to queue
	go func() {
		for {
			if len(tq.Task) == 0 {
				continue
			}
			select {
			case tq.Dequeue <- tq.Task[0]:
				if len(tq.Task) > 1 {
					tq.Task = tq.Task[1:]
				} else {
					tq.Task = []task.Task{}
				}
			default:
			}
		}
	}()
}
