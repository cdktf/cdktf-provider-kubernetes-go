// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecUpdateStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/stateful_set#rolling_update StatefulSet#rolling_update}
	RollingUpdate interface{} `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Indicates the type of the StatefulSet update strategy. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/stateful_set#type StatefulSet#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

