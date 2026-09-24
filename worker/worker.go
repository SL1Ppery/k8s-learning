package worker

import (
	"k8s-learning/controller"
	"k8s-learning/fakeapi"
	"k8s-learning/workqueue"
)

type Worker struct {
	Queue             workqueue.WorkQueue
	ControllerManager controller.ControllerManager
	FakeAPIServer     fakeapi.FakeAPIServer
}

func (w *Worker) Run() {
	for {
		key := w.Queue.Get()
		w.ControllerManager.Handle(w.FakeAPIServer.Cluster, key)
	}
}
