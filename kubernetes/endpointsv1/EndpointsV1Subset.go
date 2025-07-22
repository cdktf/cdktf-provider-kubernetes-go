// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointsv1


type EndpointsV1Subset struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoints_v1#address EndpointsV1#address}
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// not_ready_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoints_v1#not_ready_address EndpointsV1#not_ready_address}
	NotReadyAddress interface{} `field:"optional" json:"notReadyAddress" yaml:"notReadyAddress"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoints_v1#port EndpointsV1#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
}

