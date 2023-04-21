package datakubernetesresources

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesResourcesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The resource apiVersion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#api_version DataKubernetesResources#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The resource kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#kind DataKubernetesResources#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// A selector to restrict the list of returned objects by their fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#field_selector DataKubernetesResources#field_selector}
	FieldSelector *string `field:"optional" json:"fieldSelector" yaml:"fieldSelector"`
	// A selector to restrict the list of returned objects by their labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#label_selector DataKubernetesResources#label_selector}
	LabelSelector *string `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// Limit is a maximum number of responses to return for a list call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#limit DataKubernetesResources#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The resource namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#namespace DataKubernetesResources#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The response from the API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/resources#objects DataKubernetesResources#objects}
	Objects *map[string]interface{} `field:"optional" json:"objects" yaml:"objects"`
}

