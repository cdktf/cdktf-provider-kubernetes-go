package servicev1


type ServiceV1SpecSessionAffinityConfig struct {
	// client_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/service_v1#client_ip ServiceV1#client_ip}
	ClientIp *ServiceV1SpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

