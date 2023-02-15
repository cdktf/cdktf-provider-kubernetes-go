package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#name PersistentVolume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#namespace PersistentVolume#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

