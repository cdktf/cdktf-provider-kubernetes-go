package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#name StatefulSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#optional StatefulSetV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

