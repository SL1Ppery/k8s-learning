package resource

import (
	"fmt"
	"k8s-learning/k8sinterface"
)

type Deployment struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Replicas   int      `json:"replicas"`
}
type Deployments []Deployment

func (dep *Deployment) GetKind() string {
	return dep.Kind
}
func (dep *Deployment) GetName() string {
	return dep.Metadata.Name
}
func (dep *Deployment) GetInfo() {
	fmt.Printf("%-40s %-40s %-40d\n", dep.Metadata.Name, dep.Metadata.Namespace, dep.Replicas)
}

func NewDeployment() k8sinterface.Resource {
	return &Deployment{}
}
