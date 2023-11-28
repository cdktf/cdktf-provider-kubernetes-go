// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/persistent_volume#name PersistentVolume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/persistent_volume#namespace PersistentVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

