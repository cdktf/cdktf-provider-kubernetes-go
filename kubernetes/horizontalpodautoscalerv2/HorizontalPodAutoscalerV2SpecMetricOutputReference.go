// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/horizontalpodautoscalerv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HorizontalPodAutoscalerV2SpecMetricOutputReference interface {
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
	ContainerResource() HorizontalPodAutoscalerV2SpecMetricContainerResourceOutputReference
	ContainerResourceInput() *HorizontalPodAutoscalerV2SpecMetricContainerResource
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	External() HorizontalPodAutoscalerV2SpecMetricExternalOutputReference
	ExternalInput() *HorizontalPodAutoscalerV2SpecMetricExternal
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Object() HorizontalPodAutoscalerV2SpecMetricObjectOutputReference
	ObjectInput() *HorizontalPodAutoscalerV2SpecMetricObject
	Pods() HorizontalPodAutoscalerV2SpecMetricPodsOutputReference
	PodsInput() *HorizontalPodAutoscalerV2SpecMetricPods
	Resource() HorizontalPodAutoscalerV2SpecMetricResourceOutputReference
	ResourceInput() *HorizontalPodAutoscalerV2SpecMetricResource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutContainerResource(value *HorizontalPodAutoscalerV2SpecMetricContainerResource)
	PutExternal(value *HorizontalPodAutoscalerV2SpecMetricExternal)
	PutObject(value *HorizontalPodAutoscalerV2SpecMetricObject)
	PutPods(value *HorizontalPodAutoscalerV2SpecMetricPods)
	PutResource(value *HorizontalPodAutoscalerV2SpecMetricResource)
	ResetContainerResource()
	ResetExternal()
	ResetObject()
	ResetPods()
	ResetResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerV2SpecMetricOutputReference
type jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ContainerResource() HorizontalPodAutoscalerV2SpecMetricContainerResourceOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricContainerResourceOutputReference
	_jsii_.Get(
		j,
		"containerResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ContainerResourceInput() *HorizontalPodAutoscalerV2SpecMetricContainerResource {
	var returns *HorizontalPodAutoscalerV2SpecMetricContainerResource
	_jsii_.Get(
		j,
		"containerResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) External() HorizontalPodAutoscalerV2SpecMetricExternalOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricExternalOutputReference
	_jsii_.Get(
		j,
		"external",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ExternalInput() *HorizontalPodAutoscalerV2SpecMetricExternal {
	var returns *HorizontalPodAutoscalerV2SpecMetricExternal
	_jsii_.Get(
		j,
		"externalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Object() HorizontalPodAutoscalerV2SpecMetricObjectOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricObjectOutputReference
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ObjectInput() *HorizontalPodAutoscalerV2SpecMetricObject {
	var returns *HorizontalPodAutoscalerV2SpecMetricObject
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Pods() HorizontalPodAutoscalerV2SpecMetricPodsOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricPodsOutputReference
	_jsii_.Get(
		j,
		"pods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PodsInput() *HorizontalPodAutoscalerV2SpecMetricPods {
	var returns *HorizontalPodAutoscalerV2SpecMetricPods
	_jsii_.Get(
		j,
		"podsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Resource() HorizontalPodAutoscalerV2SpecMetricResourceOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricResourceOutputReference
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResourceInput() *HorizontalPodAutoscalerV2SpecMetricResource {
	var returns *HorizontalPodAutoscalerV2SpecMetricResource
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerV2SpecMetricOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HorizontalPodAutoscalerV2SpecMetricOutputReference {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerV2SpecMetricOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerV2SpecMetricOutputReference_Override(h HorizontalPodAutoscalerV2SpecMetricOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PutContainerResource(value *HorizontalPodAutoscalerV2SpecMetricContainerResource) {
	if err := h.validatePutContainerResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putContainerResource",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PutExternal(value *HorizontalPodAutoscalerV2SpecMetricExternal) {
	if err := h.validatePutExternalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExternal",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PutObject(value *HorizontalPodAutoscalerV2SpecMetricObject) {
	if err := h.validatePutObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putObject",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PutPods(value *HorizontalPodAutoscalerV2SpecMetricPods) {
	if err := h.validatePutPodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putPods",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) PutResource(value *HorizontalPodAutoscalerV2SpecMetricResource) {
	if err := h.validatePutResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putResource",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResetContainerResource() {
	_jsii_.InvokeVoid(
		h,
		"resetContainerResource",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResetExternal() {
	_jsii_.InvokeVoid(
		h,
		"resetExternal",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResetObject() {
	_jsii_.InvokeVoid(
		h,
		"resetObject",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResetPods() {
	_jsii_.InvokeVoid(
		h,
		"resetPods",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		h,
		"resetResource",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

