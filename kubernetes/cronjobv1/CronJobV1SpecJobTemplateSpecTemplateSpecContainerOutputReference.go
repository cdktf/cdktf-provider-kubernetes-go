// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference interface {
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
	Env() CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvList
	EnvFrom() CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvFromList
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
	Lifecycle() CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference
	LifecycleInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycle
	LivenessProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() CronJobV1SpecJobTemplateSpecTemplateSpecContainerPortList
	PortInput() interface{}
	ReadinessProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbe
	Resources() CronJobV1SpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference
	ResourcesInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources
	SecurityContext() CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference
	SecurityContextInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContext
	StartupProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference
	StartupProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbe
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
	VolumeDevice() CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList
	VolumeDeviceInput() interface{}
	VolumeMount() CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeMountList
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
	PutLifecycle(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycle)
	PutLivenessProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbe)
	PutPort(value interface{})
	PutReadinessProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbe)
	PutResources(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources)
	PutSecurityContext(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContext)
	PutStartupProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbe)
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

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Env() CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvFrom() CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvFromList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Lifecycle() CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) LifecycleInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycle {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) LivenessProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) LivenessProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbe {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Port() CronJobV1SpecJobTemplateSpecTemplateSpecContainerPortList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ReadinessProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ReadinessProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbe {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Resources() CronJobV1SpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResourcesInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) SecurityContext() CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) SecurityContextInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContext {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) StartupProbe() CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) StartupProbeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbe {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Stdin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Tty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) TtyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeDevice() CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList
	_jsii_.Get(
		j,
		"volumeDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeMount() CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeMountList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetStdin(val interface{}) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetStdinOnce(val interface{}) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTty(val interface{}) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutEnv(value interface{}) {
	if err := c.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnv",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutEnvFrom(value interface{}) {
	if err := c.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvFrom",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutLifecycle(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLifecycle) {
	if err := c.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutLivenessProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerLivenessProbe) {
	if err := c.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutPort(value interface{}) {
	if err := c.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutReadinessProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerReadinessProbe) {
	if err := c.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadinessProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutResources(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources) {
	if err := c.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutSecurityContext(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContext) {
	if err := c.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutStartupProbe(value *CronJobV1SpecJobTemplateSpecTemplateSpecContainerStartupProbe) {
	if err := c.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutVolumeDevice(value interface{}) {
	if err := c.validatePutVolumeDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeDevice",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) PutVolumeMount(value interface{}) {
	if err := c.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeMount",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		c,
		"resetArgs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		c,
		"resetCommand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		c,
		"resetImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		c,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		c,
		"resetStdin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		c,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		c,
		"resetTty",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetVolumeDevice() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeDevice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

