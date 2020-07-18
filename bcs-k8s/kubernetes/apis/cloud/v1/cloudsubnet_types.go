/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CloudSubnetSpec defines the desired state of CloudSubnet
type CloudSubnetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CloudSubnet. Edit CloudSubnet_types.go to remove/update
	SubnetID   string `json:"subnetID"`
	SubnetCidr string `json:"SubnetCidr"`
	VpcID      string `json:"vpcID"`
	Region     string `json:"region"`
	Zone       string `json:"zone"`
}

// CloudSubnetStatus defines the observed state of CloudSubnet
type CloudSubnetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AvailableIPNum int64  `json:"availableIPNum"`
	State          int32  `json:"state"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudSubnet is the Schema for the cloudsubnets API
type CloudSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudSubnetSpec   `json:"spec,omitempty"`
	Status CloudSubnetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudSubnetList contains a list of CloudSubnet
type CloudSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSubnet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSubnet{}, &CloudSubnetList{})
}
