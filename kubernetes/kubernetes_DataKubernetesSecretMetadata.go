// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DataKubernetesSecretMetadata struct {
	// An unstructured key value map stored with the secret that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/secret#annotations DataKubernetesSecret#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the secret.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/secret#labels DataKubernetesSecret#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the secret, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/secret#name DataKubernetesSecret#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace defines the space within which name of the secret must be unique.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/secret#namespace DataKubernetesSecret#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

