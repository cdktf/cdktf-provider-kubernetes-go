package persistentvolume


type PersistentVolumeSpecPersistentVolumeSource struct {
	// aws_elastic_block_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#aws_elastic_block_store PersistentVolume#aws_elastic_block_store}
	AwsElasticBlockStore *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azure_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#azure_disk PersistentVolume#azure_disk}
	AzureDisk *PersistentVolumeSpecPersistentVolumeSourceAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azure_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#azure_file PersistentVolume#azure_file}
	AzureFile *PersistentVolumeSpecPersistentVolumeSourceAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	// ceph_fs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#ceph_fs PersistentVolume#ceph_fs}
	CephFs *PersistentVolumeSpecPersistentVolumeSourceCephFs `field:"optional" json:"cephFs" yaml:"cephFs"`
	// cinder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#cinder PersistentVolume#cinder}
	Cinder *PersistentVolumeSpecPersistentVolumeSourceCinder `field:"optional" json:"cinder" yaml:"cinder"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#csi PersistentVolume#csi}
	Csi *PersistentVolumeSpecPersistentVolumeSourceCsi `field:"optional" json:"csi" yaml:"csi"`
	// fc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#fc PersistentVolume#fc}
	Fc *PersistentVolumeSpecPersistentVolumeSourceFc `field:"optional" json:"fc" yaml:"fc"`
	// flex_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#flex_volume PersistentVolume#flex_volume}
	FlexVolume *PersistentVolumeSpecPersistentVolumeSourceFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#flocker PersistentVolume#flocker}
	Flocker *PersistentVolumeSpecPersistentVolumeSourceFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	// gce_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#gce_persistent_disk PersistentVolume#gce_persistent_disk}
	GcePersistentDisk *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// glusterfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#glusterfs PersistentVolume#glusterfs}
	Glusterfs *PersistentVolumeSpecPersistentVolumeSourceGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// host_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#host_path PersistentVolume#host_path}
	HostPath *PersistentVolumeSpecPersistentVolumeSourceHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#iscsi PersistentVolume#iscsi}
	Iscsi *PersistentVolumeSpecPersistentVolumeSourceIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#local PersistentVolume#local}
	Local *PersistentVolumeSpecPersistentVolumeSourceLocal `field:"optional" json:"local" yaml:"local"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#nfs PersistentVolume#nfs}
	Nfs *PersistentVolumeSpecPersistentVolumeSourceNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// photon_persistent_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#photon_persistent_disk PersistentVolume#photon_persistent_disk}
	PhotonPersistentDisk *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// quobyte block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#quobyte PersistentVolume#quobyte}
	Quobyte *PersistentVolumeSpecPersistentVolumeSourceQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#rbd PersistentVolume#rbd}
	Rbd *PersistentVolumeSpecPersistentVolumeSourceRbd `field:"optional" json:"rbd" yaml:"rbd"`
	// vsphere_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#vsphere_volume PersistentVolume#vsphere_volume}
	VsphereVolume *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

