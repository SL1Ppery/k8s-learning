package controller

import (
	"fmt"
	"k8s-learning/client"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type ReplicaSetController struct {
	Client client.Client
}

func (rsc *ReplicaSetController) Handle(r k8sinterface.Resource) {
	rs := r.(*resource.ReplicaSet)
	rsc.Reconcile(rs)
}
func (rsc *ReplicaSetController) Reconcile(desiredrs *resource.ReplicaSet) {
	pods, err := rsc.Client.GetPodsByLabels(desiredrs.Selector)
	if err != nil {
		fmt.Printf("查询 ReplicaSet %s 的 Pod 失败: %v\n", desiredrs.Metadata.Name, err)
		return
	}
	currentcount := len(*pods)

	if currentcount < desiredrs.Replicas {
		needCreate := desiredrs.Replicas - currentcount
		for i := 0; i < needCreate; i++ {
			err := rsc.Client.CreatePodWithTemplate(&desiredrs.Template, desiredrs.Metadata.Name)
			if err != nil {
				fmt.Printf("ReplicaSet %s 创建 Pod 失败: %v\n", desiredrs.Metadata.Name, err)
				return
			}
		}
		fmt.Printf("ReplicaSet %s 当前有 %d 个匹配 Pod，已创建 %d 个 Pod\n",
			desiredrs.Metadata.Name, currentcount, needCreate)
	} else if currentcount == desiredrs.Replicas {
		fmt.Printf("ReplicaSet %s 无需调整，当前有 %d 个匹配 Pod\n",
			desiredrs.Metadata.Name, currentcount)
	} else {
		needDelete := currentcount - desiredrs.Replicas
		fmt.Printf("ReplicaSet %s 需要删除 %d 个 Pod\n", desiredrs.Metadata.Name, needDelete)
	}

	rsc.printAdjustedPods(desiredrs)
}

func (rsc *ReplicaSetController) printAdjustedPods(desiredrs *resource.ReplicaSet) {
	pods, err := rsc.Client.GetPodsByLabels(desiredrs.Selector)
	if err != nil {
		fmt.Printf("查询 ReplicaSet %s 调整后的 Pod 失败: %v\n", desiredrs.Metadata.Name, err)
		return
	}

	fmt.Printf("%-40s %-40s %-40s\n", "Name", "Namespace", "Status")
	for i := range *pods {
		(*pods)[i].GetInfo()
	}
}
