// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobV1SpecTemplateSpecVolumePersistentVolumeClaim struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#claim_name JobV1#claim_name}
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#read_only JobV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}
