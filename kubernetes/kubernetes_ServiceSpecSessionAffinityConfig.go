// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ServiceSpecSessionAffinityConfig struct {
	// client_ip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service#client_ip Service#client_ip}
	ClientIp *ServiceSpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

