// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference interface {
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
	InternalValue() *JobSpecTemplateSpecSecurityContextWindowsOptions
	SetInternalValue(val *JobSpecTemplateSpecSecurityContextWindowsOptions)
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

// The jsii proxy struct for JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference
type jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InternalValue() *JobSpecTemplateSpecSecurityContextWindowsOptions {
	var returns *JobSpecTemplateSpecSecurityContextWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecSecurityContextWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference_Override(j JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpec(val *string) {
	if err := j.validateSetGmsaCredentialSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpec",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpecName(val *string) {
	if err := j.validateSetGmsaCredentialSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpecName",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetHostProcess(val interface{}) {
	if err := j.validateSetHostProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostProcess",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetInternalValue(val *JobSpecTemplateSpecSecurityContextWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetRunAsUsername(val *string) {
	if err := j.validateSetRunAsUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUsername",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpec() {
	_jsii_.InvokeVoid(
		j,
		"resetGmsaCredentialSpec",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpecName() {
	_jsii_.InvokeVoid(
		j,
		"resetGmsaCredentialSpecName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetHostProcess() {
	_jsii_.InvokeVoid(
		j,
		"resetHostProcess",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetRunAsUsername() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsUsername",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

