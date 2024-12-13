// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolume struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#aws_elastic_block_store CronJobV1#aws_elastic_block_store}
	AwsElasticBlockStore *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#azure_disk CronJobV1#azure_disk}
	AzureDisk *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#azure_file CronJobV1#azure_file}
	AzureFile *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#ceph_fs CronJobV1#ceph_fs}
	CephFs *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#cinder CronJobV1#cinder}
	Cinder *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#config_map CronJobV1#config_map}
	ConfigMap *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#csi CronJobV1#csi}
	Csi *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi `field:"optional" json:"csi" yaml:"csi"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#downward_api CronJobV1#downward_api}
	DownwardApi *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#empty_dir CronJobV1#empty_dir}
	EmptyDir *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// ephemeral block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#ephemeral CronJobV1#ephemeral}
	Ephemeral *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#fc CronJobV1#fc}
	Fc *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#flex_volume CronJobV1#flex_volume}
	FlexVolume *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#flocker CronJobV1#flocker}
	Flocker *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#gce_persistent_disk CronJobV1#gce_persistent_disk}
	GcePersistentDisk *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// git_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#git_repo CronJobV1#git_repo}
	GitRepo *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#glusterfs CronJobV1#glusterfs}
	Glusterfs *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#host_path CronJobV1#host_path}
	HostPath *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#iscsi CronJobV1#iscsi}
	Iscsi *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#local CronJobV1#local}
	Local *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal `field:"optional" json:"local" yaml:"local"`
	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#name CronJobV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#nfs CronJobV1#nfs}
	Nfs *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// persistent_volume_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#persistent_volume_claim CronJobV1#persistent_volume_claim}
	PersistentVolumeClaim *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#photon_persistent_disk CronJobV1#photon_persistent_disk}
	PhotonPersistentDisk *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// projected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#projected CronJobV1#projected}
	Projected interface{} `field:"optional" json:"projected" yaml:"projected"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#quobyte CronJobV1#quobyte}
	Quobyte *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#rbd CronJobV1#rbd}
	Rbd *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#secret CronJobV1#secret}
	Secret *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret `field:"optional" json:"secret" yaml:"secret"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job_v1#vsphere_volume CronJobV1#vsphere_volume}
	VsphereVolume *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

