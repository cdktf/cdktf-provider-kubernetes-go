package persistentvolumev1


type PersistentVolumeV1SpecNodeAffinity struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume_v1#required PersistentVolumeV1#required}
	Required *PersistentVolumeV1SpecNodeAffinityRequired `field:"optional" json:"required" yaml:"required"`
}

