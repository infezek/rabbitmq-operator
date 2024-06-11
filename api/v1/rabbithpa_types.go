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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RabbitHPASpec defines the desired state of RabbitHPA
type RabbitHPASpec struct {
	Rabbit RabbitSpec `json:"rabbit"`
	HPA    HPASpec    `json:"hpa"`
}

type RabbitSpec struct {
	Namespace string `json:"namespace"`
	Host      string `json:"host"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Queue     string `json:"queue"`
}

type HPASpec struct {
	Label                string `json:"label"`
	Namespace            string `json:"namespace"`
	ReplicasetPerMessage int32  `json:"replicasetPerMessage"`
	MinReplicas          int32  `json:"minReplicas"`
	MaxReplicas          int32  `json:"maxReplicas"`
}

// RabbitHPAStatus defines the observed state of RabbitHPA
type RabbitHPAStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RabbitHPA is the Schema for the rabbithpas API
type RabbitHPA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitHPASpec   `json:"spec,omitempty"`
	Status RabbitHPAStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitHPAList contains a list of RabbitHPA
type RabbitHPAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitHPA `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RabbitHPA{}, &RabbitHPAList{})
}
