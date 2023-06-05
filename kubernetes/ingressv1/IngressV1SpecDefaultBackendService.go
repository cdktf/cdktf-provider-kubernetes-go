package ingressv1


type IngressV1SpecDefaultBackendService struct {
	// Specifies the name of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/ingress_v1#name IngressV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/ingress_v1#port IngressV1#port}
	Port *IngressV1SpecDefaultBackendServicePort `field:"required" json:"port" yaml:"port"`
}

