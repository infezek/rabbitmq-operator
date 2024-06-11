/*
Copyright 2024 Ezequiel Lopes.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"time"

	apiv1 "github.com/infezek/rabbitmq-operator/api/v1"

	"github.com/sirupsen/logrus"
	apps "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RabbitHPAReconciler reconciles a RabbitHPA object
type RabbitHPAReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *RabbitHPAReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logrus.WithField("Type", "[RECONCILE]")
	log.Infof("Starting")
	defer log.Infof("Finished")
	// Buscar as especificações do RabbitHPA
	request := &apiv1.RabbitHPA{}
	if err := r.Get(ctx, req.NamespacedName, request); err != nil {
		if errors.IsNotFound(err) {
			log.Info("RabbitHPA resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	// Buscar o deployment
	deployment := &apps.Deployment{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: request.Spec.HPA.Namespace,
		Name:      request.Spec.HPA.Label,
	}, deployment)
	if err != nil {
		log.Error(err.Error())
		return ctrl.Result{}, err
	}
	rabbitProps := request.Spec.Rabbit
	// Buscar as propriedades da fila
	gateway := NewRabbitMQGateway(rabbitProps.Host, rabbitProps.Namespace, rabbitProps.Username, rabbitProps.Password)
	queueProps, err := gateway.RequestQuery(rabbitProps.Queue)
	if err != nil {
		log.Errorf(err.Error())
		return ctrl.Result{}, err
	}
	target := request.CalculateReplicas(queueProps.MessagesReady)
	log.WithFields(logrus.Fields{
		"current": *deployment.Spec.Replicas,
		"target":  target,
	}).Infof("Current ReplicaSet Deployment")
	if *deployment.Spec.Replicas != target {
		log.Info("Updating deployment replicas")
		deployment.Spec.Replicas = &target
		err = r.Update(ctx, deployment)
		if err != nil {
			log.Errorf(err.Error())
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
}

func (r *RabbitHPAReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.RabbitHPA{}).
		Complete(r)
}
