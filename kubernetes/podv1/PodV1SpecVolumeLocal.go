package podv1


type PodV1SpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod_v1#path PodV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

