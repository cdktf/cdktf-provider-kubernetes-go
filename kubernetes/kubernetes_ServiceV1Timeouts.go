// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ServiceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service_v1#create ServiceV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

