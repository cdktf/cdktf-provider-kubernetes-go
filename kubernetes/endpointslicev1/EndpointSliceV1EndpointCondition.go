// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointslicev1


type EndpointSliceV1EndpointCondition struct {
	// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#ready EndpointSliceV1#ready}
	Ready interface{} `field:"optional" json:"ready" yaml:"ready"`
	// serving is identical to ready except that it is set regardless of the terminating state of endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#serving EndpointSliceV1#serving}
	Serving interface{} `field:"optional" json:"serving" yaml:"serving"`
	// terminating indicates that this endpoint is terminating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#terminating EndpointSliceV1#terminating}
	Terminating interface{} `field:"optional" json:"terminating" yaml:"terminating"`
}

