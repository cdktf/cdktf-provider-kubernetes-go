package statefulset


type StatefulSetSpecTemplateSpecVolumeGitRepo struct {
	// Target directory name.
	//
	// Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#directory StatefulSet#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#repository StatefulSet#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Commit hash for the specified revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#revision StatefulSet#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

