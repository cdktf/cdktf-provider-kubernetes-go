// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecAffinityNodeAffinity struct {
	// preferred_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set#preferred_during_scheduling_ignored_during_execution StatefulSet#preferred_during_scheduling_ignored_during_execution}
	PreferredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	// required_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set#required_during_scheduling_ignored_during_execution StatefulSet#required_during_scheduling_ignored_during_execution}
	RequiredDuringSchedulingIgnoredDuringExecution *StatefulSetSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

