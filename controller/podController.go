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

func (pc *PodController) Reconcile(pod *resource.Pod) {
	currentPod, err := pc.Client.GetPod(pod.Metadata.Name)
	if err != nil {
		pc.Client.CreatePod(pod)
		fmt.Printf("pod %s 已创建\n", pod.Metadata.Name)
		return
	}
	currentPod.GetInfo()
}
