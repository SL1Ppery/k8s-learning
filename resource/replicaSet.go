package resource

import (
	"fmt"
	"k8s-learning/k8sinterface"
)

type ReplicaSet struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   Metadata          `json:"metadata"`
	Replicas   int               `json:"replicas"`
	Selector   map[string]string `json:"selector"`
	Template   PodTemplate       `json:"template"`
	Status     Status            `json:"status"`
}

type PodTemplate struct {
	Metadata Metadata `json:"metadata"`
}

type ReplicaSetStatus struct {
	Replicas      int
	ReadyReplicas int
}

type ReplicaSets []ReplicaSet

func (rs *ReplicaSet) GetKind() string {
	return rs.Kind
}
func (rs *ReplicaSet) GetName() string {
	return rs.Metadata.Name
}
func (rs *ReplicaSet) GetInfo() {
	fmt.Printf("%-40s %-40s %-40d\n", rs.Metadata.Name, rs.Metadata.Namespace, rs.Replicas)
}
func NewReplicaSet() k8sinterface.Resource {
	return &ReplicaSet{}
}
