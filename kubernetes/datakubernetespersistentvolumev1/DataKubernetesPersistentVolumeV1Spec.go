// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1Spec struct {
	// Contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#access_modes DataKubernetesPersistentVolumeV1#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#capacity DataKubernetesPersistentVolumeV1#capacity}
	Capacity *map[string]*string `field:"required" json:"capacity" yaml:"capacity"`
	// persistent_volume_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#persistent_volume_source DataKubernetesPersistentVolumeV1#persistent_volume_source}
	PersistentVolumeSource *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource `field:"required" json:"persistentVolumeSource" yaml:"persistentVolumeSource"`
	// claim_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#claim_ref DataKubernetesPersistentVolumeV1#claim_ref}
	ClaimRef *DataKubernetesPersistentVolumeV1SpecClaimRef `field:"optional" json:"claimRef" yaml:"claimRef"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#mount_options DataKubernetesPersistentVolumeV1#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#node_affinity DataKubernetesPersistentVolumeV1#node_affinity}
	NodeAffinity *DataKubernetesPersistentVolumeV1SpecNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// What happens to a persistent volume when released from its claim.
	//
	// Valid options are Retain (default) and Recycle. Recycling must be supported by the volume plugin underlying this persistent volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#persistent_volume_reclaim_policy DataKubernetesPersistentVolumeV1#persistent_volume_reclaim_policy}
	PersistentVolumeReclaimPolicy *string `field:"optional" json:"persistentVolumeReclaimPolicy" yaml:"persistentVolumeReclaimPolicy"`
	// A description of the persistent volume's class. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#storage_class_name DataKubernetesPersistentVolumeV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines if a volume is intended to be used with a formatted filesystem.
	//
	// or to remain in raw block state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/data-sources/persistent_volume_v1#volume_mode DataKubernetesPersistentVolumeV1#volume_mode}
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
}

