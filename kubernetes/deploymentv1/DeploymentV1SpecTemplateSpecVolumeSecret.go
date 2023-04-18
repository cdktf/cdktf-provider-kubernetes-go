package deploymentv1


type DeploymentV1SpecTemplateSpecVolumeSecret struct {
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment_v1#default_mode DeploymentV1#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment_v1#items DeploymentV1#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Optional: Specify whether the Secret or its keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment_v1#optional DeploymentV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
	// Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment_v1#secret_name DeploymentV1#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

