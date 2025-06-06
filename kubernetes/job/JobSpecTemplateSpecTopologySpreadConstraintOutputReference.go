// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecTopologySpreadConstraintOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelSelector() JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	LabelSelectorInput() interface{}
	MatchLabelKeys() *[]*string
	SetMatchLabelKeys(val *[]*string)
	MatchLabelKeysInput() *[]*string
	MaxSkew() *float64
	SetMaxSkew(val *float64)
	MaxSkewInput() *float64
	MinDomains() *float64
	SetMinDomains(val *float64)
	MinDomainsInput() *float64
	NodeAffinityPolicy() *string
	SetNodeAffinityPolicy(val *string)
	NodeAffinityPolicyInput() *string
	NodeTaintsPolicy() *string
	SetNodeTaintsPolicy(val *string)
	NodeTaintsPolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopologyKey() *string
	SetTopologyKey(val *string)
	TopologyKeyInput() *string
	WhenUnsatisfiable() *string
	SetWhenUnsatisfiable(val *string)
	WhenUnsatisfiableInput() *string
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
	PutLabelSelector(value interface{})
	ResetLabelSelector()
	ResetMatchLabelKeys()
	ResetMaxSkew()
	ResetMinDomains()
	ResetNodeAffinityPolicy()
	ResetNodeTaintsPolicy()
	ResetTopologyKey()
	ResetWhenUnsatisfiable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobSpecTemplateSpecTopologySpreadConstraintOutputReference
type jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelector() JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList {
	var returns JobSpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomains() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomainsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiableInput",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecTopologySpreadConstraintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecTopologySpreadConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecTopologySpreadConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecTopologySpreadConstraintOutputReference_Override(j JobSpecTemplateSpecTopologySpreadConstraintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetMatchLabelKeys(val *[]*string) {
	if err := j.validateSetMatchLabelKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelKeys",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetMaxSkew(val *float64) {
	if err := j.validateSetMaxSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSkew",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetMinDomains(val *float64) {
	if err := j.validateSetMinDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDomains",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeAffinityPolicy(val *string) {
	if err := j.validateSetNodeAffinityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeAffinityPolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeTaintsPolicy(val *string) {
	if err := j.validateSetNodeTaintsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTaintsPolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetTopologyKey(val *string) {
	if err := j.validateSetTopologyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyKey",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference)SetWhenUnsatisfiable(val *string) {
	if err := j.validateSetWhenUnsatisfiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenUnsatisfiable",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) PutLabelSelector(value interface{}) {
	if err := j.validatePutLabelSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLabelSelector",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetLabelSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetLabelSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMatchLabelKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetMatchLabelKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMaxSkew() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxSkew",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMinDomains() {
	_jsii_.InvokeVoid(
		j,
		"resetMinDomains",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeAffinityPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeAffinityPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeTaintsPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeTaintsPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetTopologyKey() {
	_jsii_.InvokeVoid(
		j,
		"resetTopologyKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ResetWhenUnsatisfiable() {
	_jsii_.InvokeVoid(
		j,
		"resetWhenUnsatisfiable",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecTopologySpreadConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

