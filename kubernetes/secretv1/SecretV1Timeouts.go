package secretv1


type SecretV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/secret_v1#create SecretV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

