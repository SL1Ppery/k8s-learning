package resource

import "fmt"

type ReplicaSet struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   Metadata          `json:"metadata"`
	Replicas   int               `json:"replicas"`
	Selector   map[string]string `json:"selector"`
	Template   PodTemplate       `json:"template"`
}

type PodTemplate struct {
	Metadata Metadata `json:"metadata"`
}

type ReplicaSets []ReplicaSet

func (rs *ReplicaSet) GetKind() string {
	return rs.Kind
}
func (rs *ReplicaSet) GetName() string {
	return rs.Metadata.Name
}
func (rs *ReplicaSet) GetInfo() {
	fmt.Printf("%-20s %-20s %-20d\n", rs.Metadata.Name, rs.Metadata.Namespace, rs.Replicas)
}
func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{}
}
