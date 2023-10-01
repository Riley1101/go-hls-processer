package utils

import (
	"context"
	"fmt"
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

func (j Job) Run(queue *Queue) error {
	msg := fmt.Sprintf("\n <========== Running job %s ======> \n", j.Name)
	slog.Info(msg)
	err := j.Action()
	j.Completed = true
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

type Queue struct {
	MAX_SIZE int `default:"4"`
	name     string
	length   int
	// arrays of chan jobs
	jobs   []chan Job
	ctx    context.Context
	cancel context.CancelFunc
}

func (q *Queue) GetJobs() []chan Job {
	return q.jobs
}

func (q *Queue) isFull() bool {
	return q.length >= q.MAX_SIZE
}

func (q *Queue) AddJob(job Job) {
	q.jobs = append(q.jobs, make(chan Job))
	q.length++
	go func() {
		q.jobs[len(q.jobs)-1] <- job
	}()
}

func (q *Queue) AddJobs(jobs []Job) {
	for _, job := range jobs {
		q.AddJob(job)
	}
}
func (q *Queue) GetCompleteJobs() int {
	fmt.Println("Getting complete jobs")
	count := 0
	for _, job := range q.jobs {
		val, ok := <-job
		if ok {
			if val.Completed {
				count++
			}
		}
	}
	return count
}
func (q *Queue) RemoveJob(job chan Job) {
	for i, j := range q.jobs {
		if j == job {
			q.jobs = append(q.jobs[:i], q.jobs[i+1:]...)
		}
	}
	q.length--
}

func NewQueue(name string, size int) *Queue {
	ctx, cancel := context.WithCancel(context.Background())
	return &Queue{
		name:     name,
		jobs:     []chan Job{},
		ctx:      ctx,
		cancel:   cancel,
		MAX_SIZE: size,
	}
}
