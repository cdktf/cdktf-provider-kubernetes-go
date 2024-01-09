// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolebinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterRoleBindingConfig struct {
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
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cluster_role_binding#metadata ClusterRoleBinding#metadata}
	Metadata *ClusterRoleBindingMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// role_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cluster_role_binding#role_ref ClusterRoleBinding#role_ref}
	RoleRef *ClusterRoleBindingRoleRef `field:"required" json:"roleRef" yaml:"roleRef"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cluster_role_binding#subject ClusterRoleBinding#subject}
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cluster_role_binding#id ClusterRoleBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

