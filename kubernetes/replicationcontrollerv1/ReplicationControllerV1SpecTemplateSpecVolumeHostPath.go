// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller_v1#path ReplicationControllerV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller_v1#type ReplicationControllerV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

