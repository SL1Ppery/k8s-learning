package eventhandle

import (
	"k8s-learning/event"
	"k8s-learning/workqueue"
)

type EventHandle struct {
	Queue *workqueue.WorkQueue
}

func NewEventHandle() *EventHandle {
	return &EventHandle{
		Queue: workqueue.NewWorkQueue(100),
	}
}

func (eh *EventHandle) OnAdd(e *event.Event) {
	key := e.Namespace + "/" + e.Name
	eh.Queue.Add(key)
}
