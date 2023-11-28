// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecToleration struct {
	// Effect indicates the taint effect to match.
	//
	// Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#effect ReplicationController#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Key is the taint key that the toleration applies to.
	//
	// Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#key ReplicationController#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Operator represents a key's relationship to the value.
	//
	// Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#operator ReplicationController#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint.
	//
	// By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#toleration_seconds ReplicationController#toleration_seconds}
	TolerationSeconds *string `field:"optional" json:"tolerationSeconds" yaml:"tolerationSeconds"`
	// Value is the taint value the toleration matches to.
	//
	// If the operator is Exists, the value should be empty, otherwise just a regular string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#value ReplicationController#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

