// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobSpecTemplateSpecInitContainerStartupProbeHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#name Job#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#value Job#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}
