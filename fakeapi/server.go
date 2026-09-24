package fakeapi

import (
	"fmt"
	"k8s-learning/event"
	"k8s-learning/eventhandle"
	fakecluster "k8s-learning/fakeCluster"
	"k8s-learning/k8sinterface"
	namegenerator "k8s-learning/nameGenerator"
	"k8s-learning/resource"
)

type FakeAPIServer struct {
	Cluster      *fakecluster.FakeCluster
	Event_handel *eventhandle.EventHandle
}

type APIServer interface {
	Create(r k8sinterface.Resource) error
	Get(kind, namespace, name string) (k8sinterface.Resource, error)
	Update(r k8sinterface.Resource) error
	Delete(kind, namespace, name string) error
}

func NewFakeAPIServer() *FakeAPIServer {
	return &FakeAPIServer{
		Cluster:      fakecluster.NewFakeCluster(),
		Event_handel: eventhandle.NewEventHandle(),
	}
}

func (s *FakeAPIServer) Create(r k8sinterface.Resource) error {
	kind := r.GetKind()
	var name, namespace string
	switch kind {
	case "Pod":
		r, ok := r.(*resource.Pod)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Pod")
		}
		namespace = r.Metadata.Namespace
		name = r.Metadata.Name
		for _, pod := range s.Cluster.Pods {
			if pod.Metadata.Namespace == namespace && pod.Metadata.Name == name {
				return fmt.Errorf("the pod %s has already existed", pod.Metadata.Name)
			}
		}
		s.Cluster.Pods = append(s.Cluster.Pods, *r)
	case "Deployment":
		r, ok := r.(*resource.Deployment)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Deployment")
		}
		namespace = r.Metadata.Namespace
		name = r.Metadata.Name
		for _, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == name && dep.Metadata.Namespace == namespace {
				return fmt.Errorf("the deployment %s has already existed", dep.Metadata.Name)
			}
		}
		s.Cluster.Deployments = append(s.Cluster.Deployments, *r)
	case "ReplicaSet":
		r, ok := r.(*resource.ReplicaSet)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected: replicaset")
		}
		namespace = r.Metadata.Namespace
		name = r.Metadata.Name
		for _, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == name && rep.Metadata.Namespace == namespace {
				return fmt.Errorf("the ReplicaSet %s has already existed", rep.Metadata.Name)
			}
		}
		s.Cluster.ReplicaSets = append(s.Cluster.ReplicaSets, *r)
	default:
		return fmt.Errorf("resource type not found")
	}
	event := event.Event{
		EventType: "ADD",
		Name:      name,
		Namespace: namespace,
		Kind:      kind,
	}
	s.Event_handel.OnAdd(&event)
	return nil
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
	found := false
	switch r.GetKind() {
	case "Pod":
		desiredPod, ok := r.(*resource.Pod)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Pod")
		}
		for i, pod := range s.Cluster.Pods {
			if pod.Metadata.Name == r.GetName() && pod.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.Pods[i] = *desiredPod
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the Pod %s not found", r.GetName())
		}
	case "Deployment":
		desiredDep, ok := r.(*resource.Deployment)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected Deployment")
		}
		for i, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == r.GetName() && dep.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.Deployments[i] = *desiredDep
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the Deployment %s not found", r.GetName())
		}
	case "ReplicaSet":
		desiredRep, ok := r.(*resource.ReplicaSet)
		if !ok {
			return fmt.Errorf("resource type mismatch: expected: replicaset")
		}
		for i, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == r.GetName() && rep.Metadata.Namespace == r.GetNamespace() {
				s.Cluster.ReplicaSets[i] = *desiredRep
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the ReplicaSet %s not found", r.GetName())
		}
	default:
		return fmt.Errorf("the resource type not found")
	}
	event := event.Event{
		EventType: "UPDATE",
		Name:      r.GetName(),
		Namespace: r.GetNamespace(),
		Kind:      r.GetKind(),
	}
	s.Event_handel.OnAdd(&event)
	return nil
}
func (s *FakeAPIServer) Delete(kind, namespace, name string) error {
	found := false
	switch kind {
	case "Pod":
		for i, pod := range s.Cluster.Pods {
			if pod.Metadata.Name == name && pod.Metadata.Namespace == namespace {
				s.Cluster.Pods = append(s.Cluster.Pods[:i], s.Cluster.Pods[i+1:]...)
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the Pod %s not found", name)
		}
	case "Deployment":
		for i, dep := range s.Cluster.Deployments {
			if dep.Metadata.Name == name && dep.Metadata.Namespace == namespace {
				s.Cluster.Deployments = append(s.Cluster.Deployments[:i], s.Cluster.Deployments[i+1:]...)
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the Deployment %s not found", name)
		}
	case "ReplicaSet":
		for i, rep := range s.Cluster.ReplicaSets {
			if rep.Metadata.Name == name && rep.Metadata.Namespace == namespace {
				s.Cluster.ReplicaSets = append(s.Cluster.ReplicaSets[:i], s.Cluster.ReplicaSets[i+1:]...)
				found = true
				break
			}
		}
		if found == false {
			return fmt.Errorf("the ReplicaSet %s not found", name)
		}
	default:
		return fmt.Errorf("the resource type not found ")
	}
	event := event.Event{
		EventType: "DELETE",
		Name:      name,
		Namespace: namespace,
		Kind:      kind,
	}
	s.Event_handel.OnAdd(&event)
	return nil
}

func (fc *FakeAPIServer) GetPodsByLabels(labels map[string]string) (*[]resource.Pod, error) {
	var matchedPods []resource.Pod
	fakeClusterPods := fc.Cluster.GetFakeClusterPods()
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

func (fc *FakeAPIServer) CreatePodWithTemplate(template *resource.PodTemplate, name string) error {
	var newPod resource.Pod
	newPod.APIVersion = "v1"
	newPod.Kind = "Pod"
	newPod.Metadata = template.Metadata
	newPod.Metadata.Name = namegenerator.GenerateName(name)
	newPod.Metadata.Labels = template.Metadata.Labels
	newPod.Metadata.Namespace = template.Metadata.Namespace
	newPod.Status = resource.Status{Phase: "Pending"}

	return fc.Create(&newPod)
}
