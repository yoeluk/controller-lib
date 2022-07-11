/*
Copyright 2017 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Rdsdb is a specification for a Rdsdb resource
type Rdsdb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsdbSpec   `json:"spec"`
	Status RdsdbStatus `json:"status"`
}

// RdsdbSpec is the spec for a Rdsdb resource
type RdsdbSpec struct {
	DbName     string `json:"dbName"`
	DbUserName *int32 `json:"dbUserName"`
}

// RdsdbStatus is the status for a Rdsdb resource
type RdsdbStatus struct {
	DbCreated bool `json:"dbCreated"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RdsdbList is a list of Rdsdb resources
type RdsdbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Rdsdb `json:"items"`
}
