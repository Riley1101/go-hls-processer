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
	// run all the jobs in the Queue
	for _, job := range w.Queue.GetJobs() {
		go func(job chan Job) {
			for {
				select {
				case <-w.Queue.ctx.Done():
					log.Print("Queue is done")
					w.Queue.RemoveJob(job)
					return
				case job := <-job:
					job.Run(w.Queue)
				}
			}
		}(job)
	}
	return true
}
