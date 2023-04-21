package networkpolicyv1


type NetworkPolicyV1Metadata struct {
	// An unstructured key value map stored with the network policy that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy_v1#annotations NetworkPolicyV1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Prefix, used by the server, to generate a unique name ONLY IF the `name` field has not been provided.
	//
	// This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy_v1#generate_name NetworkPolicyV1#generate_name}
	GenerateName *string `field:"optional" json:"generateName" yaml:"generateName"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the network policy.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy_v1#labels NetworkPolicyV1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the network policy, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy_v1#name NetworkPolicyV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace defines the space within which name of the network policy must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy_v1#namespace NetworkPolicyV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

