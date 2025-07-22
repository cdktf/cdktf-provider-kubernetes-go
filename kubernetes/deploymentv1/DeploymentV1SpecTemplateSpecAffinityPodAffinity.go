// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplateSpecAffinityPodAffinity struct {
	// preferred_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/deployment_v1#preferred_during_scheduling_ignored_during_execution DeploymentV1#preferred_during_scheduling_ignored_during_execution}
	PreferredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	// required_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/deployment_v1#required_during_scheduling_ignored_during_execution DeploymentV1#required_during_scheduling_ignored_during_execution}
	RequiredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

