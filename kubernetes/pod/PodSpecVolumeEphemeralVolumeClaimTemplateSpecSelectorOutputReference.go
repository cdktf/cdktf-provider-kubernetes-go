// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/pod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference interface {
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
	InternalValue() *PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelector
	SetInternalValue(val *PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelector)
	MatchExpressions() PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorMatchExpressionsList
	MatchExpressionsInput() interface{}
	MatchLabels() *map[string]*string
	SetMatchLabels(val *map[string]*string)
	MatchLabelsInput() *map[string]*string
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
	PutMatchExpressions(value interface{})
	ResetMatchExpressions()
	ResetMatchLabels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference
type jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) InternalValue() *PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelector {
	var returns *PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) MatchExpressions() PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorMatchExpressionsList {
	var returns PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorMatchExpressionsList
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) MatchExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchExpressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) MatchLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference_Override(p PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetInternalValue(val *PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetMatchLabels(val *map[string]*string) {
	if err := j.validateSetMatchLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabels",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) PutMatchExpressions(value interface{}) {
	if err := p.validatePutMatchExpressionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMatchExpressions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ResetMatchExpressions() {
	_jsii_.InvokeVoid(
		p,
		"resetMatchExpressions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ResetMatchLabels() {
	_jsii_.InvokeVoid(
		p,
		"resetMatchLabels",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

