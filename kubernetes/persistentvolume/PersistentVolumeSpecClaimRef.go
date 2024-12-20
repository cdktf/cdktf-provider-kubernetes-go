// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecClaimRef struct {
	// The name of the PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/persistent_volume#name PersistentVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/persistent_volume#namespace PersistentVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

