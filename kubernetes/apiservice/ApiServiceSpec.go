package apiservice


type ApiServiceSpec struct {
	// Group is the API group name this server hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#group ApiService#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// GroupPriorityMinimum is the priority this group should have at least.
	//
	// Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#group_priority_minimum ApiService#group_priority_minimum}
	GroupPriorityMinimum *float64 `field:"required" json:"groupPriorityMinimum" yaml:"groupPriorityMinimum"`
	// Version is the API version this server hosts. For example, `v1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#version ApiService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// VersionPriority controls the ordering of this API version inside of its group.
	//
	// Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is `kube-like`, it will sort above non `kube-like` version strings, which are ordered lexicographically. `Kube-like` versions start with a `v`, then are followed by a number (the major version), then optionally the string `alpha` or `beta` and another number (the minor version). These are sorted first by GA > `beta` > `alpha` (where GA is a version with no suffix such as `beta` or `alpha`), and then by comparing major version, then minor version. An example sorted list of versions: `v10`, `v2`, `v1`, `v11beta2`, `v10beta3`, `v3beta1`, `v12alpha1`, `v11alpha2`, `foo1`, `foo10`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#version_priority ApiService#version_priority}
	VersionPriority *float64 `field:"required" json:"versionPriority" yaml:"versionPriority"`
	// CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate.
	//
	// If unspecified, system trust roots on the apiserver are used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#ca_bundle ApiService#ca_bundle}
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server.
	//
	// This is strongly discouraged. You should use the CABundle instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#insecure_skip_tls_verify ApiService#insecure_skip_tls_verify}
	InsecureSkipTlsVerify interface{} `field:"optional" json:"insecureSkipTlsVerify" yaml:"insecureSkipTlsVerify"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/api_service#service ApiService#service}
	Service *ApiServiceSpecService `field:"optional" json:"service" yaml:"service"`
}

