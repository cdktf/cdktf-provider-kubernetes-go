package podsecuritypolicy


type PodSecurityPolicyMetadata struct {
	// An unstructured key value map stored with the podsecuritypolicy that may be used to store arbitrary metadata.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_security_policy#annotations PodSecurityPolicy#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the podsecuritypolicy.
	//
	// May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_security_policy#labels PodSecurityPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the podsecuritypolicy, must be unique. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_security_policy#name PodSecurityPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

