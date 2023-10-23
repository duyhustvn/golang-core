package task

import "log"

type PrintTask struct {
	message string
}

func NewPrintTask(msg string) *PrintTask {
	return &PrintTask{message: msg}
}

func (p *PrintTask) Execute() error {
	log.Println(p.message)
	return nil
}
