// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ClusterRoleBindingV1RoleRef struct {
	// The API group of the user. The only value possible at the moment is `rbac.authorization.k8s.io`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cluster_role_binding_v1#api_group ClusterRoleBindingV1#api_group}
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cluster_role_binding_v1#kind ClusterRoleBindingV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the User to bind to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cluster_role_binding_v1#name ClusterRoleBindingV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}
