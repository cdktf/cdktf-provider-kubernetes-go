package deploymentv1


type DeploymentV1SpecTemplateSpecImagePullSecrets struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment_v1#name DeploymentV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

