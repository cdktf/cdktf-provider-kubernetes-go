// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference interface {
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
	GmsaCredentialSpec() *string
	SetGmsaCredentialSpec(val *string)
	GmsaCredentialSpecInput() *string
	GmsaCredentialSpecName() *string
	SetGmsaCredentialSpecName(val *string)
	GmsaCredentialSpecNameInput() *string
	HostProcess() interface{}
	SetHostProcess(val interface{})
	HostProcessInput() interface{}
	InternalValue() *DeploymentV1SpecTemplateSpecSecurityContextWindowsOptions
	SetInternalValue(val *DeploymentV1SpecTemplateSpecSecurityContextWindowsOptions)
	RunAsUsername() *string
	SetRunAsUsername(val *string)
	RunAsUsernameInput() *string
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
	ResetGmsaCredentialSpec()
	ResetGmsaCredentialSpecName()
	ResetHostProcess()
	ResetRunAsUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference
type jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InternalValue() *DeploymentV1SpecTemplateSpecSecurityContextWindowsOptions {
	var returns *DeploymentV1SpecTemplateSpecSecurityContextWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference_Override(d DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpec(val *string) {
	if err := j.validateSetGmsaCredentialSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpec",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpecName(val *string) {
	if err := j.validateSetGmsaCredentialSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpecName",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetHostProcess(val interface{}) {
	if err := j.validateSetHostProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostProcess",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetInternalValue(val *DeploymentV1SpecTemplateSpecSecurityContextWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetRunAsUsername(val *string) {
	if err := j.validateSetRunAsUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUsername",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetGmsaCredentialSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpecName() {
	_jsii_.InvokeVoid(
		d,
		"resetGmsaCredentialSpecName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetHostProcess() {
	_jsii_.InvokeVoid(
		d,
		"resetHostProcess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetRunAsUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

