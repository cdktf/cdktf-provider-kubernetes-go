// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecVolumeProjectedSourcesConfigMap struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod_v1#items PodV1#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the ConfigMap or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod_v1#optional PodV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

