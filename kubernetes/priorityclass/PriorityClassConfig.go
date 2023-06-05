package priorityclass

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PriorityClassConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#metadata PriorityClass#metadata}
	Metadata *PriorityClassMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// The value of this priority class.
	//
	// This is the actual priority that pods receive when they have the name of this class in their pod spec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#value PriorityClass#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// An arbitrary string that usually provides guidelines on when this priority class should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#description PriorityClass#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class.
	//
	// Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#global_default PriorityClass#global_default}
	GlobalDefault interface{} `field:"optional" json:"globalDefault" yaml:"globalDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#id PriorityClass#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/priority_class#preemption_policy PriorityClass#preemption_policy}
	PreemptionPolicy *string `field:"optional" json:"preemptionPolicy" yaml:"preemptionPolicy"`
}

