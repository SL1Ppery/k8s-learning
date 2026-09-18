package controller

import (
	"fmt"
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

func (cm *ControllerManager) Handle(r k8sinterface.Resource) {
	controller, ok := cm.Controllers[r.GetKind()]
	if !ok {
		fmt.Printf("没有对应的controller%v\n", r.GetKind())
		return
	}
	controller.Handle(r)
}
