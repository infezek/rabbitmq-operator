package v1

import "github.com/sirupsen/logrus"

func (r *RabbitHPA) CalculateReplicas(currentMessage int32) (target int32) {
	logrus.Info("Calculating replicas")
	defer logrus.Info("Replicas calculated")
	logrus.Info("Current message", currentMessage)
	target = currentMessage / r.Spec.HPA.ReplicasetPerMessage
	logrus.Info("Target replicas before validation", target)
	if target > r.Spec.HPA.MaxReplicas {
		return r.Spec.HPA.MaxReplicas
	}
	if target < r.Spec.HPA.MinReplicas {
		return r.Spec.HPA.MinReplicas
	}
	return target
}
