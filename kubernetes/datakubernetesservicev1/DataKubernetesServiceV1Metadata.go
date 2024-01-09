// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesservicev1


type DataKubernetesServiceV1Metadata struct {
	// An unstructured key value map stored with the service that may be used to store arbitrary metadata.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/service_v1#annotations DataKubernetesServiceV1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	//
	// May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/service_v1#labels DataKubernetesServiceV1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the service, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/service_v1#name DataKubernetesServiceV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace defines the space within which name of the service must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/service_v1#namespace DataKubernetesServiceV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

