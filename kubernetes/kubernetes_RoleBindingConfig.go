// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleBindingConfig struct {
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
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#metadata RoleBinding#metadata}
	Metadata *RoleBindingMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// role_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#role_ref RoleBinding#role_ref}
	RoleRef *RoleBindingRoleRef `field:"required" json:"roleRef" yaml:"roleRef"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#subject RoleBinding#subject}
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#id RoleBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

