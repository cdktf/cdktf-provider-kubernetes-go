// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingress


type IngressSpecRuleHttpPath struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/ingress#backend Ingress#backend}
	Backend *IngressSpecRuleHttpPathBackend `field:"optional" json:"backend" yaml:"backend"`
	// path is matched against the path of an incoming request.
	//
	// Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value "Exact" or "Prefix".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/ingress#path Ingress#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

