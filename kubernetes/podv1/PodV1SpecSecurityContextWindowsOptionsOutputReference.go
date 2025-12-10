// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/podv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodV1SpecSecurityContextWindowsOptionsOutputReference interface {
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
	InternalValue() *PodV1SpecSecurityContextWindowsOptions
	SetInternalValue(val *PodV1SpecSecurityContextWindowsOptions)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetGmsaCredentialSpec()
	ResetGmsaCredentialSpecName()
	ResetHostProcess()
	ResetRunAsUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodV1SpecSecurityContextWindowsOptionsOutputReference
type jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) HostProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) HostProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) InternalValue() *PodV1SpecSecurityContextWindowsOptions {
	var returns *PodV1SpecSecurityContextWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) RunAsUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) RunAsUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPodV1SpecSecurityContextWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodV1SpecSecurityContextWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecSecurityContextWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodV1SpecSecurityContextWindowsOptionsOutputReference_Override(p PodV1SpecSecurityContextWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpec(val *string) {
	if err := j.validateSetGmsaCredentialSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpec",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpecName(val *string) {
	if err := j.validateSetGmsaCredentialSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpecName",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetHostProcess(val interface{}) {
	if err := j.validateSetHostProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostProcess",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetInternalValue(val *PodV1SpecSecurityContextWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetRunAsUsername(val *string) {
	if err := j.validateSetRunAsUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUsername",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpec() {
	_jsii_.InvokeVoid(
		p,
		"resetGmsaCredentialSpec",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpecName() {
	_jsii_.InvokeVoid(
		p,
		"resetGmsaCredentialSpecName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ResetHostProcess() {
	_jsii_.InvokeVoid(
		p,
		"resetHostProcess",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ResetRunAsUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecSecurityContextWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

