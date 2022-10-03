package persistentvolume


type PersistentVolumeSpecNodeAffinity struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#required PersistentVolume#required}
	Required *PersistentVolumeSpecNodeAffinityRequired `field:"optional" json:"required" yaml:"required"`
}

