// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference interface {
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
	Exec() DeploymentSpecTemplateSpecContainerLivenessProbeExecOutputReference
	ExecInput() *DeploymentSpecTemplateSpecContainerLivenessProbeExec
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Grpc() DeploymentSpecTemplateSpecContainerLivenessProbeGrpcList
	GrpcInput() interface{}
	HttpGet() DeploymentSpecTemplateSpecContainerLivenessProbeHttpGetOutputReference
	HttpGetInput() *DeploymentSpecTemplateSpecContainerLivenessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *DeploymentSpecTemplateSpecContainerLivenessProbe
	SetInternalValue(val *DeploymentSpecTemplateSpecContainerLivenessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	SuccessThreshold() *float64
	SetSuccessThreshold(val *float64)
	SuccessThresholdInput() *float64
	TcpSocket() DeploymentSpecTemplateSpecContainerLivenessProbeTcpSocketList
	TcpSocketInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutExec(value *DeploymentSpecTemplateSpecContainerLivenessProbeExec)
	PutGrpc(value interface{})
	PutHttpGet(value *DeploymentSpecTemplateSpecContainerLivenessProbeHttpGet)
	PutTcpSocket(value interface{})
	ResetExec()
	ResetFailureThreshold()
	ResetGrpc()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetSuccessThreshold()
	ResetTcpSocket()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) Exec() DeploymentSpecTemplateSpecContainerLivenessProbeExecOutputReference {
	var returns DeploymentSpecTemplateSpecContainerLivenessProbeExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ExecInput() *DeploymentSpecTemplateSpecContainerLivenessProbeExec {
	var returns *DeploymentSpecTemplateSpecContainerLivenessProbeExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) Grpc() DeploymentSpecTemplateSpecContainerLivenessProbeGrpcList {
	var returns DeploymentSpecTemplateSpecContainerLivenessProbeGrpcList
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GrpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) HttpGet() DeploymentSpecTemplateSpecContainerLivenessProbeHttpGetOutputReference {
	var returns DeploymentSpecTemplateSpecContainerLivenessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) HttpGetInput() *DeploymentSpecTemplateSpecContainerLivenessProbeHttpGet {
	var returns *DeploymentSpecTemplateSpecContainerLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) InternalValue() *DeploymentSpecTemplateSpecContainerLivenessProbe {
	var returns *DeploymentSpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) SuccessThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) SuccessThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TcpSocket() DeploymentSpecTemplateSpecContainerLivenessProbeTcpSocketList {
	var returns DeploymentSpecTemplateSpecContainerLivenessProbeTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TcpSocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewDeploymentSpecTemplateSpecContainerLivenessProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecContainerLivenessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecContainerLivenessProbeOutputReference_Override(d DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetInternalValue(val *DeploymentSpecTemplateSpecContainerLivenessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetSuccessThreshold(val *float64) {
	if err := j.validateSetSuccessThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successThreshold",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PutExec(value *DeploymentSpecTemplateSpecContainerLivenessProbeExec) {
	if err := d.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PutGrpc(value interface{}) {
	if err := d.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGrpc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PutHttpGet(value *DeploymentSpecTemplateSpecContainerLivenessProbeHttpGet) {
	if err := d.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) PutTcpSocket(value interface{}) {
	if err := d.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		d,
		"resetExec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		d,
		"resetGrpc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetSuccessThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetSuccessThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		d,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLivenessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

