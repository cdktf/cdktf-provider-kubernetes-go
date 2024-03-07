// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecVolume struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#aws_elastic_block_store StatefulSet#aws_elastic_block_store}
	AwsElasticBlockStore *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#azure_disk StatefulSet#azure_disk}
	AzureDisk *StatefulSetSpecTemplateSpecVolumeAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#azure_file StatefulSet#azure_file}
	AzureFile *StatefulSetSpecTemplateSpecVolumeAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#ceph_fs StatefulSet#ceph_fs}
	CephFs *StatefulSetSpecTemplateSpecVolumeCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#cinder StatefulSet#cinder}
	Cinder *StatefulSetSpecTemplateSpecVolumeCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#config_map StatefulSet#config_map}
	ConfigMap *StatefulSetSpecTemplateSpecVolumeConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#csi StatefulSet#csi}
	Csi *StatefulSetSpecTemplateSpecVolumeCsi `field:"optional" json:"csi" yaml:"csi"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#downward_api StatefulSet#downward_api}
	DownwardApi *StatefulSetSpecTemplateSpecVolumeDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#empty_dir StatefulSet#empty_dir}
	EmptyDir *StatefulSetSpecTemplateSpecVolumeEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// ephemeral block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#ephemeral StatefulSet#ephemeral}
	Ephemeral *StatefulSetSpecTemplateSpecVolumeEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#fc StatefulSet#fc}
	Fc *StatefulSetSpecTemplateSpecVolumeFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#flex_volume StatefulSet#flex_volume}
	FlexVolume *StatefulSetSpecTemplateSpecVolumeFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#flocker StatefulSet#flocker}
	Flocker *StatefulSetSpecTemplateSpecVolumeFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#gce_persistent_disk StatefulSet#gce_persistent_disk}
	GcePersistentDisk *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// git_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#git_repo StatefulSet#git_repo}
	GitRepo *StatefulSetSpecTemplateSpecVolumeGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#glusterfs StatefulSet#glusterfs}
	Glusterfs *StatefulSetSpecTemplateSpecVolumeGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#host_path StatefulSet#host_path}
	HostPath *StatefulSetSpecTemplateSpecVolumeHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#iscsi StatefulSet#iscsi}
	Iscsi *StatefulSetSpecTemplateSpecVolumeIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#local StatefulSet#local}
	Local *StatefulSetSpecTemplateSpecVolumeLocal `field:"optional" json:"local" yaml:"local"`
	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#name StatefulSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#nfs StatefulSet#nfs}
	Nfs *StatefulSetSpecTemplateSpecVolumeNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// persistent_volume_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#persistent_volume_claim StatefulSet#persistent_volume_claim}
	PersistentVolumeClaim *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#photon_persistent_disk StatefulSet#photon_persistent_disk}
	PhotonPersistentDisk *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// projected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#projected StatefulSet#projected}
	Projected interface{} `field:"optional" json:"projected" yaml:"projected"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#quobyte StatefulSet#quobyte}
	Quobyte *StatefulSetSpecTemplateSpecVolumeQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#rbd StatefulSet#rbd}
	Rbd *StatefulSetSpecTemplateSpecVolumeRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#secret StatefulSet#secret}
	Secret *StatefulSetSpecTemplateSpecVolumeSecret `field:"optional" json:"secret" yaml:"secret"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/stateful_set#vsphere_volume StatefulSet#vsphere_volume}
	VsphereVolume *StatefulSetSpecTemplateSpecVolumeVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

