package client

import (
	"fmt"
	fakecluster "k8s-learning/fakeCluster"
	"k8s-learning/resource"
)

type FakeClient struct{}

func (fc *FakeClient) GetPod(name string) (*resource.Pod, error) {
	fakeClusterPods := fakecluster.GetFakeClusterPods()
	for i := range fakeClusterPods {
		if fakeClusterPods[i].Metadata.Name == name {
			return &fakeClusterPods[i], nil
		}
	}
	return nil, fmt.Errorf("pod not found")
}

func (fc *FakeClient) CreatePod(pod *resource.Pod) error {
	fakecluster.Cluster.Pods = append(fakecluster.Cluster.Pods, *pod)
	return nil
}

func (fc *FakeClient) UpdatePod(pod *resource.Pod) error {
	for i, p := range fakecluster.Cluster.Pods {
		if p.Metadata.Name == pod.Metadata.Name {
			fakecluster.Cluster.Pods[i] = *pod
			return nil
		}
	}
	return fmt.Errorf("pod not found")
}

func (fc *FakeClient) GetDeployment(name string) (*resource.Deployment, error) {
	fakeClusterDeployments := fakecluster.GetFakeClusterDeployment()
	for i := range fakeClusterDeployments {
		if fakeClusterDeployments[i].Metadata.Name == name {
			return &fakeClusterDeployments[i], nil
		}
	}
	return nil, fmt.Errorf("deployment not found")
}

func (fc *FakeClient) CreateDeployment(dep *resource.Deployment) error {
	fakecluster.Cluster.Deployments = append(fakecluster.Cluster.Deployments, *dep)
	return nil
}

func (fc *FakeClient) UpdateDeployment(dep *resource.Deployment) error {
	for i, d := range fakecluster.Cluster.Deployments {
		if d.Metadata.Name == dep.Metadata.Name {
			fakecluster.Cluster.Deployments[i] = *dep
			return nil
		}
	}
	return fmt.Errorf("deployment not found")
}
