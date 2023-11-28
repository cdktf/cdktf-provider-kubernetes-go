// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/replicationcontroller/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference interface {
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
	InternalValue() *ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions
	SetInternalValue(val *ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions)
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

// The jsii proxy struct for ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference
type jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GmsaCredentialSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gmsaCredentialSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) HostProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InternalValue() *ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions {
	var returns *ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) RunAsUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference_Override(r ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpec(val *string) {
	if err := j.validateSetGmsaCredentialSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpec",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetGmsaCredentialSpecName(val *string) {
	if err := j.validateSetGmsaCredentialSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gmsaCredentialSpecName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetHostProcess(val interface{}) {
	if err := j.validateSetHostProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostProcess",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetInternalValue(val *ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetRunAsUsername(val *string) {
	if err := j.validateSetRunAsUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUsername",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpec() {
	_jsii_.InvokeVoid(
		r,
		"resetGmsaCredentialSpec",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetGmsaCredentialSpecName() {
	_jsii_.InvokeVoid(
		r,
		"resetGmsaCredentialSpecName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetHostProcess() {
	_jsii_.InvokeVoid(
		r,
		"resetHostProcess",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ResetRunAsUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetRunAsUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

