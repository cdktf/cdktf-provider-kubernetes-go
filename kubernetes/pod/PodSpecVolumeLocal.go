package pod


type PodSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#path Pod#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

