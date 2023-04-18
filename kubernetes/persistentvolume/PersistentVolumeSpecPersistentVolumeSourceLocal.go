package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourceLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/persistent_volume#path PersistentVolume#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

