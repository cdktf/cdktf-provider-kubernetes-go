// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobSpecTemplateSpecVolumeCsiNodePublishSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#name Job#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}
