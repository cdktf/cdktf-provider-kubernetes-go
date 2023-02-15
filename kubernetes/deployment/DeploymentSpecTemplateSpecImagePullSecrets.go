package deployment


type DeploymentSpecTemplateSpecImagePullSecrets struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#name Deployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

