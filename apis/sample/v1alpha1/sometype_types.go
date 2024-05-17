/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// SomeTypeParameters are the configurable fields of a SomeType.
type SomeTypeParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// SomeTypeObservation are the observable fields of a SomeType.
type SomeTypeObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A SomeTypeSpec defines the desired state of a SomeType.
type SomeTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SomeTypeParameters `json:"forProvider"`
}

// A SomeTypeStatus represents the observed state of a SomeType.
type SomeTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SomeTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A SomeType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type SomeType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SomeTypeSpec   `json:"spec"`
	Status SomeTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SomeTypeList contains a list of SomeType
type SomeTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SomeType `json:"items"`
}

// SomeType type metadata.
var (
	SomeTypeKind             = reflect.TypeOf(SomeType{}).Name()
	SomeTypeGroupKind        = schema.GroupKind{Group: Group, Kind: SomeTypeKind}.String()
	SomeTypeKindAPIVersion   = SomeTypeKind + "." + SchemeGroupVersion.String()
	SomeTypeGroupVersionKind = SchemeGroupVersion.WithKind(SomeTypeKind)
)

func init() {
	SchemeBuilder.Register(&SomeType{}, &SomeTypeList{})
}
