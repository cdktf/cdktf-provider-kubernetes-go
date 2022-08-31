// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type IngressSpecRuleHttp struct {
	// path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress#path Ingress#path}
	Path interface{} `field:"required" json:"path" yaml:"path"`
}

