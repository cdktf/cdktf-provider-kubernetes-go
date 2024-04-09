// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointslicev1


type EndpointSliceV1EndpointTargetRef struct {
	// Name of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/endpoint_slice_v1#name EndpointSliceV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/endpoint_slice_v1#field_path EndpointSliceV1#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// Namespace of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/endpoint_slice_v1#namespace EndpointSliceV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specific resourceVersion to which this reference is made, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/endpoint_slice_v1#resource_version EndpointSliceV1#resource_version}
	ResourceVersion *string `field:"optional" json:"resourceVersion" yaml:"resourceVersion"`
	// If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/endpoint_slice_v1#uid EndpointSliceV1#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

