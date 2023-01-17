package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference interface {
	cdktf.ComplexObject
	AllowPrivilegeEscalation() interface{}
	SetAllowPrivilegeEscalation(val interface{})
	AllowPrivilegeEscalationInput() interface{}
	Capabilities() JobV1SpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference
	CapabilitiesInput() *JobV1SpecTemplateSpecInitContainerSecurityContextCapabilities
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
	InternalValue() *JobV1SpecTemplateSpecInitContainerSecurityContext
	SetInternalValue(val *JobV1SpecTemplateSpecInitContainerSecurityContext)
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
	SeccompProfile() JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile
	SeLinuxOptions() JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions
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
	PutCapabilities(value *JobV1SpecTemplateSpecInitContainerSecurityContextCapabilities)
	PutSeccompProfile(value *JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions)
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

// The jsii proxy struct for JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference
type jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) Capabilities() JobV1SpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerSecurityContextCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) CapabilitiesInput() *JobV1SpecTemplateSpecInitContainerSecurityContextCapabilities {
	var returns *JobV1SpecTemplateSpecInitContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) InternalValue() *JobV1SpecTemplateSpecInitContainerSecurityContext {
	var returns *JobV1SpecTemplateSpecInitContainerSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) SeccompProfile() JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) SeccompProfileInput() *JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile {
	var returns *JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) SeLinuxOptions() JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference {
	var returns JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) SeLinuxOptionsInput() *JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions {
	var returns *JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecInitContainerSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecInitContainerSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecInitContainerSecurityContextOutputReference_Override(j JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetAllowPrivilegeEscalation(val interface{}) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetInternalValue(val *JobV1SpecTemplateSpecInitContainerSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetReadOnlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) PutCapabilities(value *JobV1SpecTemplateSpecInitContainerSecurityContextCapabilities) {
	if err := j.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) PutSeccompProfile(value *JobV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile) {
	if err := j.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeccompProfile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) PutSeLinuxOptions(value *JobV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions) {
	if err := j.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSeLinuxOptions",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		j,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		j,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		j,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		j,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		j,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		j,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		j,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecInitContainerSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

