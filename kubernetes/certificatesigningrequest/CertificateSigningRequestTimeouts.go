// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatesigningrequest


type CertificateSigningRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/certificate_signing_request#create CertificateSigningRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

