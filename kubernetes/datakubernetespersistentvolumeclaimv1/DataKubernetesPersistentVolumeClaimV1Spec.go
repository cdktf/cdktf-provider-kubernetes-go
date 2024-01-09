// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumeclaimv1


type DataKubernetesPersistentVolumeClaimV1Spec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/persistent_volume_claim_v1#selector DataKubernetesPersistentVolumeClaimV1#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/persistent_volume_claim_v1#storage_class_name DataKubernetesPersistentVolumeClaimV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/persistent_volume_claim_v1#volume_mode DataKubernetesPersistentVolumeClaimV1#volume_mode}
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/persistent_volume_claim_v1#volume_name DataKubernetesPersistentVolumeClaimV1#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

