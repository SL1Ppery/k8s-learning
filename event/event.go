package event

type EventType string

const (
	Add    EventType = "ADD"
	Update EventType = "UPDATE"
	Delete EventType = "DELETE"
)

type Event struct {
	EventType
	Name      string
	Namespace string
	Kind      string
}
