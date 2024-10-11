// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolume struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#aws_elastic_block_store CronJob#aws_elastic_block_store}
	AwsElasticBlockStore *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#azure_disk CronJob#azure_disk}
	AzureDisk *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#azure_file CronJob#azure_file}
	AzureFile *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#ceph_fs CronJob#ceph_fs}
	CephFs *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#cinder CronJob#cinder}
	Cinder *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#config_map CronJob#config_map}
	ConfigMap *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#csi CronJob#csi}
	Csi *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi `field:"optional" json:"csi" yaml:"csi"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#downward_api CronJob#downward_api}
	DownwardApi *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#empty_dir CronJob#empty_dir}
	EmptyDir *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// ephemeral block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#ephemeral CronJob#ephemeral}
	Ephemeral *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#fc CronJob#fc}
	Fc *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#flex_volume CronJob#flex_volume}
	FlexVolume *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#flocker CronJob#flocker}
	Flocker *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#gce_persistent_disk CronJob#gce_persistent_disk}
	GcePersistentDisk *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// git_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#git_repo CronJob#git_repo}
	GitRepo *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#glusterfs CronJob#glusterfs}
	Glusterfs *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#host_path CronJob#host_path}
	HostPath *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#iscsi CronJob#iscsi}
	Iscsi *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#local CronJob#local}
	Local *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal `field:"optional" json:"local" yaml:"local"`
	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#name CronJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#nfs CronJob#nfs}
	Nfs *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// persistent_volume_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#persistent_volume_claim CronJob#persistent_volume_claim}
	PersistentVolumeClaim *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#photon_persistent_disk CronJob#photon_persistent_disk}
	PhotonPersistentDisk *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// projected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#projected CronJob#projected}
	Projected interface{} `field:"optional" json:"projected" yaml:"projected"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#quobyte CronJob#quobyte}
	Quobyte *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#rbd CronJob#rbd}
	Rbd *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#secret CronJob#secret}
	Secret *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret `field:"optional" json:"secret" yaml:"secret"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#vsphere_volume CronJob#vsphere_volume}
	VsphereVolume *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

