// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecToleration struct {
	// Effect indicates the taint effect to match.
	//
	// Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/daemonset#effect Daemonset#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Key is the taint key that the toleration applies to.
	//
	// Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/daemonset#key Daemonset#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Operator represents a key's relationship to the value.
	//
	// Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/daemonset#operator Daemonset#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint.
	//
	// By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/daemonset#toleration_seconds Daemonset#toleration_seconds}
	TolerationSeconds *string `field:"optional" json:"tolerationSeconds" yaml:"tolerationSeconds"`
	// Value is the taint value the toleration matches to.
	//
	// If the operator is Exists, the value should be empty, otherwise just a regular string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/daemonset#value Daemonset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

