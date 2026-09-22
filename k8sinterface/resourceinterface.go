package k8sinterface

type Resource interface {
	GetKind() string
	GetName() string
	GetNamespace() string
	GetInfo()
}

type Controller interface {
	Handle(resource Resource)
}
