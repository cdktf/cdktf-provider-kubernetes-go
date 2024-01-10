// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/horizontalpodautoscalerv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HorizontalPodAutoscalerV2SpecBehaviorScaleUpList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) HorizontalPodAutoscalerV2SpecBehaviorScaleUpOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerV2SpecBehaviorScaleUpList
type jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerV2SpecBehaviorScaleUpList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HorizontalPodAutoscalerV2SpecBehaviorScaleUpList {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerV2SpecBehaviorScaleUpListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorScaleUpList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerV2SpecBehaviorScaleUpList_Override(h HorizontalPodAutoscalerV2SpecBehaviorScaleUpList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorScaleUpList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := h.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		h,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) Get(index *float64) HorizontalPodAutoscalerV2SpecBehaviorScaleUpOutputReference {
	if err := h.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns HorizontalPodAutoscalerV2SpecBehaviorScaleUpOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorScaleUpList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

