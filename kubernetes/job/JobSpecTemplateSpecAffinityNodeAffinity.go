// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecAffinityNodeAffinity struct {
	// preferred_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/job#preferred_during_scheduling_ignored_during_execution Job#preferred_during_scheduling_ignored_during_execution}
	PreferredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	// required_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/job#required_during_scheduling_ignored_during_execution Job#required_during_scheduling_ignored_during_execution}
	RequiredDuringSchedulingIgnoredDuringExecution *JobSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

