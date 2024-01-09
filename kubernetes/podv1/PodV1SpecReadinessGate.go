// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#condition_type PodV1#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

