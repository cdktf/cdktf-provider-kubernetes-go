// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointslicev1


type EndpointSliceV1Endpoint struct {
	// addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#addresses EndpointSliceV1#addresses}
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#condition EndpointSliceV1#condition}
	Condition *EndpointSliceV1EndpointCondition `field:"optional" json:"condition" yaml:"condition"`
	// hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#hostname EndpointSliceV1#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// nodeName represents the name of the Node hosting this endpoint.
	//
	// This can be used to determine endpoints local to a Node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#node_name EndpointSliceV1#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
	// target_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#target_ref EndpointSliceV1#target_ref}
	TargetRef *EndpointSliceV1EndpointTargetRef `field:"optional" json:"targetRef" yaml:"targetRef"`
	// zone is the name of the Zone this endpoint exists in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/endpoint_slice_v1#zone EndpointSliceV1#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

