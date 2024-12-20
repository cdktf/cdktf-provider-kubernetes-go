// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte struct {
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/persistent_volume_v1#registry DataKubernetesPersistentVolumeV1#registry}
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	// Volume is a string that references an already created Quobyte volume by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/persistent_volume_v1#volume DataKubernetesPersistentVolumeV1#volume}
	Volume *string `field:"required" json:"volume" yaml:"volume"`
	// Group to map volume access to Default is no group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/persistent_volume_v1#group DataKubernetesPersistentVolumeV1#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/persistent_volume_v1#read_only DataKubernetesPersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// User to map volume access to Defaults to serivceaccount user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/persistent_volume_v1#user DataKubernetesPersistentVolumeV1#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

