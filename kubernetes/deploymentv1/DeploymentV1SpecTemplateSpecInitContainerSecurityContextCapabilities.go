package deploymentv1


type DeploymentV1SpecTemplateSpecInitContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#add DeploymentV1#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#drop DeploymentV1#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

