// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpoints


type EndpointsSubsetAddress struct {
	// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/endpoints#ip Endpoints#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The Hostname of this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/endpoints#hostname Endpoints#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Node hosting this endpoint. This can be used to determine endpoints local to a node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/endpoints#node_name Endpoints#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
}

