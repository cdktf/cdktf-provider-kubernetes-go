// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicy


type NetworkPolicySpecEgressTo struct {
	// ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#ip_block NetworkPolicy#ip_block}
	IpBlock *NetworkPolicySpecEgressToIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#namespace_selector NetworkPolicy#namespace_selector}
	NamespaceSelector *NetworkPolicySpecEgressToNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// pod_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#pod_selector NetworkPolicy#pod_selector}
	PodSelector *NetworkPolicySpecEgressToPodSelector `field:"optional" json:"podSelector" yaml:"podSelector"`
}

