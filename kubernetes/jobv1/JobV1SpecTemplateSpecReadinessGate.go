// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/job_v1#condition_type JobV1#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

