// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod_v1#optional PodV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

