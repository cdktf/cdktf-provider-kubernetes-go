// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumeclaimv1


type PersistentVolumeClaimV1Spec struct {
	// A set of the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#access_modes PersistentVolumeClaimV1#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#resources PersistentVolumeClaimV1#resources}
	Resources *PersistentVolumeClaimV1SpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#selector PersistentVolumeClaimV1#selector}
	Selector *PersistentVolumeClaimV1SpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#storage_class_name PersistentVolumeClaimV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#volume_mode PersistentVolumeClaimV1#volume_mode}
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim_v1#volume_name PersistentVolumeClaimV1#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

