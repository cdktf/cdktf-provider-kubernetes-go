package statefulset


type StatefulSetSpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#name StatefulSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#value StatefulSet#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

