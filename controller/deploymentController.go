package controller

import (
	"fmt"
	"k8s-learning/client"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type DeploymentController struct {
	Client client.Client
}

func (dc *DeploymentController) Handle(r k8sinterface.Resource) {
	dep := r.(*resource.Deployment)
	dc.Reconcile(dep)
}

func (dc *DeploymentController) Reconcile(desireddep *resource.Deployment) {
	currentDep, err := dc.Client.GetDeployment(desireddep.GetName())
	if err != nil {
		dc.Client.CreateDeployment(desireddep)
		return
	}
	currentDep.GetInfo()

	if desireddep.Replicas != currentDep.Replicas {
		currentDep.Replicas = desireddep.Replicas
		err = dc.Client.UpdateDeployment(currentDep)
		if err != nil {
			fmt.Printf("更新deployment状态失败: %v\n", err)
			return
		}
		fmt.Printf("deployment状态 %s 已更新\n", desireddep.GetName())
	}
}
