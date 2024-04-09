// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecNodeAffinity struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/persistent_volume#required PersistentVolume#required}
	Required *PersistentVolumeSpecNodeAffinityRequired `field:"optional" json:"required" yaml:"required"`
}

