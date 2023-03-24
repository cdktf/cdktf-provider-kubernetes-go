package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() JobV1SpecTemplateSpecAffinityOutputReference
	AffinityInput() *JobV1SpecTemplateSpecAffinity
	AutomountServiceAccountToken() interface{}
	SetAutomountServiceAccountToken(val interface{})
	AutomountServiceAccountTokenInput() interface{}
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
	Container() JobV1SpecTemplateSpecContainerList
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() JobV1SpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *JobV1SpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() interface{}
	SetEnableServiceLinks(val interface{})
	EnableServiceLinksInput() interface{}
	// Experimental.
	Fqn() *string
	HostAliases() JobV1SpecTemplateSpecHostAliasesList
	HostAliasesInput() interface{}
	HostIpc() interface{}
	SetHostIpc(val interface{})
	HostIpcInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	HostNetwork() interface{}
	SetHostNetwork(val interface{})
	HostNetworkInput() interface{}
	HostPid() interface{}
	SetHostPid(val interface{})
	HostPidInput() interface{}
	ImagePullSecrets() JobV1SpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() interface{}
	InitContainer() JobV1SpecTemplateSpecInitContainerList
	InitContainerInput() interface{}
	InternalValue() *JobV1SpecTemplateSpec
	SetInternalValue(val *JobV1SpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() JobV1SpecTemplateSpecReadinessGateList
	ReadinessGateInput() interface{}
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	RuntimeClassName() *string
	SetRuntimeClassName(val *string)
	RuntimeClassNameInput() *string
	SchedulerName() *string
	SetSchedulerName(val *string)
	SchedulerNameInput() *string
	SecurityContext() JobV1SpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *JobV1SpecTemplateSpecSecurityContext
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	ShareProcessNamespace() interface{}
	SetShareProcessNamespace(val interface{})
	ShareProcessNamespaceInput() interface{}
	Subdomain() *string
	SetSubdomain(val *string)
	SubdomainInput() *string
	TerminationGracePeriodSeconds() *float64
	SetTerminationGracePeriodSeconds(val *float64)
	TerminationGracePeriodSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Toleration() JobV1SpecTemplateSpecTolerationList
	TolerationInput() interface{}
	TopologySpreadConstraint() JobV1SpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() interface{}
	Volume() JobV1SpecTemplateSpecVolumeList
	VolumeInput() interface{}
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
	PutAffinity(value *JobV1SpecTemplateSpecAffinity)
	PutContainer(value interface{})
	PutDnsConfig(value *JobV1SpecTemplateSpecDnsConfig)
	PutHostAliases(value interface{})
	PutImagePullSecrets(value interface{})
	PutInitContainer(value interface{})
	PutReadinessGate(value interface{})
	PutSecurityContext(value *JobV1SpecTemplateSpecSecurityContext)
	PutToleration(value interface{})
	PutTopologySpreadConstraint(value interface{})
	PutVolume(value interface{})
	ResetActiveDeadlineSeconds()
	ResetAffinity()
	ResetAutomountServiceAccountToken()
	ResetContainer()
	ResetDnsConfig()
	ResetDnsPolicy()
	ResetEnableServiceLinks()
	ResetHostAliases()
	ResetHostIpc()
	ResetHostname()
	ResetHostNetwork()
	ResetHostPid()
	ResetImagePullSecrets()
	ResetInitContainer()
	ResetNodeName()
	ResetNodeSelector()
	ResetPriorityClassName()
	ResetReadinessGate()
	ResetRestartPolicy()
	ResetRuntimeClassName()
	ResetSchedulerName()
	ResetSecurityContext()
	ResetServiceAccountName()
	ResetShareProcessNamespace()
	ResetSubdomain()
	ResetTerminationGracePeriodSeconds()
	ResetToleration()
	ResetTopologySpreadConstraint()
	ResetVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecOutputReference
type jsiiProxy_JobV1SpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Affinity() JobV1SpecTemplateSpecAffinityOutputReference {
	var returns JobV1SpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) AffinityInput() *JobV1SpecTemplateSpecAffinity {
	var returns *JobV1SpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) AutomountServiceAccountToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Container() JobV1SpecTemplateSpecContainerList {
	var returns JobV1SpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) DnsConfig() JobV1SpecTemplateSpecDnsConfigOutputReference {
	var returns JobV1SpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) DnsConfigInput() *JobV1SpecTemplateSpecDnsConfig {
	var returns *JobV1SpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) EnableServiceLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) EnableServiceLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostAliases() JobV1SpecTemplateSpecHostAliasesList {
	var returns JobV1SpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostIpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostIpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostPid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) HostPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ImagePullSecrets() JobV1SpecTemplateSpecImagePullSecretsList {
	var returns JobV1SpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ImagePullSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) InitContainer() JobV1SpecTemplateSpecInitContainerList {
	var returns JobV1SpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) InitContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) InternalValue() *JobV1SpecTemplateSpec {
	var returns *JobV1SpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ReadinessGate() JobV1SpecTemplateSpecReadinessGateList {
	var returns JobV1SpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ReadinessGateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) RuntimeClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) RuntimeClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) SchedulerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) SchedulerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) SecurityContext() JobV1SpecTemplateSpecSecurityContextOutputReference {
	var returns JobV1SpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) SecurityContextInput() *JobV1SpecTemplateSpecSecurityContext {
	var returns *JobV1SpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ShareProcessNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ShareProcessNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Toleration() JobV1SpecTemplateSpecTolerationList {
	var returns JobV1SpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TolerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TopologySpreadConstraint() JobV1SpecTemplateSpecTopologySpreadConstraintList {
	var returns JobV1SpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) TopologySpreadConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Volume() JobV1SpecTemplateSpecVolumeList {
	var returns JobV1SpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecOutputReference_Override(j JobV1SpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetAutomountServiceAccountToken(val interface{}) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetEnableServiceLinks(val interface{}) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetHostIpc(val interface{}) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetHostPid(val interface{}) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetInternalValue(val *JobV1SpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetRuntimeClassName(val *string) {
	if err := j.validateSetRuntimeClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeClassName",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetSchedulerName(val *string) {
	if err := j.validateSetSchedulerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulerName",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetShareProcessNamespace(val interface{}) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutAffinity(value *JobV1SpecTemplateSpecAffinity) {
	if err := j.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAffinity",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutContainer(value interface{}) {
	if err := j.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putContainer",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutDnsConfig(value *JobV1SpecTemplateSpecDnsConfig) {
	if err := j.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDnsConfig",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutHostAliases(value interface{}) {
	if err := j.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHostAliases",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutImagePullSecrets(value interface{}) {
	if err := j.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putImagePullSecrets",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutInitContainer(value interface{}) {
	if err := j.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitContainer",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutReadinessGate(value interface{}) {
	if err := j.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putReadinessGate",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutSecurityContext(value *JobV1SpecTemplateSpecSecurityContext) {
	if err := j.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutToleration(value interface{}) {
	if err := j.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putToleration",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value interface{}) {
	if err := j.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTopologySpreadConstraint",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) PutVolume(value interface{}) {
	if err := j.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		j,
		"resetAffinity",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		j,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		j,
		"resetContainer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		j,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		j,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		j,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		j,
		"resetHostname",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		j,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		j,
		"resetHostPid",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		j,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		j,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		j,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		j,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetRuntimeClassName() {
	_jsii_.InvokeVoid(
		j,
		"resetRuntimeClassName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetSchedulerName() {
	_jsii_.InvokeVoid(
		j,
		"resetSchedulerName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		j,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		j,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		j,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		j,
		"resetToleration",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		j,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

