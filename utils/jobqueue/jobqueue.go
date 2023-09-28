package utils

import (
	"context"
	"log"
	"log/slog"
)

type Job struct {
	Name      string
	Completed bool `default:"false"`
	Action    func() error
}

func (j Job) GetName() string {
	return j.Name
}

func (j Job) Run() error {
	slog.Info("========== Running job %s \n ======>", j.Name)
	err := j.Action()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

type Queue struct {
	MAX_SIZE int `default:"2"`
	name     string
	length   int
	jobs     chan Job
	ctx      context.Context
	cancel   context.CancelFunc
}

func (q *Queue) GetJobs() chan Job {
	return q.jobs
}

// AddJob sends job to the channel.
func (q *Queue) AddJob(job Job) {
	if len(q.jobs) == q.MAX_SIZE {
		slog.Error("Queue is full")
		return
	}
	q.jobs <- job
}

func (q *Queue) AddJobs(jobs []Job) {
	for _, job := range jobs {
		if len(q.jobs) == q.MAX_SIZE {
			slog.Error("Queue is full")
			return
		}
		q.AddJob(job)
	}
}

func NewQueue(name string, size int) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	return &Queue{
		name:     name,
		jobs:     make(chan Job, size),
		ctx:      ctx,
		cancel:   cancel,
		MAX_SIZE: size,
	}
}
