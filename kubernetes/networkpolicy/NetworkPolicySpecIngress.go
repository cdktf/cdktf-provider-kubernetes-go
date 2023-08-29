// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicy


type NetworkPolicySpecIngress struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/network_policy#from NetworkPolicy#from}
	From interface{} `field:"optional" json:"from" yaml:"from"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/network_policy#ports NetworkPolicy#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

