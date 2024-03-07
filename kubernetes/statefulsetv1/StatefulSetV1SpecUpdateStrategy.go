// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecUpdateStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set_v1#rolling_update StatefulSetV1#rolling_update}
	RollingUpdate interface{} `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Indicates the type of the StatefulSet update strategy. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set_v1#type StatefulSetV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

