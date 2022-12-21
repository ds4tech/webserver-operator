/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebserverSpec defines the desired state of Webserver
type WebserverSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// Size defines the number of Webserver instances
	// The following markers will use OpenAPI v3 schema to validate the value
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:ExclusiveMaximum=false
	Size int32 `json:"size,omitempty"`
}

// WebserverStatus defines the observed state of Webserver
type WebserverStatus struct {
	// Represents the observations of a Webserver's current state.
	// Webserver.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// Webserver.status.conditions.status are one of True, False, Unknown.
	// Webserver.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// Webserver.status.conditions.Message is a human readable message indicating details about the transition.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Webserver is the Schema for the webservers API
type Webserver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebserverSpec   `json:"spec,omitempty"`
	Status WebserverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WebserverList contains a list of Webserver
type WebserverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webserver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Webserver{}, &WebserverList{})
}
