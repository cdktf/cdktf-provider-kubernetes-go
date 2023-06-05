package jobv1


type JobV1SpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/job_v1#path JobV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

