// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/cronjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference interface {
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
	Env() CronJobSpecJobTemplateSpecTemplateSpecContainerEnvList
	EnvFrom() CronJobSpecJobTemplateSpecTemplateSpecContainerEnvFromList
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
	Lifecycle() CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference
	LifecycleInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycle
	LivenessProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() CronJobSpecJobTemplateSpecTemplateSpecContainerPortList
	PortInput() interface{}
	ReadinessProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbe
	Resources() CronJobSpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference
	ResourcesInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerResources
	SecurityContext() CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference
	SecurityContextInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContext
	StartupProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference
	StartupProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbe
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
	VolumeDevice() CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList
	VolumeDeviceInput() interface{}
	VolumeMount() CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeMountList
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
	PutLifecycle(value *CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycle)
	PutLivenessProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbe)
	PutPort(value interface{})
	PutReadinessProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbe)
	PutResources(value *CronJobSpecJobTemplateSpecTemplateSpecContainerResources)
	PutSecurityContext(value *CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContext)
	PutStartupProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbe)
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

// The jsii proxy struct for CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference
type jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Env() CronJobSpecJobTemplateSpecTemplateSpecContainerEnvList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvFrom() CronJobSpecJobTemplateSpecTemplateSpecContainerEnvFromList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Lifecycle() CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) LifecycleInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycle {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) LivenessProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) LivenessProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbe {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Port() CronJobSpecJobTemplateSpecTemplateSpecContainerPortList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ReadinessProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ReadinessProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbe {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Resources() CronJobSpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResourcesInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerResources {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) SecurityContext() CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) SecurityContextInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContext {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) StartupProbe() CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) StartupProbeInput() *CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbe {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Stdin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) StdinOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Tty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) TtyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeDevice() CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeDeviceList
	_jsii_.Get(
		j,
		"volumeDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeMount() CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeMountList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) VolumeMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewCronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecJobTemplateSpecTemplateSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference_Override(c CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetStdin(val interface{}) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetStdinOnce(val interface{}) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetTty(val interface{}) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutEnv(value interface{}) {
	if err := c.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnv",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutEnvFrom(value interface{}) {
	if err := c.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvFrom",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutLifecycle(value *CronJobSpecJobTemplateSpecTemplateSpecContainerLifecycle) {
	if err := c.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutLivenessProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerLivenessProbe) {
	if err := c.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutPort(value interface{}) {
	if err := c.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutReadinessProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerReadinessProbe) {
	if err := c.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadinessProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutResources(value *CronJobSpecJobTemplateSpecTemplateSpecContainerResources) {
	if err := c.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResources",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutSecurityContext(value *CronJobSpecJobTemplateSpecTemplateSpecContainerSecurityContext) {
	if err := c.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutStartupProbe(value *CronJobSpecJobTemplateSpecTemplateSpecContainerStartupProbe) {
	if err := c.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutVolumeDevice(value interface{}) {
	if err := c.validatePutVolumeDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeDevice",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) PutVolumeMount(value interface{}) {
	if err := c.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolumeMount",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		c,
		"resetArgs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		c,
		"resetCommand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		c,
		"resetEnv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		c,
		"resetImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		c,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		c,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		c,
		"resetStdin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		c,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		c,
		"resetTty",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetVolumeDevice() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeDevice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		c,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

