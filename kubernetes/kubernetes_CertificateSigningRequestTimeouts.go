// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CertificateSigningRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/certificate_signing_request#create CertificateSigningRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

