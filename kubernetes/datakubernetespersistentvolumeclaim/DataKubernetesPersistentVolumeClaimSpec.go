package datakubernetespersistentvolumeclaim


type DataKubernetesPersistentVolumeClaimSpec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/persistent_volume_claim#selector DataKubernetesPersistentVolumeClaim#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/persistent_volume_claim#storage_class_name DataKubernetesPersistentVolumeClaim#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/persistent_volume_claim#volume_name DataKubernetesPersistentVolumeClaim#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

