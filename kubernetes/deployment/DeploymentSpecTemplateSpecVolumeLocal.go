package deployment


type DeploymentSpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/deployment#path Deployment#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

