// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpoints


type EndpointsSubset struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/endpoints#address Endpoints#address}
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// not_ready_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/endpoints#not_ready_address Endpoints#not_ready_address}
	NotReadyAddress interface{} `field:"optional" json:"notReadyAddress" yaml:"notReadyAddress"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/endpoints#port Endpoints#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
}

