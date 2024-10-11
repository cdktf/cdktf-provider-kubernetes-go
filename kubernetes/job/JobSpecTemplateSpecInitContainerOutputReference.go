// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecInitContainerOutputReference interface {
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
	Env() JobSpecTemplateSpecInitContainerEnvList
	EnvFrom() JobSpecTemplateSpecInitContainerEnvFromList
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
	Lifecycle() JobSpecTemplateSpecInitContainerLifecycleOutputReference
	LifecycleInput() *JobSpecTemplateSpecInitContainerLifecycle
	LivenessProbe() JobSpecTemplateSpecInitContainerLivenessProbeOutputReference
	LivenessProbeInput() *JobSpecTemplateSpecInitContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() JobSpecTemplateSpecInitContainerPortList
	PortInput() interface{}
	ReadinessProbe() JobSpecTemplateSpecInitContainerReadinessProbeOutputReference
	ReadinessProbeInput() *JobSpecTemplateSpecInitContainerReadinessProbe
	Resources() JobSpecTemplateSpecInitContainerResourcesOutputReference
	ResourcesInput() *JobSpecTemplateSpecInitContainerResources
	SecurityContext() JobSpecTemplateSpecInitContainerSecurityContextOutputReference
	SecurityContextInput() *JobSpecTemplateSpecInitContainerSecurityContext
	StartupProbe() JobSpecTemplateSpecInitContainerStartupProbeOutputReference
	StartupProbeInput() *JobSpecTemplateSpecInitContainerStartupProbe
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
	VolumeDevice() JobSpecTemplateSpecInitContainerVolumeDeviceList
	VolumeDeviceInput() interface{}
	VolumeMount() JobSpecTemplateSpecInitContainerVolumeMountList
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
	PutLifecycle(value *JobSpecTemplateSpecInitContainerLifecycle)
	PutLivenessProbe(value *JobSpecTemplateSpecInitContainerLivenessProbe)
	PutPort(value interface{})
	PutReadinessProbe(value *JobSpecTemplateSpecInitContainerReadinessProbe)
	PutResources(value *JobSpecTemplateSpecInitContainerResources)
	PutSecurityContext(value *JobSpecTemplateSpecInitContainerSecurityContext)
	PutStartupProbe(value *JobSpecTemplateSpecInitContainerStartupProbe)
	PutVolumeDevice(value interface{})
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
	ResetVolumeDevice()
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

// The jsii proxy struct for JobSpecTemplateSpecInitContainerOutputReference
type jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Env() JobSpecTemplateSpecInitContainerEnvList {
	var returns JobSpecTemplateSpecInitContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) EnvFrom() JobSpecTemplateSpecInitContainerEnvFromList {
	var returns JobSpecTemplateSpecInitContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) EnvFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Lifecycle() JobSpecTemplateSpecInitContainerLifecycleOutputReference {
	var returns JobSpecTemplateSpecInitContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) LifecycleInput() *JobSpecTemplateSpecInitContainerLifecycle {
	var returns *JobSpecTemplateSpecInitContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) LivenessProbe() JobSpecTemplateSpecInitContainerLivenessProbeOutputReference {
	var returns JobSpecTemplateSpecInitContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) LivenessProbeInput() *JobSpecTemplateSpecInitContainerLivenessProbe {
	var returns *JobSpecTemplateSpecInitContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Port() JobSpecTemplateSpecInitContainerPortList {
	var returns JobSpecTemplateSpecInitContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ReadinessProbe() JobSpecTemplateSpecInitContainerReadinessProbeOutputReference {
	var returns JobSpecTemplateSpecInitContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ReadinessProbeInput() *JobSpecTemplateSpecInitContainerReadinessProbe {
	var returns *JobSpecTemplateSpecInitContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Resources() JobSpecTemplateSpecInitContainerResourcesOutputReference {
	var returns JobSpecTemplateSpecInitContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResourcesInput() *JobSpecTemplateSpecInitContainerResources {
	var returns *JobSpecTemplateSpecInitContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) SecurityContext() JobSpecTemplateSpecInitContainerSecurityContextOutputReference {
	var returns JobSpecTemplateSpecInitContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) SecurityContextInput() *JobSpecTemplateSpecInitContainerSecurityContext {
	var returns *JobSpecTemplateSpecInitContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) StartupProbe() JobSpecTemplateSpecInitContainerStartupProbeOutputReference {
	var returns JobSpecTemplateSpecInitContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) StartupProbeInput() *JobSpecTemplateSpecInitContainerStartupProbe {
	var returns *JobSpecTemplateSpecInitContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Stdin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) StdinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) StdinOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) StdinOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Tty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) TtyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) VolumeDevice() JobSpecTemplateSpecInitContainerVolumeDeviceList {
	var returns JobSpecTemplateSpecInitContainerVolumeDeviceList
	_jsii_.Get(
		j,
		"volumeDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) VolumeDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) VolumeMount() JobSpecTemplateSpecInitContainerVolumeMountList {
	var returns JobSpecTemplateSpecInitContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) VolumeMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecInitContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecInitContainerOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecInitContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecInitContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecInitContainerOutputReference_Override(j JobSpecTemplateSpecInitContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecInitContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetStdin(val interface{}) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetStdinOnce(val interface{}) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetTty(val interface{}) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutEnv(value interface{}) {
	if err := j.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEnv",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutEnvFrom(value interface{}) {
	if err := j.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEnvFrom",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutLifecycle(value *JobSpecTemplateSpecInitContainerLifecycle) {
	if err := j.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutLivenessProbe(value *JobSpecTemplateSpecInitContainerLivenessProbe) {
	if err := j.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutPort(value interface{}) {
	if err := j.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPort",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutReadinessProbe(value *JobSpecTemplateSpecInitContainerReadinessProbe) {
	if err := j.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putReadinessProbe",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutResources(value *JobSpecTemplateSpecInitContainerResources) {
	if err := j.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putResources",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutSecurityContext(value *JobSpecTemplateSpecInitContainerSecurityContext) {
	if err := j.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutStartupProbe(value *JobSpecTemplateSpecInitContainerStartupProbe) {
	if err := j.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutVolumeDevice(value interface{}) {
	if err := j.validatePutVolumeDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolumeDevice",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) PutVolumeMount(value interface{}) {
	if err := j.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolumeMount",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		j,
		"resetArgs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		j,
		"resetCommand",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		j,
		"resetEnv",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		j,
		"resetImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		j,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		j,
		"resetResources",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		j,
		"resetStdin",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		j,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		j,
		"resetTty",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetVolumeDevice() {
	_jsii_.InvokeVoid(
		j,
		"resetVolumeDevice",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		j,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

