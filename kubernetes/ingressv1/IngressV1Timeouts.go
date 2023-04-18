package ingressv1


type IngressV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/ingress_v1#create IngressV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/ingress_v1#delete IngressV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

