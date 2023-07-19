package persistentvolumev1


type PersistentVolumeV1SpecClaimRef struct {
	// The name of the PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/persistent_volume_v1#name PersistentVolumeV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/persistent_volume_v1#namespace PersistentVolumeV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

