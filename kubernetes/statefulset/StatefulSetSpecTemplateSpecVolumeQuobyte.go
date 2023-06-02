package statefulset


type StatefulSetSpecTemplateSpecVolumeQuobyte struct {
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#registry StatefulSet#registry}
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	// Volume is a string that references an already created Quobyte volume by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#volume StatefulSet#volume}
	Volume *string `field:"required" json:"volume" yaml:"volume"`
	// Group to map volume access to Default is no group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#group StatefulSet#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#read_only StatefulSet#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// User to map volume access to Defaults to serivceaccount user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#user StatefulSet#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

