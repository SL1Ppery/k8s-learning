package k8sinterface

type Resource interface {
	GetKind() string
	GetName() string
	GetInfo()
}

type Controller interface {
	Handle(resource Resource)
}
