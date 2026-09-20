package main

import (
	"k8s-learning/client"
	"k8s-learning/controller"
	"k8s-learning/factory"
	myparser "k8s-learning/parser"
	"k8s-learning/resource"
)

func main() {
	registry := factory.NewRegistry()
	registry.Register("Pod", resource.NewPod)
	registry.Register("Deployment", resource.NewDeployment)
	registry.Register("ReplicaSet", resource.NewReplicaSet)
	controllerManager := controller.NewControllerManager()
	fakeClient := &client.FakeClient{}
	podController := &controller.PodController{Client: fakeClient}
	deploymentController := &controller.DeploymentController{Client: fakeClient}
	replicaSetController := &controller.ReplicaSetController{Client: fakeClient}
	controllerManager.Registry("Pod", podController)
	controllerManager.Registry("Deployment", deploymentController)
	controllerManager.Registry("ReplicaSet", replicaSetController)
	resources := myparser.Parser("test.yaml", registry)
	for _, r := range resources {
		controllerManager.Handle(r)
	}
}
