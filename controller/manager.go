package controller

import (
	"fmt"
	fakecluster "k8s-learning/fakeCluster"
	"k8s-learning/k8sinterface"
)

type ControllerManager struct {
	Controllers map[string]k8sinterface.Controller
}

func NewControllerManager() *ControllerManager {
	return &ControllerManager{
		Controllers: make(map[string]k8sinterface.Controller),
	}
}

func (cm *ControllerManager) Registry(kind string, controller k8sinterface.Controller) {
	cm.Controllers[kind] = controller
}

func (cm *ControllerManager) Handle(fakecluster *fakecluster.FakeCluster, key string) {
	resource := fakecluster.GetResourceByMetadata(key)

	if resource == nil {
		fmt.Println("资源不存在:", key)
		return
	}

	controller, ok := cm.Controllers[resource.GetKind()]
	if !ok {
		fmt.Printf("没有%s控制器\n", resource.GetKind())
		return
	}

	controller.Handle(resource)
}
