// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/horizontalpodautoscalerv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HorizontalPodAutoscalerV2SpecBehaviorOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *HorizontalPodAutoscalerV2SpecBehavior
	SetInternalValue(val *HorizontalPodAutoscalerV2SpecBehavior)
	ScaleDown() HorizontalPodAutoscalerV2SpecBehaviorScaleDownList
	ScaleDownInput() interface{}
	ScaleUp() HorizontalPodAutoscalerV2SpecBehaviorScaleUpList
	ScaleUpInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutScaleDown(value interface{})
	PutScaleUp(value interface{})
	ResetScaleDown()
	ResetScaleUp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerV2SpecBehaviorOutputReference
type jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) InternalValue() *HorizontalPodAutoscalerV2SpecBehavior {
	var returns *HorizontalPodAutoscalerV2SpecBehavior
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ScaleDown() HorizontalPodAutoscalerV2SpecBehaviorScaleDownList {
	var returns HorizontalPodAutoscalerV2SpecBehaviorScaleDownList
	_jsii_.Get(
		j,
		"scaleDown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ScaleDownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleDownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ScaleUp() HorizontalPodAutoscalerV2SpecBehaviorScaleUpList {
	var returns HorizontalPodAutoscalerV2SpecBehaviorScaleUpList
	_jsii_.Get(
		j,
		"scaleUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ScaleUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerV2SpecBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HorizontalPodAutoscalerV2SpecBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerV2SpecBehaviorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerV2SpecBehaviorOutputReference_Override(h HorizontalPodAutoscalerV2SpecBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference)SetInternalValue(val *HorizontalPodAutoscalerV2SpecBehavior) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) PutScaleDown(value interface{}) {
	if err := h.validatePutScaleDownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putScaleDown",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) PutScaleUp(value interface{}) {
	if err := h.validatePutScaleUpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putScaleUp",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ResetScaleDown() {
	_jsii_.InvokeVoid(
		h,
		"resetScaleDown",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ResetScaleUp() {
	_jsii_.InvokeVoid(
		h,
		"resetScaleUp",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

