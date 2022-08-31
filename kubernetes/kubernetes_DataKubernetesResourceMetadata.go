// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DataKubernetesResourceMetadata struct {
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#name DataKubernetesResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#namespace DataKubernetesResource#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

