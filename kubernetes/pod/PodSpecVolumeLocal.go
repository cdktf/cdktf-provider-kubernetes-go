package pod


type PodSpecVolumeLocal struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#path Pod#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

