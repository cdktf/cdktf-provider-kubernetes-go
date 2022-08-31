// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ServiceV1SpecSessionAffinityConfig struct {
	// client_ip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service_v1#client_ip ServiceV1#client_ip}
	ClientIp *ServiceV1SpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

