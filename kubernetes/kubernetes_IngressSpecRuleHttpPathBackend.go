// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type IngressSpecRuleHttpPathBackend struct {
	// Specifies the name of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress#service_name Ingress#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Specifies the port of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress#service_port Ingress#service_port}
	ServicePort *string `field:"optional" json:"servicePort" yaml:"servicePort"`
}

