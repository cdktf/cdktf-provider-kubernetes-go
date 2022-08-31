// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type RoleBindingSubject struct {
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#kind RoleBinding#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the resource to bind to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#name RoleBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API group of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#api_group RoleBinding#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// The Namespace of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/role_binding#namespace RoleBinding#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

