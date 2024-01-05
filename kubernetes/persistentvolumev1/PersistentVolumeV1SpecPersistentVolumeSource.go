// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSource struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#aws_elastic_block_store PersistentVolumeV1#aws_elastic_block_store}
	AwsElasticBlockStore *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#azure_disk PersistentVolumeV1#azure_disk}
	AzureDisk *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#azure_file PersistentVolumeV1#azure_file}
	AzureFile *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#ceph_fs PersistentVolumeV1#ceph_fs}
	CephFs *PersistentVolumeV1SpecPersistentVolumeSourceCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#cinder PersistentVolumeV1#cinder}
	Cinder *PersistentVolumeV1SpecPersistentVolumeSourceCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#csi PersistentVolumeV1#csi}
	Csi *PersistentVolumeV1SpecPersistentVolumeSourceCsi `field:"optional" json:"csi" yaml:"csi"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#fc PersistentVolumeV1#fc}
	Fc *PersistentVolumeV1SpecPersistentVolumeSourceFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#flex_volume PersistentVolumeV1#flex_volume}
	FlexVolume *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#flocker PersistentVolumeV1#flocker}
	Flocker *PersistentVolumeV1SpecPersistentVolumeSourceFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#gce_persistent_disk PersistentVolumeV1#gce_persistent_disk}
	GcePersistentDisk *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#glusterfs PersistentVolumeV1#glusterfs}
	Glusterfs *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#host_path PersistentVolumeV1#host_path}
	HostPath *PersistentVolumeV1SpecPersistentVolumeSourceHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#iscsi PersistentVolumeV1#iscsi}
	Iscsi *PersistentVolumeV1SpecPersistentVolumeSourceIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#local PersistentVolumeV1#local}
	Local *PersistentVolumeV1SpecPersistentVolumeSourceLocal `field:"optional" json:"local" yaml:"local"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#nfs PersistentVolumeV1#nfs}
	Nfs *PersistentVolumeV1SpecPersistentVolumeSourceNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#photon_persistent_disk PersistentVolumeV1#photon_persistent_disk}
	PhotonPersistentDisk *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#quobyte PersistentVolumeV1#quobyte}
	Quobyte *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#rbd PersistentVolumeV1#rbd}
	Rbd *PersistentVolumeV1SpecPersistentVolumeSourceRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/persistent_volume_v1#vsphere_volume PersistentVolumeV1#vsphere_volume}
	VsphereVolume *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

