// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecVolumeDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set#container_name StatefulSet#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set#resource StatefulSet#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set#divisor StatefulSet#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

