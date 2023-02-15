package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeGitRepo struct {
	// Target directory name.
	//
	// Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#directory StatefulSetV1#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#repository StatefulSetV1#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Commit hash for the specified revision.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#revision StatefulSetV1#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

