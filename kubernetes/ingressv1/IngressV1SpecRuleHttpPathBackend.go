package ingressv1


type IngressV1SpecRuleHttpPathBackend struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#resource IngressV1#resource}
	Resource *IngressV1SpecRuleHttpPathBackendResource `field:"optional" json:"resource" yaml:"resource"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#service IngressV1#service}
	Service *IngressV1SpecRuleHttpPathBackendService `field:"optional" json:"service" yaml:"service"`
}

