package factory

import "k8s-learning/k8sinterface"

type Creator func() k8sinterface.Resource

type Registry struct {
	creators map[string]Creator
}

func NewRegistry() *Registry {
	return &Registry{
		creators: make(map[string]Creator),
	}
}

func (r *Registry) Register(kind string, creator Creator) {
	r.creators[kind] = creator
}

func (r *Registry) Create(kind string) k8sinterface.Resource {
	creator, ok := r.creators[kind]
	if !ok {
		return nil
	}
	return creator()
}
