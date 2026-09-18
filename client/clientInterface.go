package client

import "k8s-learning/resource"

type Client interface {
	GetPod(name string) (*resource.Pod, error)
	CreatePod(pod *resource.Pod) error
	UpdatePod(pod *resource.Pod) error

	GetDeployment(name string) (*resource.Deployment, error)
	CreateDeployment(dep *resource.Deployment) error
	UpdateDeployment(dep *resource.Deployment) error
}
