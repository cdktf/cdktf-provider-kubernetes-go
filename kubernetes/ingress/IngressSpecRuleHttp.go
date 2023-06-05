package ingress


type IngressSpecRuleHttp struct {
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/ingress#path Ingress#path}
	Path interface{} `field:"required" json:"path" yaml:"path"`
}

