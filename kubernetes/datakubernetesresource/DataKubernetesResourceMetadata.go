package datakubernetesresource


type DataKubernetesResourceMetadata struct {
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/data-sources/resource#name DataKubernetesResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/data-sources/resource#namespace DataKubernetesResource#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

