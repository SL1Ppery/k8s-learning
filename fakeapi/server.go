package fakeapi

import (
	"fmt"
	fakecluster "k8s-learning/fakeCluster"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type FakeAPIServer struct {
	Cluster *fakecluster.FakeCluster
}

type APIServer interface {
	Create(r k8sinterface.Resource) error
	Get(kind, namespace, name string) (k8sinterface.Resource, error)
	Update(r k8sinterface.Resource) error
	Delete(kind, namespace, name string) error
}

func NewFakeAPIServer() *FakeAPIServer {
	return &FakeAPIServer{Cluster: fakecluster.NewFakeCluster()}
}

func (s *FakeAPIServer) Create(r k8sinterface.Resource) error {
	kind := r.GetKind()
	switch kind {
	case "Pod":
		r, ok := r.(*resource.Pod)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Pod")
		}
		namespace := r.Metadata.Namespace
		name := r.Metadata.Name
		for _, pod := range s.Cluster.Pods {
			if pod.Metadata.Namespace == namespace && pod.Metadata.Name == name {
				return fmt.Errorf("the pod %s has already existed", pod.Metadata.Name)
			}
		}
		s.Cluster.Pods = append(s.Cluster.Pods, *r)
		return nil
	case "Deployment":
		r, ok := r.(*resource.Deployment)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Deployment")
		}
		namespace := r.Metadata.Namespace
		name := r.Metadata.Name
		for _, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == name && dep.Metadata.Namespace == namespace {
				return fmt.Errorf("the deployment %s has already existed", dep.Metadata.Name)
			}
		}
		s.Cluster.Deployments = append(s.Cluster.Deployments, *r)
		return nil
	case "ReplicaSet":
		r, ok := r.(*resource.ReplicaSet)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected: replicaset")
		}
		namespace := r.Metadata.Namespace
		name := r.Metadata.Name
		for _, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == name && rep.Metadata.Namespace == namespace {
				return fmt.Errorf("the ReplicaSet %s has already existed", rep.Metadata.Name)
			}
		}
		s.Cluster.ReplicaSets = append(s.Cluster.ReplicaSets, *r)
		return nil
	}
	return fmt.Errorf("resource type not found")
}

func (s *FakeAPIServer) Get(kind, namespace, name string) (k8sinterface.Resource, error) {
	switch kind {
	case "Pod":
		for i, pod := range s.Cluster.Pods {
			if pod.Metadata.Name == name && pod.Metadata.Namespace == namespace {
				return &s.Cluster.Pods[i], nil
			}
		}
		return nil, fmt.Errorf("the pod %s not existed", name)
	case "Deployment":
		for i, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == name && dep.Metadata.Namespace == namespace {
				return &s.Cluster.Deployments[i], nil
			}
		}
		return nil, fmt.Errorf("the deployment %s not existed", name)
	case "ReplicaSet":
		for i, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == name && rep.Metadata.Namespace == namespace {
				return &s.Cluster.ReplicaSets[i], nil
			}
		}
		return nil, fmt.Errorf("the replicaset %s not existed", name)
	default:
		return nil, fmt.Errorf("the resource type not found")
	}
}

func (s *FakeAPIServer) Update(r k8sinterface.Resource) error {
	switch r.GetKind() {
	case "Pod":
		for i, pod := range s.Cluster.Pods {
			desiredPod, ok := r.(*resource.Pod)
			if !ok {
				return fmt.Errorf("resource type mismatch: expected Pod")
			}
			if pod.Metadata.Name == r.GetName() && pod.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.Pods[i] = *desiredPod
				return nil
			}
		}
		return fmt.Errorf("the pod %s not found", r.GetName())
	case "Deployment":
		for i, dep := range s.Cluster.Deployments {
			desiredDep, ok := r.(*resource.Deployment)
			if !ok {
				return fmt.Errorf("resource type mismatch: expected Deployment")
			}
			if dep.Metadata.Name == r.GetName() && dep.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.Deployments[i] = *desiredDep
				return nil
			}
		}
		return fmt.Errorf("the Deployment %s not found", r.GetName())
	case "ReplicaSet":
		for i, rep := range s.Cluster.ReplicaSets {
			desiredRep, ok := r.(*resource.ReplicaSet)
			if !ok {
				return fmt.Errorf("resource type mismatch: expected: replicaset")
			}
			if rep.Metadata.Name == r.GetName() && rep.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.ReplicaSets[i] = *desiredRep
				return nil
			}
		}
		return fmt.Errorf("the ReplicaSet %s not found", r.GetName())
	default:
		return fmt.Errorf("the resource type not found")
	}

}
func (s *FakeAPIServer) Delete(kind, namespace, name string) error {
	switch kind {
	case "Pod":
		for i, pod := range s.Cluster.Pods {
			if pod.Metadata.Name == name && pod.Metadata.Namespace == namespace {
				s.Cluster.Pods = append(s.Cluster.Pods[:i], s.Cluster.Pods[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("the Pod %s not existed", name)
	case "Deployment":
		for i, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == name && dep.Metadata.Namespace == namespace {
				s.Cluster.Deployments = append(s.Cluster.Deployments[:i], s.Cluster.Deployments[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("the Deployment %s not existed", name)
	case "ReplicaSet":
		for i, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == name && rep.Metadata.Namespace == namespace {
				s.Cluster.ReplicaSets = append(s.Cluster.ReplicaSets[:i], s.Cluster.ReplicaSets[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("the ReplicaSet %s not existed", name)
	default:
		return fmt.Errorf("the resource type not found ")
	}
}
