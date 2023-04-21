package deployment


type DeploymentSpecTemplateSpecContainerEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#config_map_ref Deployment#config_map_ref}
	ConfigMapRef *DeploymentSpecTemplateSpecContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#prefix Deployment#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#secret_ref Deployment#secret_ref}
	SecretRef *DeploymentSpecTemplateSpecContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

