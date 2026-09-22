package main

import (
	"fmt"
	"k8s-learning/client"
	"k8s-learning/controller"
	"k8s-learning/factory"
	"k8s-learning/fakeapi"
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
	server := fakeapi.NewFakeAPIServer()
	for _, r := range resources {
		server.Create(r)
	}
	pods := server.Cluster.Pods
	for _, pod := range pods {
		pod.GetInfo()
	}
	r, _ := server.Get("Deployment", "ruoyi", "ruoyi")
	r.GetInfo()
	dep := resource.Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata:   resource.Metadata{Name: "ruoyi", Namespace: "prod"},
		Replicas:   3,
	}
	server.Update(&dep)
	r, _ = server.Get("Deployment", "prod", "ruoyi")
	r.GetInfo()
	server.Delete("Deployment", "prod", "ruoyi")
	r, err := server.Get("Deployment", "prod", "ruoyi")
	if err != nil {
		fmt.Println("delete falsed:", err)
	}
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// wq := workqueue.NewWorkQueue(10)
	// worker := &worker.Worker{Queue: *wq}
	// go func() {
	// 	defer wg.Done()
	// 	worker.Run()
	// }()
	// wq.Add("default/pod1")
	// wq.Add("default/pod2")
	// wg.Wait()
}
