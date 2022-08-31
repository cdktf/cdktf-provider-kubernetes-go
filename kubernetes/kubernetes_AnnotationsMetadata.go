// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type AnnotationsMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/annotations#name Annotations#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/annotations#namespace Annotations#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

