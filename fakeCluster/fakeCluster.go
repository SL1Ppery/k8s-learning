package fakecluster

import "k8s-learning/resource"

type FakeCluster struct {
	Pods        []resource.Pod
	Deployments []resource.Deployment
}

var Cluster = FakeCluster{
	Pods: []resource.Pod{
		{APIVersion: "v1", Kind: "Pod", Metadata: resource.Metadata{Name: "nginx", Namespace: "default"}, Status: resource.Status{Phase: "Running"}},
		{APIVersion: "v1", Kind: "Pod", Metadata: resource.Metadata{Name: "mysql", Namespace: "default"}, Status: resource.Status{Phase: "Pending"}},
	},
	Deployments: []resource.Deployment{
		{APIVersion: "apps/v1", Kind: "Deployment", Metadata: resource.Metadata{Name: "nginx-deployment", Namespace: "default"}, Replicas: 3},
		{APIVersion: "apps/v1", Kind: "Deployment", Metadata: resource.Metadata{Name: "ruoyi", Namespace: "ruoyi"}, Replicas: 3},
	},
}

func GetFakeClusterPods() []resource.Pod {
	return Cluster.Pods
}

func GetFakeClusterDeployment() []resource.Deployment {
	return Cluster.Deployments
}
