// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointslicev1


type EndpointSliceV1Port struct {
	// The application protocol for this port.
	//
	// This is used as a hint for implementations to offer richer behavior for protocols that they understand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoint_slice_v1#app_protocol EndpointSliceV1#app_protocol}
	AppProtocol *string `field:"required" json:"appProtocol" yaml:"appProtocol"`
	// port represents the port number of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoint_slice_v1#port EndpointSliceV1#port}
	Port *string `field:"required" json:"port" yaml:"port"`
	// name represents the name of this port. All ports in an EndpointSlice must have a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoint_slice_v1#name EndpointSliceV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/endpoint_slice_v1#protocol EndpointSliceV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

