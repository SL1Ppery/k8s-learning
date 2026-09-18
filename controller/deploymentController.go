package controller

import (
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

func (dc *DeploymentController) Reconcile(dep *resource.Deployment) {
	currentDep, err := dc.Client.GetDeployment(dep.GetName())
	if err != nil {
		dc.Client.CreateDeployment(dep)
		return
	}
	currentDep.GetInfo()
}
