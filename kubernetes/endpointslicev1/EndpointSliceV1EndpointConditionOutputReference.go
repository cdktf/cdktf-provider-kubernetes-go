// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointslicev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/endpointslicev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EndpointSliceV1EndpointConditionOutputReference interface {
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
	InternalValue() *EndpointSliceV1EndpointCondition
	SetInternalValue(val *EndpointSliceV1EndpointCondition)
	Ready() interface{}
	SetReady(val interface{})
	ReadyInput() interface{}
	Serving() interface{}
	SetServing(val interface{})
	ServingInput() interface{}
	Terminating() interface{}
	SetTerminating(val interface{})
	TerminatingInput() interface{}
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
	ResetReady()
	ResetServing()
	ResetTerminating()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EndpointSliceV1EndpointConditionOutputReference
type jsiiProxy_EndpointSliceV1EndpointConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) InternalValue() *EndpointSliceV1EndpointCondition {
	var returns *EndpointSliceV1EndpointCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) Ready() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ReadyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) Serving() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serving",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ServingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) Terminating() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminating",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) TerminatingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminatingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEndpointSliceV1EndpointConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EndpointSliceV1EndpointConditionOutputReference {
	_init_.Initialize()

	if err := validateNewEndpointSliceV1EndpointConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EndpointSliceV1EndpointConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.endpointSliceV1.EndpointSliceV1EndpointConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEndpointSliceV1EndpointConditionOutputReference_Override(e EndpointSliceV1EndpointConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.endpointSliceV1.EndpointSliceV1EndpointConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetInternalValue(val *EndpointSliceV1EndpointCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetReady(val interface{}) {
	if err := j.validateSetReadyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ready",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetServing(val interface{}) {
	if err := j.validateSetServingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serving",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetTerminating(val interface{}) {
	if err := j.validateSetTerminatingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminating",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ResetReady() {
	_jsii_.InvokeVoid(
		e,
		"resetReady",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ResetServing() {
	_jsii_.InvokeVoid(
		e,
		"resetServing",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ResetTerminating() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminating",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EndpointSliceV1EndpointConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

