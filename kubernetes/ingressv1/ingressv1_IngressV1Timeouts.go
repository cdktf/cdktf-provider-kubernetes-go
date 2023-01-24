package ingressv1


type IngressV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#create IngressV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#delete IngressV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

