package ingressv1


type IngressV1SpecRuleHttp struct {
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/ingress_v1#path IngressV1#path}
	Path interface{} `field:"required" json:"path" yaml:"path"`
}

