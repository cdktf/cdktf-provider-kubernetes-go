// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PersistentVolumeV1SpecClaimRef struct {
	// The name of the PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#name PersistentVolumeV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#namespace PersistentVolumeV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

