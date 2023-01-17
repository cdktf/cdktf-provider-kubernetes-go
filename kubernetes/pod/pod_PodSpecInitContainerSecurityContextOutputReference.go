package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/pod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodSpecInitContainerSecurityContextOutputReference interface {
	cdktf.ComplexObject
	AllowPrivilegeEscalation() interface{}
	SetAllowPrivilegeEscalation(val interface{})
	AllowPrivilegeEscalationInput() interface{}
	Capabilities() PodSpecInitContainerSecurityContextCapabilitiesOutputReference
	CapabilitiesInput() *PodSpecInitContainerSecurityContextCapabilities
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
	InternalValue() *PodSpecInitContainerSecurityContext
	SetInternalValue(val *PodSpecInitContainerSecurityContext)
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
	SeccompProfile() PodSpecInitContainerSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *PodSpecInitContainerSecurityContextSeccompProfile
	SeLinuxOptions() PodSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *PodSpecInitContainerSecurityContextSeLinuxOptions
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
	PutCapabilities(value *PodSpecInitContainerSecurityContextCapabilities)
	PutSeccompProfile(value *PodSpecInitContainerSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *PodSpecInitContainerSecurityContextSeLinuxOptions)
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

// The jsii proxy struct for PodSpecInitContainerSecurityContextOutputReference
type jsiiProxy_PodSpecInitContainerSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) AllowPrivilegeEscalationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) Capabilities() PodSpecInitContainerSecurityContextCapabilitiesOutputReference {
	var returns PodSpecInitContainerSecurityContextCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) CapabilitiesInput() *PodSpecInitContainerSecurityContextCapabilities {
	var returns *PodSpecInitContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) InternalValue() *PodSpecInitContainerSecurityContext {
	var returns *PodSpecInitContainerSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ReadOnlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) SeccompProfile() PodSpecInitContainerSecurityContextSeccompProfileOutputReference {
	var returns PodSpecInitContainerSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) SeccompProfileInput() *PodSpecInitContainerSecurityContextSeccompProfile {
	var returns *PodSpecInitContainerSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) SeLinuxOptions() PodSpecInitContainerSecurityContextSeLinuxOptionsOutputReference {
	var returns PodSpecInitContainerSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) SeLinuxOptionsInput() *PodSpecInitContainerSecurityContextSeLinuxOptions {
	var returns *PodSpecInitContainerSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPodSpecInitContainerSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSpecInitContainerSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecInitContainerSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecInitContainerSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSpecInitContainerSecurityContextOutputReference_Override(p PodSpecInitContainerSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecInitContainerSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetAllowPrivilegeEscalation(val interface{}) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetInternalValue(val *PodSpecInitContainerSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetReadOnlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) PutCapabilities(value *PodSpecInitContainerSecurityContextCapabilities) {
	if err := p.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) PutSeccompProfile(value *PodSpecInitContainerSecurityContextSeccompProfile) {
	if err := p.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeccompProfile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) PutSeLinuxOptions(value *PodSpecInitContainerSecurityContextSeLinuxOptions) {
	if err := p.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeLinuxOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		p,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecInitContainerSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

