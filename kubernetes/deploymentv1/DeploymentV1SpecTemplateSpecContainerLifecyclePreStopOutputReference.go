// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference interface {
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
	Exec() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExecOutputReference
	ExecInput() *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExec
	// Experimental.
	Fqn() *string
	HttpGet() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGetOutputReference
	HttpGetInput() *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGet
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TcpSocket() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList
	TcpSocketInput() interface{}
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
	PutExec(value *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExec)
	PutHttpGet(value *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGet)
	PutTcpSocket(value interface{})
	ResetExec()
	ResetHttpGet()
	ResetTcpSocket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference
type jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) Exec() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExecOutputReference {
	var returns DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ExecInput() *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExec {
	var returns *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) HttpGet() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGetOutputReference {
	var returns DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) HttpGetInput() *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGet {
	var returns *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) TcpSocket() DeploymentV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList {
	var returns DeploymentV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) TcpSocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference_Override(d DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) PutExec(value *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopExec) {
	if err := d.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) PutHttpGet(value *DeploymentV1SpecTemplateSpecContainerLifecyclePreStopHttpGet) {
	if err := d.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) PutTcpSocket(value interface{}) {
	if err := d.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		d,
		"resetExec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		d,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerLifecyclePreStopOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

