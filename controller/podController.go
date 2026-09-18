package controller

import (
	"fmt"
	"k8s-learning/client"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type PodController struct {
	Client client.Client
}

func (pc *PodController) Handle(r k8sinterface.Resource) {
	pod := r.(*resource.Pod)
	pc.Reconcile(pod)
}

func (pc *PodController) Reconcile(desiredpod *resource.Pod) {
	currentPod, err := pc.Client.GetPod(desiredpod.Metadata.Name)
	if err != nil {
		pc.Client.CreatePod(desiredpod)
		fmt.Printf("pod %s 已创建\n", desiredpod.Metadata.Name)
		return
	}
	currentPod.GetInfo()
	if desiredpod.Status.Phase != currentPod.Status.Phase {
		fmt.Printf("Pod: %s 当前状态与期望值不符,当前：%s 期望：%s", desiredpod.Metadata.Name, currentPod.Status.Phase, desiredpod.Status.Phase)
	}
	currentPod.Status.Phase = desiredpod.Status.Phase
	err = pc.Client.UpdatePod(currentPod)
	if err != nil {
		fmt.Printf("更新pod状态失败: %v\n", err)
		return
	} else {
		fmt.Printf("pod:%s 状态已更新\n", desiredpod.Metadata.Name)
	}

}
