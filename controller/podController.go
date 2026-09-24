package controller

import (
	"fmt"
	"k8s-learning/fakeapi"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type PodController struct {
	FakeAPIServer *fakeapi.FakeAPIServer
}

func (pc *PodController) Handle(r k8sinterface.Resource) {
	pod := r.(*resource.Pod)
	pc.Reconcile(pod)
}

func (pc *PodController) Reconcile(desiredpod *resource.Pod) error {
	currentPod, err := pc.FakeAPIServer.Get(desiredpod.Kind, desiredpod.Metadata.Namespace, desiredpod.Metadata.Name)
	if err != nil {
		pc.FakeAPIServer.Create(desiredpod)
		fmt.Printf("pod %s 已创建\n", desiredpod.Metadata.Name)
		return nil
	}
	pod, ok := currentPod.(*resource.Pod)
	if !ok {
		return fmt.Errorf("resource type mismatch: expected Pod")
	}
	pod.GetInfo()
	if desiredpod.Status.Phase != pod.Status.Phase {
		fmt.Printf("Pod: %s 当前状态与期望值不符,当前：%s 期望：%s", desiredpod.Metadata.Name, pod.Status.Phase, desiredpod.Status.Phase)
	}
	pod.Status.Phase = desiredpod.Status.Phase
	err = pc.FakeAPIServer.Update(pod)
	if err != nil {
		return fmt.Errorf("更新pod状态失败: %v\n", err)
	} else {
		fmt.Printf("pod:%s 状态已更新\n", desiredpod.Metadata.Name)
		return nil
	}
}
