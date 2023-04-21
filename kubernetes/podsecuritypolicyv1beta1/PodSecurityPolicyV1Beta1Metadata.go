package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1Metadata struct {
	// An unstructured key value map stored with the podsecuritypolicy that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy_v1beta1#annotations PodSecurityPolicyV1Beta1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the podsecuritypolicy.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy_v1beta1#labels PodSecurityPolicyV1Beta1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the podsecuritypolicy, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy_v1beta1#name PodSecurityPolicyV1Beta1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

