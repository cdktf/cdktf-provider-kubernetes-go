// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm struct {
	// empty topology key is interpreted by the scheduler as 'all topologies'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment_v1#topology_key DeploymentV1#topology_key}
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment_v1#label_selector DeploymentV1#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment_v1#namespaces DeploymentV1#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment_v1#namespace_selector DeploymentV1#namespace_selector}
	NamespaceSelector interface{} `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

