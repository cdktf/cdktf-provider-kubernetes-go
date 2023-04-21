package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#name StatefulSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#optional StatefulSetV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

