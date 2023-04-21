package job


type JobSpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#path Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

