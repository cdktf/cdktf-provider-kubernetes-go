// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs struct {
	// The endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#endpoints_name DataKubernetesPersistentVolumeV1#endpoints_name}
	EndpointsName *string `field:"required" json:"endpointsName" yaml:"endpointsName"`
	// The Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#path DataKubernetesPersistentVolumeV1#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#read_only DataKubernetesPersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

