// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolume struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#aws_elastic_block_store Pod#aws_elastic_block_store}
	AwsElasticBlockStore *PodSpecVolumeAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#azure_disk Pod#azure_disk}
	AzureDisk *PodSpecVolumeAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#azure_file Pod#azure_file}
	AzureFile *PodSpecVolumeAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#ceph_fs Pod#ceph_fs}
	CephFs *PodSpecVolumeCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#cinder Pod#cinder}
	Cinder *PodSpecVolumeCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#config_map Pod#config_map}
	ConfigMap *PodSpecVolumeConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#csi Pod#csi}
	Csi *PodSpecVolumeCsi `field:"optional" json:"csi" yaml:"csi"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#downward_api Pod#downward_api}
	DownwardApi *PodSpecVolumeDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#empty_dir Pod#empty_dir}
	EmptyDir *PodSpecVolumeEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// ephemeral block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#ephemeral Pod#ephemeral}
	Ephemeral *PodSpecVolumeEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#fc Pod#fc}
	Fc *PodSpecVolumeFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#flex_volume Pod#flex_volume}
	FlexVolume *PodSpecVolumeFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#flocker Pod#flocker}
	Flocker *PodSpecVolumeFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#gce_persistent_disk Pod#gce_persistent_disk}
	GcePersistentDisk *PodSpecVolumeGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// git_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#git_repo Pod#git_repo}
	GitRepo *PodSpecVolumeGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#glusterfs Pod#glusterfs}
	Glusterfs *PodSpecVolumeGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#host_path Pod#host_path}
	HostPath *PodSpecVolumeHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#iscsi Pod#iscsi}
	Iscsi *PodSpecVolumeIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#local Pod#local}
	Local *PodSpecVolumeLocal `field:"optional" json:"local" yaml:"local"`
	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#name Pod#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#nfs Pod#nfs}
	Nfs *PodSpecVolumeNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// persistent_volume_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#persistent_volume_claim Pod#persistent_volume_claim}
	PersistentVolumeClaim *PodSpecVolumePersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#photon_persistent_disk Pod#photon_persistent_disk}
	PhotonPersistentDisk *PodSpecVolumePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// projected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#projected Pod#projected}
	Projected interface{} `field:"optional" json:"projected" yaml:"projected"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#quobyte Pod#quobyte}
	Quobyte *PodSpecVolumeQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#rbd Pod#rbd}
	Rbd *PodSpecVolumeRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#secret Pod#secret}
	Secret *PodSpecVolumeSecret `field:"optional" json:"secret" yaml:"secret"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#vsphere_volume Pod#vsphere_volume}
	VsphereVolume *PodSpecVolumeVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

