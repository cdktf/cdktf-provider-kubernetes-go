// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type IngressV1SpecRuleHttp struct {
	// path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#path IngressV1#path}
	Path interface{} `field:"required" json:"path" yaml:"path"`
}

