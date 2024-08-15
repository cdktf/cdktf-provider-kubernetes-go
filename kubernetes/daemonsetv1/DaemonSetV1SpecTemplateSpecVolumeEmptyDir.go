// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeEmptyDir struct {
	// What type of storage medium should back this directory.
	//
	// The default is "" which means to use the node's default medium. Must be one of ["" "Memory" "HugePages" "HugePages-2Mi" "HugePages-1Gi"]. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemon_set_v1#medium DaemonSetV1#medium}
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemon_set_v1#size_limit DaemonSetV1#size_limit}
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

