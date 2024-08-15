// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecVolumeDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#container_name Daemonset#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#resource Daemonset#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#divisor Daemonset#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

