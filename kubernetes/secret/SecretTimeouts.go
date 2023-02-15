package secret


type SecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/secret#create Secret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

