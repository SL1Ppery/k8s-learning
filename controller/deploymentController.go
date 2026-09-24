package controller

import (
	"fmt"
	"k8s-learning/fakeapi"
	"k8s-learning/k8sinterface"
	"k8s-learning/resource"
)

type DeploymentController struct {
	FakeAPIServer *fakeapi.FakeAPIServer
}

func (dc *DeploymentController) Handle(r k8sinterface.Resource) {
	dep := r.(*resource.Deployment)
	dc.Reconcile(dep)
}

func (dc *DeploymentController) Reconcile(desireddep *resource.Deployment) error {
	currentDep, err := dc.FakeAPIServer.Get(desireddep.Kind, desireddep.Metadata.Namespace, desireddep.Metadata.Name)
	if err != nil {
		err = dc.FakeAPIServer.Create(desireddep)
		if err != nil {
			return fmt.Errorf("Deployment %s created failed", currentDep.GetName())
		}
		return nil
	}
	dep, ok := currentDep.(*resource.Deployment)
	if !ok {
		return fmt.Errorf("resource type mismatch")
	}
	fmt.Printf("%-40s %-40s %-40s\n", "Deployment/Name", "Namespace", "Replicas")
	dep.GetInfo()

	if desireddep.Replicas != dep.Replicas {
		dep.Replicas = desireddep.Replicas
		err = dc.FakeAPIServer.Update(dep)
		if err != nil {
			return fmt.Errorf("更新deployment状态失败: %v\n", err)
		}
		fmt.Printf("deployment状态 %s 已更新\n", desireddep.GetName())
	}
	return nil
}
