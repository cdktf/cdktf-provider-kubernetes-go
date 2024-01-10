// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference interface {
	cdktf.ComplexObject
	ApiVersion() *string
	SetApiVersion(val *string)
	ApiVersionInput() *string
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
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef
	SetInternalValue(val *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef)
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
	ResetApiVersion()
	ResetFieldPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference
type jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) InternalValue() *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef {
	var returns *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference_Override(j JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetApiVersion(val *string) {
	if err := j.validateSetApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetInternalValue(val *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ResetApiVersion() {
	_jsii_.InvokeVoid(
		j,
		"resetApiVersion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		j,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

