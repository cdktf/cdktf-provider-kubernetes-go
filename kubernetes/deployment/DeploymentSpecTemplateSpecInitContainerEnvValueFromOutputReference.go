// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference interface {
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
	ConfigMapKeyRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference
	FieldRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *DeploymentSpecTemplateSpecInitContainerEnvValueFrom
	SetInternalValue(val *DeploymentSpecTemplateSpecInitContainerEnvValueFrom)
	ResourceFieldRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef
	SecretKeyRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef)
	PutResourceFieldRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef)
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

// The jsii proxy struct for DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ConfigMapKeyRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ConfigMapKeyRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef {
	var returns *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) FieldRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) FieldRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef {
	var returns *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) InternalValue() *DeploymentSpecTemplateSpecInitContainerEnvValueFrom {
	var returns *DeploymentSpecTemplateSpecInitContainerEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResourceFieldRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResourceFieldRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef {
	var returns *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) SecretKeyRef() DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) SecretKeyRefInput() *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef {
	var returns *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference_Override(d DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference)SetInternalValue(val *DeploymentSpecTemplateSpecInitContainerEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) PutConfigMapKeyRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef) {
	if err := d.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMapKeyRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) PutFieldRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromFieldRef) {
	if err := d.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) PutResourceFieldRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef) {
	if err := d.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) PutSecretKeyRef(value *DeploymentSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef) {
	if err := d.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecretKeyRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

