// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyv1


type NetworkPolicyV1SpecIngressFrom struct {
	// ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/network_policy_v1#ip_block NetworkPolicyV1#ip_block}
	IpBlock *NetworkPolicyV1SpecIngressFromIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/network_policy_v1#namespace_selector NetworkPolicyV1#namespace_selector}
	NamespaceSelector *NetworkPolicyV1SpecIngressFromNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// pod_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/network_policy_v1#pod_selector NetworkPolicyV1#pod_selector}
	PodSelector *NetworkPolicyV1SpecIngressFromPodSelector `field:"optional" json:"podSelector" yaml:"podSelector"`
}

