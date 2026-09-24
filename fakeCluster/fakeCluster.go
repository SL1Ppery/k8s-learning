package fakecluster

import (
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
	"strings"
)

type FakeCluster struct {
	Pods        []resource.Pod
	Deployments []resource.Deployment
	ReplicaSets []resource.ReplicaSet
}

func NewFakeCluster() *FakeCluster {
	return &FakeCluster{}
}

func (f *FakeCluster) GetFakeClusterPods() []resource.Pod {
	return f.Pods
}

func (f *FakeCluster) GetFakeClusterDeployment() []resource.Deployment {
	return f.Deployments
}

func (f *FakeCluster) GetFakeClusterReplicaSets() []resource.ReplicaSet {
	return f.ReplicaSets
}

func (f *FakeCluster) GetResourceByMetadata(key string) k8sinterface.Resource {
	parts := strings.SplitN(key, "/", 2)

	if len(parts) != 2 {
		return nil
	}

	namespace := parts[0]
	name := parts[1]

	for i := range f.Pods {
		if f.Pods[i].Metadata.Namespace == namespace &&
			f.Pods[i].Metadata.Name == name {
			return &f.Pods[i]
		}
	}

	for i := range f.Deployments {
		if f.Deployments[i].Metadata.Namespace == namespace &&
			f.Deployments[i].Metadata.Name == name {
			return &f.Deployments[i]
		}
	}

	for i := range f.ReplicaSets {
		if f.ReplicaSets[i].Metadata.Namespace == namespace &&
			f.ReplicaSets[i].Metadata.Name == name {
			return &f.ReplicaSets[i]
		}
	}

	return nil
}
func (fc *FakeCluster) GetPodsByLabels(labels map[string]string) (*[]resource.Pod, error) {
	var matchedPods []resource.Pod
	fakeClusterPods := fc.GetFakeClusterPods()
	for i := range fakeClusterPods {
		match := true
		for k, v := range labels {
			if fakeClusterPods[i].Metadata.Labels[k] != v {
				match = false
				break
			}
		}
		if match {
			matchedPods = append(matchedPods, fakeClusterPods[i])
		}
	}
	return &matchedPods, nil
}
