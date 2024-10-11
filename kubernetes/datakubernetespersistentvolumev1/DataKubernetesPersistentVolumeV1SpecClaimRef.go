// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecClaimRef struct {
	// The name of the PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/data-sources/persistent_volume_v1#name DataKubernetesPersistentVolumeV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/data-sources/persistent_volume_v1#namespace DataKubernetesPersistentVolumeV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

