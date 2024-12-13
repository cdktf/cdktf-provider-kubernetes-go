// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec struct {
	// A set of the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#access_modes PodV1#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#resources PodV1#resources}
	Resources *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#selector PodV1#selector}
	Selector *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#storage_class_name PodV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#volume_mode PodV1#volume_mode}
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#volume_name PodV1#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

