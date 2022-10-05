package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecContainerOutputReference interface {
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
	Env() ReplicationControllerV1SpecTemplateSpecContainerEnvList
	EnvFrom() ReplicationControllerV1SpecTemplateSpecContainerEnvFromList
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
	Lifecycle() ReplicationControllerV1SpecTemplateSpecContainerLifecycleOutputReference
	LifecycleInput() *ReplicationControllerV1SpecTemplateSpecContainerLifecycle
	LivenessProbe() ReplicationControllerV1SpecTemplateSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() ReplicationControllerV1SpecTemplateSpecContainerPortList
	PortInput() interface{}
	ReadinessProbe() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe
	Resources() ReplicationControllerV1SpecTemplateSpecContainerResourcesOutputReference
	ResourcesInput() *ReplicationControllerV1SpecTemplateSpecContainerResources
	SecurityContext() ReplicationControllerV1SpecTemplateSpecContainerSecurityContextOutputReference
	SecurityContextInput() *ReplicationControllerV1SpecTemplateSpecContainerSecurityContext
	StartupProbe() ReplicationControllerV1SpecTemplateSpecContainerStartupProbeOutputReference
	StartupProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerStartupProbe
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
	VolumeMount() ReplicationControllerV1SpecTemplateSpecContainerVolumeMountList
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
	PutLifecycle(value *ReplicationControllerV1SpecTemplateSpecContainerLifecycle)
	PutLivenessProbe(value *ReplicationControllerV1SpecTemplateSpecContainerLivenessProbe)
	PutPort(value interface{})
	PutReadinessProbe(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe)
	PutResources(value *ReplicationControllerV1SpecTemplateSpecContainerResources)
	PutSecurityContext(value *ReplicationControllerV1SpecTemplateSpecContainerSecurityContext)
	PutStartupProbe(value *ReplicationControllerV1SpecTemplateSpecContainerStartupProbe)
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

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecContainerOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Env() ReplicationControllerV1SpecTemplateSpecContainerEnvList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) EnvFrom() ReplicationControllerV1SpecTemplateSpecContainerEnvFromList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) EnvFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Lifecycle() ReplicationControllerV1SpecTemplateSpecContainerLifecycleOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) LifecycleInput() *ReplicationControllerV1SpecTemplateSpecContainerLifecycle {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) LivenessProbe() ReplicationControllerV1SpecTemplateSpecContainerLivenessProbeOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) LivenessProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerLivenessProbe {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Port() ReplicationControllerV1SpecTemplateSpecContainerPortList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ReadinessProbe() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ReadinessProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Resources() ReplicationControllerV1SpecTemplateSpecContainerResourcesOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResourcesInput() *ReplicationControllerV1SpecTemplateSpecContainerResources {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) SecurityContext() ReplicationControllerV1SpecTemplateSpecContainerSecurityContextOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) SecurityContextInput() *ReplicationControllerV1SpecTemplateSpecContainerSecurityContext {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) StartupProbe() ReplicationControllerV1SpecTemplateSpecContainerStartupProbeOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) StartupProbeInput() *ReplicationControllerV1SpecTemplateSpecContainerStartupProbe {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Stdin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) StdinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) StdinOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) StdinOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Tty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) TtyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) VolumeMount() ReplicationControllerV1SpecTemplateSpecContainerVolumeMountList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) VolumeMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReplicationControllerV1SpecTemplateSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecContainerOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetStdin(val interface{}) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetStdinOnce(val interface{}) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetTty(val interface{}) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference)SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutEnv(value interface{}) {
	if err := r.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEnv",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutEnvFrom(value interface{}) {
	if err := r.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEnvFrom",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutLifecycle(value *ReplicationControllerV1SpecTemplateSpecContainerLifecycle) {
	if err := r.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutLivenessProbe(value *ReplicationControllerV1SpecTemplateSpecContainerLivenessProbe) {
	if err := r.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutPort(value interface{}) {
	if err := r.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPort",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutReadinessProbe(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe) {
	if err := r.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putReadinessProbe",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutResources(value *ReplicationControllerV1SpecTemplateSpecContainerResources) {
	if err := r.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResources",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutSecurityContext(value *ReplicationControllerV1SpecTemplateSpecContainerSecurityContext) {
	if err := r.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutStartupProbe(value *ReplicationControllerV1SpecTemplateSpecContainerStartupProbe) {
	if err := r.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) PutVolumeMount(value interface{}) {
	if err := r.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVolumeMount",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		r,
		"resetArgs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		r,
		"resetCommand",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		r,
		"resetEnv",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		r,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		r,
		"resetImage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		r,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		r,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		r,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		r,
		"resetResources",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		r,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		r,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		r,
		"resetStdin",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		r,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		r,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		r,
		"resetTty",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		r,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		r,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

