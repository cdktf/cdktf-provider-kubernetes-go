package jobv1


type JobV1SpecTemplateSpecVolume struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#aws_elastic_block_store JobV1#aws_elastic_block_store}
	AwsElasticBlockStore *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#azure_disk JobV1#azure_disk}
	AzureDisk *JobV1SpecTemplateSpecVolumeAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#azure_file JobV1#azure_file}
	AzureFile *JobV1SpecTemplateSpecVolumeAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#ceph_fs JobV1#ceph_fs}
	CephFs *JobV1SpecTemplateSpecVolumeCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#cinder JobV1#cinder}
	Cinder *JobV1SpecTemplateSpecVolumeCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#config_map JobV1#config_map}
	ConfigMap *JobV1SpecTemplateSpecVolumeConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#csi JobV1#csi}
	Csi *JobV1SpecTemplateSpecVolumeCsi `field:"optional" json:"csi" yaml:"csi"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#downward_api JobV1#downward_api}
	DownwardApi *JobV1SpecTemplateSpecVolumeDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// empty_dir block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#empty_dir JobV1#empty_dir}
	EmptyDir *JobV1SpecTemplateSpecVolumeEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#fc JobV1#fc}
	Fc *JobV1SpecTemplateSpecVolumeFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#flex_volume JobV1#flex_volume}
	FlexVolume *JobV1SpecTemplateSpecVolumeFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#flocker JobV1#flocker}
	Flocker *JobV1SpecTemplateSpecVolumeFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#gce_persistent_disk JobV1#gce_persistent_disk}
	GcePersistentDisk *JobV1SpecTemplateSpecVolumeGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// git_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#git_repo JobV1#git_repo}
	GitRepo *JobV1SpecTemplateSpecVolumeGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#glusterfs JobV1#glusterfs}
	Glusterfs *JobV1SpecTemplateSpecVolumeGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#host_path JobV1#host_path}
	HostPath *JobV1SpecTemplateSpecVolumeHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#iscsi JobV1#iscsi}
	Iscsi *JobV1SpecTemplateSpecVolumeIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#local JobV1#local}
	Local *JobV1SpecTemplateSpecVolumeLocal `field:"optional" json:"local" yaml:"local"`
	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#name JobV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#nfs JobV1#nfs}
	Nfs *JobV1SpecTemplateSpecVolumeNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// persistent_volume_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#persistent_volume_claim JobV1#persistent_volume_claim}
	PersistentVolumeClaim *JobV1SpecTemplateSpecVolumePersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#photon_persistent_disk JobV1#photon_persistent_disk}
	PhotonPersistentDisk *JobV1SpecTemplateSpecVolumePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// projected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#projected JobV1#projected}
	Projected interface{} `field:"optional" json:"projected" yaml:"projected"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#quobyte JobV1#quobyte}
	Quobyte *JobV1SpecTemplateSpecVolumeQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#rbd JobV1#rbd}
	Rbd *JobV1SpecTemplateSpecVolumeRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#secret JobV1#secret}
	Secret *JobV1SpecTemplateSpecVolumeSecret `field:"optional" json:"secret" yaml:"secret"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#vsphere_volume JobV1#vsphere_volume}
	VsphereVolume *JobV1SpecTemplateSpecVolumeVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

