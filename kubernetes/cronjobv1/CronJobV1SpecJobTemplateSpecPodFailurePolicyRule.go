// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecPodFailurePolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/cron_job_v1#action CronJobV1#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// on_exit_codes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/cron_job_v1#on_exit_codes CronJobV1#on_exit_codes}
	OnExitCodes *CronJobV1SpecJobTemplateSpecPodFailurePolicyRuleOnExitCodes `field:"optional" json:"onExitCodes" yaml:"onExitCodes"`
	// on_pod_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/cron_job_v1#on_pod_condition CronJobV1#on_pod_condition}
	OnPodCondition interface{} `field:"optional" json:"onPodCondition" yaml:"onPodCondition"`
}

