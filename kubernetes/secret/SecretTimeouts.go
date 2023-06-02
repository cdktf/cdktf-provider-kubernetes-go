package secret


type SecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/secret#create Secret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

