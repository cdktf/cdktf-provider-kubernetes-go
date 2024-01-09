// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecContainerEnvValueFromConfigMapKeyRef struct {
	// The key to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#key PodV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#optional PodV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

