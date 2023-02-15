package deployment


type DeploymentSpecTemplateSpecInitContainerEnvValueFrom struct {
	// config_map_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#config_map_key_ref Deployment#config_map_key_ref}
	ConfigMapKeyRef *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#field_ref Deployment#field_ref}
	FieldRef *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#resource_field_ref Deployment#resource_field_ref}
	ResourceFieldRef *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#secret_key_ref Deployment#secret_key_ref}
	SecretKeyRef *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

