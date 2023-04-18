package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/persistent_volume_v1#name PersistentVolumeV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/persistent_volume_v1#namespace PersistentVolumeV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

