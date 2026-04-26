/*
Copyright 2026.

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

// WallAppSpec defines the desired state of WallApp
type WallAppSpec struct {
	// Foo *string `json:"foo,omitempty"`
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`

	StorageSize string `json:"storageSize"`
	DbPassword  string `json:"dbPassword"`
}

// WallAppStatus defines the observed state of WallApp.
type WallAppStatus struct {
	AvailableReplicas int32              `json:"avaibleReplicas"`
	DatacaseStatus    string             `json:"databaseStatus"`
	Conditions        []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WallApp is the Schema for the wallapps API
type WallApp struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of WallApp
	// +required
	Spec WallAppSpec `json:"spec"`

	// status defines the observed state of WallApp
	// +optional
	Status WallAppStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// WallAppList contains a list of WallApp
type WallAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []WallApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WallApp{}, &WallAppList{})
}
