// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecPodFailurePolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/job_v1#action JobV1#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// on_exit_codes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/job_v1#on_exit_codes JobV1#on_exit_codes}
	OnExitCodes *JobV1SpecPodFailurePolicyRuleOnExitCodes `field:"optional" json:"onExitCodes" yaml:"onExitCodes"`
	// on_pod_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/job_v1#on_pod_condition JobV1#on_pod_condition}
	OnPodCondition interface{} `field:"optional" json:"onPodCondition" yaml:"onPodCondition"`
}

