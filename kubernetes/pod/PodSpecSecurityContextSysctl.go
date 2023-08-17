package pod


type PodSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#name Pod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#value Pod#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

