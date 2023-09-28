package utils

type Pool struct {
	Workers  *[]Worker
	MAX_SIZE int
}

func NewPool(size int) *Pool {
	return &Pool{
		Workers:  &[]Worker{},
		MAX_SIZE: size,
	}
}

func (p *Pool) AddWorker(worker Worker) {
	if len(*p.Workers) == p.MAX_SIZE {
		return
	}
	*p.Workers = append(*p.Workers, worker)
}

func (p *Pool) Start() {
	for _, worker := range *p.Workers {
		go worker.DoWork()
	}
}
