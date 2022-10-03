package persistentvolume


type PersistentVolumeSpecClaimRef struct {
	// The name of the PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#name PersistentVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the PersistentVolumeClaim. Uses 'default' namespace if none is specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#namespace PersistentVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

