// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSecurityPolicyV1Beta1SpecAllowedHostPaths struct {
	// pathPrefix is the path prefix that the host volume must match.
	//
	// It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.
	//
	// Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#path_prefix PodSecurityPolicyV1Beta1#path_prefix}
	PathPrefix *string `field:"required" json:"pathPrefix" yaml:"pathPrefix"`
	// when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#read_only PodSecurityPolicyV1Beta1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

