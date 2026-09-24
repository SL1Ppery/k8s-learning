package main

import (
	"fmt"
	"k8s-learning/controller"
	"k8s-learning/factory"
	"k8s-learning/fakeapi"
	myparser "k8s-learning/parser"
	"k8s-learning/resource"
	"k8s-learning/worker"
)

func main() {
	//init
	registry := factory.NewRegistry()
	registry.Register("Pod", resource.NewPod)
	registry.Register("Deployment", resource.NewDeployment)
	registry.Register("ReplicaSet", resource.NewReplicaSet)
	apiserver := fakeapi.NewFakeAPIServer()
	controllerManager := controller.NewControllerManager()
	podController := &controller.PodController{FakeAPIServer: apiserver}
	deploymentController := &controller.DeploymentController{FakeAPIServer: apiserver}
	replicaSetController := &controller.ReplicaSetController{FakeAPIServer: apiserver}
	controllerManager.Registry("Pod", podController)
	controllerManager.Registry("Deployment", deploymentController)
	controllerManager.Registry("ReplicaSet", replicaSetController)
	worker := worker.Worker{
		Queue:             *apiserver.Event_handel.Queue,
		ControllerManager: *controllerManager,
		FakeAPIServer:     *apiserver,
	}
	go worker.Run()

	resources := myparser.Parser("test.yaml", registry)
	for _, r := range resources {
		apiserver.Create(r)
	}
	pod, err := apiserver.Get("Pod", "default", "nginx")
	if err != nil {
		fmt.Println(err)
	}
	pod.GetInfo()
}
