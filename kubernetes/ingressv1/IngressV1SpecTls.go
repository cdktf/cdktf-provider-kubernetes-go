// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressv1


type IngressV1SpecTls struct {
	// hosts is a list of hosts included in the TLS certificate.
	//
	// The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/ingress_v1#hosts IngressV1#hosts}
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// secretName is the name of the secret used to terminate TLS traffic on port 443.
	//
	// Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the "Host" header is used for routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/ingress_v1#secret_name IngressV1#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

