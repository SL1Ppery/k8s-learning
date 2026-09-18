package client

import (
	"fmt"
	fakecluster "k8s-learning/fakeCluster"
	nameGenerator "k8s-learning/nameGenerator"
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

func (fc *FakeClient) GetPodsByLabels(labels map[string]string) (*[]resource.Pod, error) {
	var matchedPods []resource.Pod
	fakeClusterPods := fakecluster.GetFakeClusterPods()
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

func (fc *FakeClient) CreatePodWithTemplate(template *resource.PodTemplate, name string) error {
	var newPod resource.Pod
	newPod.Metadata = template.Metadata
	newPod.Metadata.Name = nameGenerator.GenerateName(name)
	fakecluster.Cluster.Pods = append(fakecluster.Cluster.Pods, newPod)
	if fakecluster.Cluster.Pods[len(fakecluster.Cluster.Pods)-1].Metadata.Name == newPod.Metadata.Name {
		return nil
	} else {
		return fmt.Errorf("failed to create pod with template")
	}
}

// func (fc *FakeClient) GetReplicaSet(name string) (*resource.ReplicaSet, error) {
// 	fakeClusterReplicaSets := fakecluster.GetFakeClusterReplicaSets()
// 	for i := range fakeClusterReplicaSets {
// 		if fakeClusterReplicaSets[i].Metadata.Name == name {
// 			return &fakeClusterReplicaSets[i], nil
// 		}
// 	}
// 	return nil, fmt.Errorf("replica set not found")
// }

// func (fc *FakeClient) CreateReplicaSet(rs *resource.ReplicaSet) error {
// 	fakecluster.Cluster.ReplicaSets = append(fakecluster.Cluster.ReplicaSets, *rs)
// 	return nil
// }

// func (fc *FakeClient) UpdateReplicaSet(rs *resource.ReplicaSet) error {
// 	for i, r := range fakecluster.Cluster.ReplicaSets {
// 		if r.Metadata.Name == rs.Metadata.Name {
// 			fakecluster.Cluster.ReplicaSets[i] = *rs
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("replica set not found")
// }
