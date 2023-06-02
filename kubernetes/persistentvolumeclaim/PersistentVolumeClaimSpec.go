package persistentvolumeclaim


type PersistentVolumeClaimSpec struct {
	// A set of the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/persistent_volume_claim#access_modes PersistentVolumeClaim#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/persistent_volume_claim#resources PersistentVolumeClaim#resources}
	Resources *PersistentVolumeClaimSpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/persistent_volume_claim#selector PersistentVolumeClaim#selector}
	Selector *PersistentVolumeClaimSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/persistent_volume_claim#storage_class_name PersistentVolumeClaim#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/persistent_volume_claim#volume_name PersistentVolumeClaim#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

