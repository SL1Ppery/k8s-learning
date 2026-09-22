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

func (cm *ControllerManager) Handle(key string) {
	rescource := fakecluster.GetResourceByMetadata(key)
	controller, ok := cm.Controllers[rescource.GetName()]
	if !ok {
		fmt.Printf("没有%s控制器", rescource.GetKind())
		return
	}
	controller.Handle(rescource)
}
