package worker

import (
	"fmt"
	"k8s-learning/workqueue"
)

type Worker struct {
	Queue workqueue.WorkQueue
}

func (w *Worker) Run() {
	for {
		key := w.Queue.Get()
		fmt.Printf("Processing key: %v\n", key)
	}
}
