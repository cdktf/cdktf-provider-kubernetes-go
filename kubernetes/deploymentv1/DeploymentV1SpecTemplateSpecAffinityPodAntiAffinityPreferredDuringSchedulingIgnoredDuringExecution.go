// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment_v1#pod_affinity_term DeploymentV1#pod_affinity_term}
	PodAffinityTerm *DeploymentV1SpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment_v1#weight DeploymentV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

