// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumeclaimv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/datakubernetespersistentvolumeclaimv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference interface {
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
	MatchExpressions() DataKubernetesPersistentVolumeClaimV1SpecSelectorMatchExpressionsList
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

// The jsii proxy struct for DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference
type jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) MatchExpressions() DataKubernetesPersistentVolumeClaimV1SpecSelectorMatchExpressionsList {
	var returns DataKubernetesPersistentVolumeClaimV1SpecSelectorMatchExpressionsList
	_jsii_.Get(
		j,
		"matchExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) MatchExpressionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchExpressionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) MatchLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) MatchLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeClaimV1.DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference_Override(d DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeClaimV1.DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetMatchLabels(val *map[string]*string) {
	if err := j.validateSetMatchLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchLabels",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) PutMatchExpressions(value interface{}) {
	if err := d.validatePutMatchExpressionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMatchExpressions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ResetMatchExpressions() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchExpressions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ResetMatchLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeClaimV1SpecSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

