package namespace


type NamespaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/namespace#delete Namespace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

