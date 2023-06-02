package deploymentv1


type DeploymentV1SpecTemplateSpecContainerEnvValueFrom struct {
	// config_map_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment_v1#config_map_key_ref DeploymentV1#config_map_key_ref}
	ConfigMapKeyRef *DeploymentV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment_v1#field_ref DeploymentV1#field_ref}
	FieldRef *DeploymentV1SpecTemplateSpecContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment_v1#resource_field_ref DeploymentV1#resource_field_ref}
	ResourceFieldRef *DeploymentV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment_v1#secret_key_ref DeploymentV1#secret_key_ref}
	SecretKeyRef *DeploymentV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

