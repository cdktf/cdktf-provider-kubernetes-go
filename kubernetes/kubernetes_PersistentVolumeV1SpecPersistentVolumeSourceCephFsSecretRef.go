// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#name PersistentVolumeV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#namespace PersistentVolumeV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

