// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#aws_elastic_block_store DataKubernetesPersistentVolumeV1#aws_elastic_block_store}
	AwsElasticBlockStore *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#azure_disk DataKubernetesPersistentVolumeV1#azure_disk}
	AzureDisk *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#azure_file DataKubernetesPersistentVolumeV1#azure_file}
	AzureFile *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#ceph_fs DataKubernetesPersistentVolumeV1#ceph_fs}
	CephFs *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#cinder DataKubernetesPersistentVolumeV1#cinder}
	Cinder *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#csi DataKubernetesPersistentVolumeV1#csi}
	Csi *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi `field:"optional" json:"csi" yaml:"csi"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#fc DataKubernetesPersistentVolumeV1#fc}
	Fc *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#flex_volume DataKubernetesPersistentVolumeV1#flex_volume}
	FlexVolume *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#flocker DataKubernetesPersistentVolumeV1#flocker}
	Flocker *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#gce_persistent_disk DataKubernetesPersistentVolumeV1#gce_persistent_disk}
	GcePersistentDisk *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#glusterfs DataKubernetesPersistentVolumeV1#glusterfs}
	Glusterfs *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#host_path DataKubernetesPersistentVolumeV1#host_path}
	HostPath *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#iscsi DataKubernetesPersistentVolumeV1#iscsi}
	Iscsi *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#local DataKubernetesPersistentVolumeV1#local}
	Local *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal `field:"optional" json:"local" yaml:"local"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#nfs DataKubernetesPersistentVolumeV1#nfs}
	Nfs *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#photon_persistent_disk DataKubernetesPersistentVolumeV1#photon_persistent_disk}
	PhotonPersistentDisk *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#quobyte DataKubernetesPersistentVolumeV1#quobyte}
	Quobyte *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#rbd DataKubernetesPersistentVolumeV1#rbd}
	Rbd *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/data-sources/persistent_volume_v1#vsphere_volume DataKubernetesPersistentVolumeV1#vsphere_volume}
	VsphereVolume *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

