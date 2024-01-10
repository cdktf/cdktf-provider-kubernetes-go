// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/podv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodV1SpecContainerOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() PodV1SpecContainerEnvList
	EnvFrom() PodV1SpecContainerEnvFromList
	EnvFromInput() interface{}
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullPolicy() *string
	SetImagePullPolicy(val *string)
	ImagePullPolicyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lifecycle() PodV1SpecContainerLifecycleOutputReference
	LifecycleInput() *PodV1SpecContainerLifecycle
	LivenessProbe() PodV1SpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *PodV1SpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() PodV1SpecContainerPortList
	PortInput() interface{}
	ReadinessProbe() PodV1SpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *PodV1SpecContainerReadinessProbe
	Resources() PodV1SpecContainerResourcesOutputReference
	ResourcesInput() *PodV1SpecContainerResources
	SecurityContext() PodV1SpecContainerSecurityContextOutputReference
	SecurityContextInput() *PodV1SpecContainerSecurityContext
	StartupProbe() PodV1SpecContainerStartupProbeOutputReference
	StartupProbeInput() *PodV1SpecContainerStartupProbe
	Stdin() interface{}
	SetStdin(val interface{})
	StdinInput() interface{}
	StdinOnce() interface{}
	SetStdinOnce(val interface{})
	StdinOnceInput() interface{}
	TerminationMessagePath() *string
	SetTerminationMessagePath(val *string)
	TerminationMessagePathInput() *string
	TerminationMessagePolicy() *string
	SetTerminationMessagePolicy(val *string)
	TerminationMessagePolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tty() interface{}
	SetTty(val interface{})
	TtyInput() interface{}
	VolumeMount() PodV1SpecContainerVolumeMountList
	VolumeMountInput() interface{}
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	PutEnv(value interface{})
	PutEnvFrom(value interface{})
	PutLifecycle(value *PodV1SpecContainerLifecycle)
	PutLivenessProbe(value *PodV1SpecContainerLivenessProbe)
	PutPort(value interface{})
	PutReadinessProbe(value *PodV1SpecContainerReadinessProbe)
	PutResources(value *PodV1SpecContainerResources)
	PutSecurityContext(value *PodV1SpecContainerSecurityContext)
	PutStartupProbe(value *PodV1SpecContainerStartupProbe)
	PutVolumeMount(value interface{})
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetEnvFrom()
	ResetImage()
	ResetImagePullPolicy()
	ResetLifecycle()
	ResetLivenessProbe()
	ResetPort()
	ResetReadinessProbe()
	ResetResources()
	ResetSecurityContext()
	ResetStartupProbe()
	ResetStdin()
	ResetStdinOnce()
	ResetTerminationMessagePath()
	ResetTerminationMessagePolicy()
	ResetTty()
	ResetVolumeMount()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodV1SpecContainerOutputReference
type jsiiProxy_PodV1SpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Env() PodV1SpecContainerEnvList {
	var returns PodV1SpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) EnvFrom() PodV1SpecContainerEnvFromList {
	var returns PodV1SpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) EnvFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Lifecycle() PodV1SpecContainerLifecycleOutputReference {
	var returns PodV1SpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) LifecycleInput() *PodV1SpecContainerLifecycle {
	var returns *PodV1SpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) LivenessProbe() PodV1SpecContainerLivenessProbeOutputReference {
	var returns PodV1SpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) LivenessProbeInput() *PodV1SpecContainerLivenessProbe {
	var returns *PodV1SpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Port() PodV1SpecContainerPortList {
	var returns PodV1SpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ReadinessProbe() PodV1SpecContainerReadinessProbeOutputReference {
	var returns PodV1SpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ReadinessProbeInput() *PodV1SpecContainerReadinessProbe {
	var returns *PodV1SpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Resources() PodV1SpecContainerResourcesOutputReference {
	var returns PodV1SpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) ResourcesInput() *PodV1SpecContainerResources {
	var returns *PodV1SpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) SecurityContext() PodV1SpecContainerSecurityContextOutputReference {
	var returns PodV1SpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) SecurityContextInput() *PodV1SpecContainerSecurityContext {
	var returns *PodV1SpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) StartupProbe() PodV1SpecContainerStartupProbeOutputReference {
	var returns PodV1SpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) StartupProbeInput() *PodV1SpecContainerStartupProbe {
	var returns *PodV1SpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Stdin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) StdinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) StdinOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) StdinOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) Tty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) TtyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) VolumeMount() PodV1SpecContainerVolumeMountList {
	var returns PodV1SpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) VolumeMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewPodV1SpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PodV1SpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPodV1SpecContainerOutputReference_Override(p PodV1SpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetStdin(val interface{}) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetStdinOnce(val interface{}) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetTty(val interface{}) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutEnv(value interface{}) {
	if err := p.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnv",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutEnvFrom(value interface{}) {
	if err := p.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvFrom",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutLifecycle(value *PodV1SpecContainerLifecycle) {
	if err := p.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutLivenessProbe(value *PodV1SpecContainerLivenessProbe) {
	if err := p.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutPort(value interface{}) {
	if err := p.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutReadinessProbe(value *PodV1SpecContainerReadinessProbe) {
	if err := p.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReadinessProbe",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutResources(value *PodV1SpecContainerResources) {
	if err := p.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putResources",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutSecurityContext(value *PodV1SpecContainerSecurityContext) {
	if err := p.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutStartupProbe(value *PodV1SpecContainerStartupProbe) {
	if err := p.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) PutVolumeMount(value interface{}) {
	if err := p.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVolumeMount",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		p,
		"resetArgs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		p,
		"resetCommand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		p,
		"resetEnv",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		p,
		"resetImage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		p,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		p,
		"resetPort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		p,
		"resetResources",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		p,
		"resetStdin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		p,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		p,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		p,
		"resetTty",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		p,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PodV1SpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

