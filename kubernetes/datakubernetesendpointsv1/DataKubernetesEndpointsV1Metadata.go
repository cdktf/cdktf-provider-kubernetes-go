// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesendpointsv1


type DataKubernetesEndpointsV1Metadata struct {
	// An unstructured key value map stored with the endpoints that may be used to store arbitrary metadata.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/data-sources/endpoints_v1#annotations DataKubernetesEndpointsV1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Prefix, used by the server, to generate a unique name ONLY IF the `name` field has not been provided.
	//
	// This value will also be combined with a unique suffix. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/data-sources/endpoints_v1#generate_name DataKubernetesEndpointsV1#generate_name}
	GenerateName *string `field:"optional" json:"generateName" yaml:"generateName"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the endpoints.
	//
	// May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/data-sources/endpoints_v1#labels DataKubernetesEndpointsV1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the endpoints, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/data-sources/endpoints_v1#name DataKubernetesEndpointsV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace defines the space within which name of the endpoints must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/data-sources/endpoints_v1#namespace DataKubernetesEndpointsV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

