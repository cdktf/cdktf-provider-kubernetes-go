// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference interface {
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
	LabelSelector() DeploymentV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList
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

// The jsii proxy struct for DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference
type jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelector() DeploymentV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList {
	var returns DeploymentV1SpecTemplateSpecTopologySpreadConstraintLabelSelectorList
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) LabelSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MatchLabelKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchLabelKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MaxSkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSkewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomains() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) MinDomainsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeAffinityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAffinityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) NodeTaintsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTaintsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) TopologyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) WhenUnsatisfiableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whenUnsatisfiableInput",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference_Override(d DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMatchLabelKeys(val *[]*string) {
	if err := j.validateSetMatchLabelKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabelKeys",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMaxSkew(val *float64) {
	if err := j.validateSetMaxSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSkew",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetMinDomains(val *float64) {
	if err := j.validateSetMinDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDomains",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeAffinityPolicy(val *string) {
	if err := j.validateSetNodeAffinityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeAffinityPolicy",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetNodeTaintsPolicy(val *string) {
	if err := j.validateSetNodeTaintsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTaintsPolicy",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetTopologyKey(val *string) {
	if err := j.validateSetTopologyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyKey",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference)SetWhenUnsatisfiable(val *string) {
	if err := j.validateSetWhenUnsatisfiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whenUnsatisfiable",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) PutLabelSelector(value interface{}) {
	if err := d.validatePutLabelSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLabelSelector",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetLabelSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetLabelSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMatchLabelKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchLabelKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMaxSkew() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxSkew",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetMinDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetMinDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeAffinityPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeAffinityPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetNodeTaintsPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeTaintsPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetTopologyKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologyKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ResetWhenUnsatisfiable() {
	_jsii_.InvokeVoid(
		d,
		"resetWhenUnsatisfiable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecTopologySpreadConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

