// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference interface {
	cdktf.ComplexObject
	AllowPrivilegeEscalation() interface{}
	SetAllowPrivilegeEscalation(val interface{})
	AllowPrivilegeEscalationInput() interface{}
	Capabilities() DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference
	CapabilitiesInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilities
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
	InternalValue() *DeploymentSpecTemplateSpecInitContainerSecurityContext
	SetInternalValue(val *DeploymentSpecTemplateSpecInitContainerSecurityContext)
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	ReadOnlyRootFilesystem() interface{}
	SetReadOnlyRootFilesystem(val interface{})
	ReadOnlyRootFilesystemInput() interface{}
	RunAsGroup() *string
	SetRunAsGroup(val *string)
	RunAsGroupInput() *string
	RunAsNonRoot() interface{}
	SetRunAsNonRoot(val interface{})
	RunAsNonRootInput() interface{}
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	SeccompProfile() DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfile
	SeLinuxOptions() DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions
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
	PutCapabilities(value *DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilities)
	PutSeccompProfile(value *DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions)
	ResetAllowPrivilegeEscalation()
	ResetCapabilities()
	ResetPrivileged()
	ResetReadOnlyRootFilesystem()
	ResetRunAsGroup()
	ResetRunAsNonRoot()
	ResetRunAsUser()
	ResetSeccompProfile()
	ResetSeLinuxOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) Capabilities() DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) CapabilitiesInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilities {
	var returns *DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) InternalValue() *DeploymentSpecTemplateSpecInitContainerSecurityContext {
	var returns *DeploymentSpecTemplateSpecInitContainerSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) SeccompProfile() DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) SeccompProfileInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfile {
	var returns *DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) SeLinuxOptions() DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference {
	var returns DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) SeLinuxOptionsInput() *DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions {
	var returns *DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecInitContainerSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference_Override(d DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetAllowPrivilegeEscalation(val interface{}) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetInternalValue(val *DeploymentSpecTemplateSpecInitContainerSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetReadOnlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) PutCapabilities(value *DeploymentSpecTemplateSpecInitContainerSecurityContextCapabilities) {
	if err := d.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) PutSeccompProfile(value *DeploymentSpecTemplateSpecInitContainerSecurityContextSeccompProfile) {
	if err := d.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSeccompProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) PutSeLinuxOptions(value *DeploymentSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions) {
	if err := d.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSeLinuxOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		d,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

