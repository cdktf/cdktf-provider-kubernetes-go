// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobSpecTemplateSpecInitContainerEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#config_map_ref Job#config_map_ref}
	ConfigMapRef *JobSpecTemplateSpecInitContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#prefix Job#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#secret_ref Job#secret_ref}
	SecretRef *JobSpecTemplateSpecInitContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}
