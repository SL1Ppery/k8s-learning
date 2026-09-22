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

var Cluster = FakeCluster{
	Pods: []resource.Pod{
		{APIVersion: "v1", Kind: "Pod", Metadata: resource.Metadata{Name: "nginx-replicaset-feu24", Namespace: "default", Labels: map[string]string{"app": "nginx"}}, Status: resource.Status{Phase: "Running"}},
		{APIVersion: "v1", Kind: "Pod", Metadata: resource.Metadata{Name: "mysql", Namespace: "default", Labels: map[string]string{"app": "mysql"}}, Status: resource.Status{Phase: "Pending"}},
	},
	Deployments: []resource.Deployment{
		{APIVersion: "apps/v1", Kind: "Deployment", Metadata: resource.Metadata{Name: "nginx-deployment", Namespace: "default"}, Replicas: 3},
		{APIVersion: "apps/v1", Kind: "Deployment", Metadata: resource.Metadata{Name: "ruoyi", Namespace: "ruoyi"}, Replicas: 3},
	},
	ReplicaSets: []resource.ReplicaSet{
		{APIVersion: "apps/v1", Kind: "ReplicaSet", Metadata: resource.Metadata{Name: "nginx-replicaset", Namespace: "default"}, Replicas: 2, Selector: map[string]string{"app": "nginx"}},
		{APIVersion: "apps/v1", Kind: "ReplicaSet", Metadata: resource.Metadata{Name: "ruoyi-replicaset", Namespace: "ruoyi"}, Replicas: 2, Selector: map[string]string{"app": "ruoyi"}},
	},
}

func NewFakeCluster() *FakeCluster {
	return &FakeCluster{}
}

func GetFakeClusterPods() []resource.Pod {
	return Cluster.Pods
}

func GetFakeClusterDeployment() []resource.Deployment {
	return Cluster.Deployments
}

func GetFakeClusterReplicaSets() []resource.ReplicaSet {
	return Cluster.ReplicaSets
}

func GetResourceByMetadata(key string) k8sinterface.Resource {
	parts := strings.SplitN(key, "/", 2)

	if len(parts) != 2 {
		return nil
	}

	namespace := parts[0]
	name := parts[1]

	for _, pod := range Cluster.Pods {
		if pod.Metadata.Namespace == namespace &&
			pod.Metadata.Name == name {
			return &pod
		}
	}

	for _, dep := range Cluster.Deployments {
		if dep.Metadata.Namespace == namespace &&
			dep.Metadata.Name == name {
			return &dep
		}
	}

	for _, rs := range Cluster.ReplicaSets {
		if rs.Metadata.Namespace == namespace &&
			rs.Metadata.Name == name {
			return &rs
		}
	}

	return nil
}
