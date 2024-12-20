// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatesigningrequest


type CertificateSigningRequestSpec struct {
	// Base64-encoded PKCS#10 CSR data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/certificate_signing_request#request CertificateSigningRequest#request}
	Request *string `field:"required" json:"request" yaml:"request"`
	// Requested signer for the request.
	//
	// It is a qualified name in the form: `scope-hostname.io/name`.If empty, it will be defaulted: 1. If it's a kubelet client certificate, it is assigned `kubernetes.io/kube-apiserver-client-kubelet`.2. If it's a kubelet serving certificate, it is assigned `kubernetes.io/kubelet-serving`.3. Otherwise, it is assigned `kubernetes.io/legacy-unknown`. Distribution of trust for signers happens out of band.You can select on this field using `spec.signerName`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/certificate_signing_request#signer_name CertificateSigningRequest#signer_name}
	SignerName *string `field:"optional" json:"signerName" yaml:"signerName"`
	// allowedUsages specifies a set of usage contexts the key will be valid for. See: 	https://tools.ietf.org/html/rfc5280#section-4.2.1.3 	https://tools.ietf.org/html/rfc5280#section-4.2.1.12.
	//
	// Valid values are:
	//  "signing",
	//  "digital signature",
	//  "content commitment",
	//  "key encipherment",
	//  "key agreement",
	//  "data encipherment",
	//  "cert sign",
	//  "crl sign",
	//  "encipher only",
	//  "decipher only",
	//  "any",
	//  "server auth",
	//  "client auth",
	//  "code signing",
	//  "email protection",
	//  "s/mime",
	//  "ipsec end system",
	//  "ipsec tunnel",
	//  "ipsec user",
	//  "timestamping",
	//  "ocsp signing",
	//  "microsoft sgc",
	//  "netscape sgc"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/certificate_signing_request#usages CertificateSigningRequest#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
}

