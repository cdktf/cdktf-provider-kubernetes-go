// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/deployment#node_affinity Deployment#node_affinity}
	NodeAffinity *DeploymentSpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/deployment#pod_affinity Deployment#pod_affinity}
	PodAffinity *DeploymentSpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/deployment#pod_anti_affinity Deployment#pod_anti_affinity}
	PodAntiAffinity *DeploymentSpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

