// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference interface {
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
	ConfigMapKeyRef() JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference
	FieldRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *JobV1SpecTemplateSpecInitContainerEnvValueFrom
	SetInternalValue(val *JobV1SpecTemplateSpecInitContainerEnvValueFrom)
	ResourceFieldRef() JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef
	SecretKeyRef() JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRef)
	PutResourceFieldRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef)
	ResetConfigMapKeyRef()
	ResetFieldRef()
	ResetResourceFieldRef()
	ResetSecretKeyRef()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference
type jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ConfigMapKeyRef() JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ConfigMapKeyRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef {
	var returns *JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) FieldRef() JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) FieldRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRef {
	var returns *JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) InternalValue() *JobV1SpecTemplateSpecInitContainerEnvValueFrom {
	var returns *JobV1SpecTemplateSpecInitContainerEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResourceFieldRef() JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResourceFieldRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef {
	var returns *JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) SecretKeyRef() JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) SecretKeyRefInput() *JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef {
	var returns *JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecInitContainerEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference_Override(j JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference)SetInternalValue(val *JobV1SpecTemplateSpecInitContainerEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) PutConfigMapKeyRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef) {
	if err := j.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfigMapKeyRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) PutFieldRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromFieldRef) {
	if err := j.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) PutResourceFieldRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef) {
	if err := j.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) PutSecretKeyRef(value *JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef) {
	if err := j.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecretKeyRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		j,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		j,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		j,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		j,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

