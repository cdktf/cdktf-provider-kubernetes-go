package manifest


type ManifestFieldManager struct {
	// Force changes against conflicts.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#force_conflicts Manifest#force_conflicts}
	ForceConflicts interface{} `field:"optional" json:"forceConflicts" yaml:"forceConflicts"`
	// The name to use for the field manager when creating and updating the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#name Manifest#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

