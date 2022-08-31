// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type IngressV1SpecDefaultBackend struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#resource IngressV1#resource}
	Resource *IngressV1SpecDefaultBackendResource `field:"optional" json:"resource" yaml:"resource"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#service IngressV1#service}
	Service *IngressV1SpecDefaultBackendService `field:"optional" json:"service" yaml:"service"`
}

