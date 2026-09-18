package client

import "k8s-learning/resource"

type Client interface {
	GetPod(name string) (*resource.Pod, error)
	CreatePod(pod *resource.Pod) error
	UpdatePod(pod *resource.Pod) error

	GetDeployment(name string) (*resource.Deployment, error)
	CreateDeployment(dep *resource.Deployment) error
	UpdateDeployment(dep *resource.Deployment) error

	GetReplicaSet(name string) (*resource.ReplicaSet, error)
	CreateReplicaSet(rs *resource.ReplicaSet) error
	UpdateReplicaSet(rs *resource.ReplicaSet) error
	GetPodsByLabels(labels map[string]string) (*[]resource.Pod, error)
	CreatePodWithTemplate(template *resource.PodTemplate, name string) error
}
