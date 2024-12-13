// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesendpointsv1


type DataKubernetesEndpointsV1Subset struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/endpoints_v1#address DataKubernetesEndpointsV1#address}
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// not_ready_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/endpoints_v1#not_ready_address DataKubernetesEndpointsV1#not_ready_address}
	NotReadyAddress interface{} `field:"optional" json:"notReadyAddress" yaml:"notReadyAddress"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/endpoints_v1#port DataKubernetesEndpointsV1#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
}

