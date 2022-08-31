// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesResourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#api_version DataKubernetesResource#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The resource kind.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#kind DataKubernetesResource#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#metadata DataKubernetesResource#metadata}
	Metadata *DataKubernetesResourceMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// The response from the API server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/resource#object DataKubernetesResource#object}
	Object *map[string]interface{} `field:"optional" json:"object" yaml:"object"`
}

