package daemonset


type DaemonsetSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#level Daemonset#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Role is a SELinux role label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#role Daemonset#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type is a SELinux type label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#type Daemonset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User is a SELinux user label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#user Daemonset#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

