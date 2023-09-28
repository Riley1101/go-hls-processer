package utils

import (
	"log"
)

type Worker struct {
	Queue *Queue
}

func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

func (w *Worker) DoWork() bool {
	for {
		select {
		case <-w.Queue.ctx.Done():
			return true
		case job := <-w.Queue.jobs:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}
		default:
			return true
		}
	}
}
