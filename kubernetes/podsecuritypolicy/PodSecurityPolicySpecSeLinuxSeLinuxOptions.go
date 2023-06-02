package podsecuritypolicy


type PodSecurityPolicySpecSeLinuxSeLinuxOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod_security_policy#level PodSecurityPolicy#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod_security_policy#role PodSecurityPolicy#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod_security_policy#type PodSecurityPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod_security_policy#user PodSecurityPolicy#user}.
	User *string `field:"required" json:"user" yaml:"user"`
}

