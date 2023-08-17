package deploymentv1


type DeploymentV1SpecTemplateSpecVolumeConfigMap struct {
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment_v1#default_mode DeploymentV1#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment_v1#items DeploymentV1#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment_v1#name DeploymentV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the ConfigMap or its keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment_v1#optional DeploymentV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

