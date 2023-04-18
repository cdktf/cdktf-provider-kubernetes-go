package job


type JobSpecTemplateSpecVolumeQuobyte struct {
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#registry Job#registry}
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	// Volume is a string that references an already created Quobyte volume by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#volume Job#volume}
	Volume *string `field:"required" json:"volume" yaml:"volume"`
	// Group to map volume access to Default is no group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#group Job#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#read_only Job#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// User to map volume access to Defaults to serivceaccount user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#user Job#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

