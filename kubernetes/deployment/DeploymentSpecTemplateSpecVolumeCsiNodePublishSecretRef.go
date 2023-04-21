package deployment


type DeploymentSpecTemplateSpecVolumeCsiNodePublishSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#name Deployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

