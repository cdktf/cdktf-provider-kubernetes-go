package jobv1


type JobV1SpecTemplateSpecSecurityContextSeccompProfile struct {
	// Localhost Profile indicates a profile defined in a file on the node should be used.
	//
	// The profile must be preconfigured on the node to work.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#localhost_profile JobV1#localhost_profile}
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
	// Type indicates which kind of seccomp profile will be applied. Valid options are: Localhost, RuntimeDefault, Unconfined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#type JobV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

