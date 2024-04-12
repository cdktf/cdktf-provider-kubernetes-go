// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicy


type PodSecurityPolicySpecAllowedHostPaths struct {
	// pathPrefix is the path prefix that the host volume must match.
	//
	// It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.
	//
	// Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/pod_security_policy#path_prefix PodSecurityPolicy#path_prefix}
	PathPrefix *string `field:"required" json:"pathPrefix" yaml:"pathPrefix"`
	// when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/pod_security_policy#read_only PodSecurityPolicy#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

