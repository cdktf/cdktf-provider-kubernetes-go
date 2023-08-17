package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec struct {
	// A set of the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#access_modes ReplicationController#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#resources ReplicationController#resources}
	Resources *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#selector ReplicationController#selector}
	Selector *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#storage_class_name ReplicationController#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#volume_name ReplicationController#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

