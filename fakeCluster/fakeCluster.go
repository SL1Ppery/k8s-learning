package fakecluster

import "k8s-learning/resource"

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

func GetFakeClusterPods() []resource.Pod {
	return Cluster.Pods
}

func GetFakeClusterDeployment() []resource.Deployment {
	return Cluster.Deployments
}

func GetFakeClusterReplicaSets() []resource.ReplicaSet {
	return Cluster.ReplicaSets
}
