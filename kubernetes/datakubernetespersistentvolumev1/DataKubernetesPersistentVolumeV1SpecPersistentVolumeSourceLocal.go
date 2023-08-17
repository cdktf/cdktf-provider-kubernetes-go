package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/data-sources/persistent_volume_v1#path DataKubernetesPersistentVolumeV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

