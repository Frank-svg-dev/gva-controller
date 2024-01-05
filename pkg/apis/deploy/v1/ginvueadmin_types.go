/*
Copyright 2024 pp.

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
type ResourceHandle struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type BackendHandle struct {
	Replicas        *int32         `json:"replicas,omitempty"`
	Image           string         `json:"image,omitempty"`
	ImagePullPolicy string         `json:"imagePullPolicy,omitempty"`
	ContainerPort   *int32         `json:"container_port,omitempty"`
	Resource        ResourceHandle `json:"resource,omitempty"`
	Limit           ResourceHandle `json:"limit,omitempty"`
}

type FrontHandle struct {
	Replicas        *int32         `json:"replicas,omitempty"`
	Image           string         `json:"image,omitempty"`
	ImagePullPolicy string         `json:"imagePullPolicy,omitempty"`
	ContainerPort   *int32         `json:"container_port,omitempty"`
	Resource        ResourceHandle `json:"resource,omitempty"`
	Limit           ResourceHandle `json:"limit,omitempty"`
}

// GinVueAdminSpec defines the desired state of GinVueAdmin
type GinVueAdminSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name        string `json:"name,omitempty"`
	ClusterMode string `json:"cluster_mode,omitempty"`
	Backend     string `json:"backend,omitempty"`
	Front       string `json:"front,omitempty"`
}

// GinVueAdminStatus defines the observed state of GinVueAdmin
type GinVueAdminStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GinVueAdmin is the Schema for the ginvueadmins API
type GinVueAdmin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GinVueAdminSpec   `json:"spec,omitempty"`
	Status GinVueAdminStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GinVueAdminList contains a list of GinVueAdmin
type GinVueAdminList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GinVueAdmin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GinVueAdmin{}, &GinVueAdminList{})
}
