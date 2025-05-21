// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecPodFailurePolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/cron_job#action CronJob#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// on_exit_codes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/cron_job#on_exit_codes CronJob#on_exit_codes}
	OnExitCodes *CronJobSpecJobTemplateSpecPodFailurePolicyRuleOnExitCodes `field:"optional" json:"onExitCodes" yaml:"onExitCodes"`
	// on_pod_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/cron_job#on_pod_condition CronJob#on_pod_condition}
	OnPodCondition interface{} `field:"optional" json:"onPodCondition" yaml:"onPodCondition"`
}

