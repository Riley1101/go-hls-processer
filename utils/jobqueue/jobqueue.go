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
	log.Printf("Job running: %s", j.GetName())
	err := j.Action()
	if err != nil {
		return err
	}
	j.Completed = true
	return nil
}

type Queue struct {
	name   string
	length int
	jobs   chan Job
	ctx    context.Context
	cancel context.CancelFunc
}

func (q *Queue) GetJobs() chan Job {
	return q.jobs
}

// AddJob sends job to the channel.
func (q *Queue) AddJob(job Job) {
	MAX_QUEUE_SIZE := 2

	slog.Info("Adding job to queue", q.length)
	if len(q.jobs) == MAX_QUEUE_SIZE {
		slog.Error("Queue is full")
		return
	}
	q.jobs <- job
}

func NewQueue(name string, size int) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	return &Queue{
		name:   name,
		jobs:   make(chan Job, size),
		ctx:    ctx,
		cancel: cancel,
	}
}
