// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PersistentVolumeV1SpecNodeAffinity struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#required PersistentVolumeV1#required}
	Required *PersistentVolumeV1SpecNodeAffinityRequired `field:"optional" json:"required" yaml:"required"`
}

