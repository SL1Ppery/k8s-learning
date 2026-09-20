package resource

import (
	"fmt"
	"k8s-learning/k8sinterface"
)

type Pod struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Status     Status   `json:"status"`
}
type Metadata struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}
type Status struct {
	Phase string `json:"phase"`
}
type Pods []Pod

func (pod *Pod) GetKind() string {
	return pod.Kind
}

func (pod *Pod) GetName() string {
	return pod.Metadata.Name
}

func (pod *Pod) GetInfo() {
	fmt.Printf("%-40s %-40s %-40s\n", pod.Metadata.Name, pod.Metadata.Namespace, pod.Status.Phase)
}
func NewPod() k8sinterface.Resource {
	return &Pod{}
}
