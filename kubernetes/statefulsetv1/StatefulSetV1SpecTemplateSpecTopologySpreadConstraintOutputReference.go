// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/statefulsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference interface {
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
	LabelSelector() StatefulSetV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList
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

// The jsii proxy struct for StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference
type jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelector() StatefulSetV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList {
	var returns StatefulSetV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomains() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomainsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiableInput",
		&returns,
	)
	return returns
}


func NewStatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference_Override(s StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMatchLabelKeys(val *[]*string) {
	if err := j.validateSetMatchLabelKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelKeys",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMaxSkew(val *float64) {
	if err := j.validateSetMaxSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSkew",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMinDomains(val *float64) {
	if err := j.validateSetMinDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDomains",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeAffinityPolicy(val *string) {
	if err := j.validateSetNodeAffinityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeAffinityPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeTaintsPolicy(val *string) {
	if err := j.validateSetNodeTaintsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTaintsPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTopologyKey(val *string) {
	if err := j.validateSetTopologyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyKey",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetWhenUnsatisfiable(val *string) {
	if err := j.validateSetWhenUnsatisfiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenUnsatisfiable",
		val,
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) PutLabelSelector(value interface{}) {
	if err := s.validatePutLabelSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLabelSelector",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetLabelSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMatchLabelKeys() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchLabelKeys",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMaxSkew() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSkew",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMinDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetMinDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeAffinityPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeAffinityPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeTaintsPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeTaintsPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetTopologyKey() {
	_jsii_.InvokeVoid(
		s,
		"resetTopologyKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetWhenUnsatisfiable() {
	_jsii_.InvokeVoid(
		s,
		"resetWhenUnsatisfiable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

