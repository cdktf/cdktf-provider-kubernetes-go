// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecPodFailurePolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/job#action Job#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// on_exit_codes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/job#on_exit_codes Job#on_exit_codes}
	OnExitCodes *JobSpecPodFailurePolicyRuleOnExitCodes `field:"optional" json:"onExitCodes" yaml:"onExitCodes"`
	// on_pod_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/job#on_pod_condition Job#on_pod_condition}
	OnPodCondition interface{} `field:"optional" json:"onPodCondition" yaml:"onPodCondition"`
}

