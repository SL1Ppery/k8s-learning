package controller

import (
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
		// Handle error
		return
	}
	currentcount := len(pods)

	if currentcount < desiredrs.Replicas {
		needCreate := desiredrs.Replicas - currentcount
		for i := 0; i < needCreate; i++ {
			err := rsc.Client.CreatePodWithTemplate(&desiredrs.Template, desiredrs.Metadata.Name)
			if err != nil {
				// Handle error
				return
			}
		}
	}
	if currentcount == desiredrs.Replicas {
		return
	}
	if currentcount > desiredrs.Replicas {
		needDelete := currentcount - desiredrs.Replicas
		for i := 0; i < needDelete; i++ {
			podToDelete := (*pods)[i]
			err := rsc.Client.DeletePod(podToDelete.Metadata.Name)
			if err != nil {
				// Handle error
				return
			}
		}
	}
}
