// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/statefulset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference interface {
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
	InternalValue() *StatefulSetSpecTemplateSpecSecurityContextWindowsOptions
	SetInternalValue(val *StatefulSetSpecTemplateSpecSecurityContextWindowsOptions)
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

// The jsii proxy struct for StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference
type jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InternalValue() *StatefulSetSpecTemplateSpecSecurityContextWindowsOptions {
	var returns *StatefulSetSpecTemplateSpecSecurityContextWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference_Override(s StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpec(val *string) {
	if err := j.validateSetGmsaCredentialSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpec",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpecName(val *string) {
	if err := j.validateSetGmsaCredentialSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpecName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetHostProcess(val interface{}) {
	if err := j.validateSetHostProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostProcess",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetInternalValue(val *StatefulSetSpecTemplateSpecSecurityContextWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetRunAsUsername(val *string) {
	if err := j.validateSetRunAsUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUsername",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetGmsaCredentialSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpecName() {
	_jsii_.InvokeVoid(
		s,
		"resetGmsaCredentialSpecName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetHostProcess() {
	_jsii_.InvokeVoid(
		s,
		"resetHostProcess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetRunAsUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetRunAsUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

