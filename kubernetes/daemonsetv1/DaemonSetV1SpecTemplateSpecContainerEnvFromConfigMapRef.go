package daemonsetv1


type DaemonSetV1SpecTemplateSpecContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemon_set_v1#optional DaemonSetV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

