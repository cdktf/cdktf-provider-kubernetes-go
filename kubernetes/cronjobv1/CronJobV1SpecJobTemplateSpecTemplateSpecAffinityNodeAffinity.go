// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity struct {
	// preferred_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#preferred_during_scheduling_ignored_during_execution CronJobV1#preferred_during_scheduling_ignored_during_execution}
	PreferredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	// required_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#required_during_scheduling_ignored_during_execution CronJobV1#required_during_scheduling_ignored_during_execution}
	RequiredDuringSchedulingIgnoredDuringExecution *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

