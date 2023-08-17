package statefulset


type StatefulSetSpecTemplateSpecContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#name StatefulSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#optional StatefulSet#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

