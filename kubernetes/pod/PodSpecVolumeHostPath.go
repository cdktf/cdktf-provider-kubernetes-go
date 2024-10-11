// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#path Pod#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#type Pod#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

