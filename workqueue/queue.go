package workqueue

type WorkQueue struct {
	queue chan string
}

func NewWorkQueue(size int) *WorkQueue {
	return &WorkQueue{
		queue: make(chan string, size),
	}
}

func (wq *WorkQueue) Add(key string) {
	wq.queue <- key
}

func (wq *WorkQueue) Get() string {
	return <-wq.queue
}
