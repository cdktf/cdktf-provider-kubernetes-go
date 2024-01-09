// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatesigningrequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateSigningRequestConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/certificate_signing_request#metadata CertificateSigningRequest#metadata}
	Metadata *CertificateSigningRequestMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/certificate_signing_request#spec CertificateSigningRequest#spec}
	Spec *CertificateSigningRequestSpec `field:"required" json:"spec" yaml:"spec"`
	// Automatically approve the CertificateSigningRequest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/certificate_signing_request#auto_approve CertificateSigningRequest#auto_approve}
	AutoApprove interface{} `field:"optional" json:"autoApprove" yaml:"autoApprove"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/certificate_signing_request#id CertificateSigningRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/certificate_signing_request#timeouts CertificateSigningRequest#timeouts}
	Timeouts *CertificateSigningRequestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

