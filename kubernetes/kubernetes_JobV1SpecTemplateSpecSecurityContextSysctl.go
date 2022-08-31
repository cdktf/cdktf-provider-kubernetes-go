// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobV1SpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#name JobV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#value JobV1#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

