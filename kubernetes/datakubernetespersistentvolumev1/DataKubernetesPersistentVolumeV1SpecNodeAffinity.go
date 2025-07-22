// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecNodeAffinity struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/data-sources/persistent_volume_v1#required DataKubernetesPersistentVolumeV1#required}
	Required *DataKubernetesPersistentVolumeV1SpecNodeAffinityRequired `field:"optional" json:"required" yaml:"required"`
}

