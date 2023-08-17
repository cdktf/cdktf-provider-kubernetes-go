package service


type ServiceSpecSessionAffinityConfig struct {
	// client_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/service#client_ip Service#client_ip}
	ClientIp *ServiceSpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

