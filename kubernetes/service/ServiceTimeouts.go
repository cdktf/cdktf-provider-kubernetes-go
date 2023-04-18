package service


type ServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/service#create Service#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

