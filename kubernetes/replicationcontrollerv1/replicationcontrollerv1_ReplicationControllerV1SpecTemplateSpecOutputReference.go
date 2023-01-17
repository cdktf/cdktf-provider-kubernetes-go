package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() ReplicationControllerV1SpecTemplateSpecAffinityOutputReference
	AffinityInput() *ReplicationControllerV1SpecTemplateSpecAffinity
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
	Container() ReplicationControllerV1SpecTemplateSpecContainerList
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() ReplicationControllerV1SpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *ReplicationControllerV1SpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() interface{}
	SetEnableServiceLinks(val interface{})
	EnableServiceLinksInput() interface{}
	// Experimental.
	Fqn() *string
	HostAliases() ReplicationControllerV1SpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() ReplicationControllerV1SpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() interface{}
	InitContainer() ReplicationControllerV1SpecTemplateSpecInitContainerList
	InitContainerInput() interface{}
	InternalValue() *ReplicationControllerV1SpecTemplateSpec
	SetInternalValue(val *ReplicationControllerV1SpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() ReplicationControllerV1SpecTemplateSpecReadinessGateList
	ReadinessGateInput() interface{}
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	RuntimeClassName() *string
	SetRuntimeClassName(val *string)
	RuntimeClassNameInput() *string
	SecurityContext() ReplicationControllerV1SpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *ReplicationControllerV1SpecTemplateSpecSecurityContext
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
	Toleration() ReplicationControllerV1SpecTemplateSpecTolerationList
	TolerationInput() interface{}
	TopologySpreadConstraint() ReplicationControllerV1SpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() interface{}
	Volume() ReplicationControllerV1SpecTemplateSpecVolumeList
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
	PutAffinity(value *ReplicationControllerV1SpecTemplateSpecAffinity)
	PutContainer(value interface{})
	PutDnsConfig(value *ReplicationControllerV1SpecTemplateSpecDnsConfig)
	PutHostAliases(value interface{})
	PutImagePullSecrets(value interface{})
	PutInitContainer(value interface{})
	PutReadinessGate(value interface{})
	PutSecurityContext(value *ReplicationControllerV1SpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Affinity() ReplicationControllerV1SpecTemplateSpecAffinityOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) AffinityInput() *ReplicationControllerV1SpecTemplateSpecAffinity {
	var returns *ReplicationControllerV1SpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) AutomountServiceAccountToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Container() ReplicationControllerV1SpecTemplateSpecContainerList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) DnsConfig() ReplicationControllerV1SpecTemplateSpecDnsConfigOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) DnsConfigInput() *ReplicationControllerV1SpecTemplateSpecDnsConfig {
	var returns *ReplicationControllerV1SpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) EnableServiceLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) EnableServiceLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostAliases() ReplicationControllerV1SpecTemplateSpecHostAliasesList {
	var returns ReplicationControllerV1SpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostIpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostIpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostPid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) HostPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ImagePullSecrets() ReplicationControllerV1SpecTemplateSpecImagePullSecretsList {
	var returns ReplicationControllerV1SpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ImagePullSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) InitContainer() ReplicationControllerV1SpecTemplateSpecInitContainerList {
	var returns ReplicationControllerV1SpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) InitContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) InternalValue() *ReplicationControllerV1SpecTemplateSpec {
	var returns *ReplicationControllerV1SpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ReadinessGate() ReplicationControllerV1SpecTemplateSpecReadinessGateList {
	var returns ReplicationControllerV1SpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ReadinessGateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) RuntimeClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) RuntimeClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) SecurityContext() ReplicationControllerV1SpecTemplateSpecSecurityContextOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) SecurityContextInput() *ReplicationControllerV1SpecTemplateSpecSecurityContext {
	var returns *ReplicationControllerV1SpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ShareProcessNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ShareProcessNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Toleration() ReplicationControllerV1SpecTemplateSpecTolerationList {
	var returns ReplicationControllerV1SpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TolerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TopologySpreadConstraint() ReplicationControllerV1SpecTemplateSpecTopologySpreadConstraintList {
	var returns ReplicationControllerV1SpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) TopologySpreadConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Volume() ReplicationControllerV1SpecTemplateSpecVolumeList {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerV1SpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetAutomountServiceAccountToken(val interface{}) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetEnableServiceLinks(val interface{}) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetHostIpc(val interface{}) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetHostPid(val interface{}) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetInternalValue(val *ReplicationControllerV1SpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetRuntimeClassName(val *string) {
	if err := j.validateSetRuntimeClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeClassName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetShareProcessNamespace(val interface{}) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutAffinity(value *ReplicationControllerV1SpecTemplateSpecAffinity) {
	if err := r.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAffinity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutContainer(value interface{}) {
	if err := r.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putContainer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutDnsConfig(value *ReplicationControllerV1SpecTemplateSpecDnsConfig) {
	if err := r.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDnsConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutHostAliases(value interface{}) {
	if err := r.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHostAliases",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutImagePullSecrets(value interface{}) {
	if err := r.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putImagePullSecrets",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutInitContainer(value interface{}) {
	if err := r.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putInitContainer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutReadinessGate(value interface{}) {
	if err := r.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putReadinessGate",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutSecurityContext(value *ReplicationControllerV1SpecTemplateSpecSecurityContext) {
	if err := r.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutToleration(value interface{}) {
	if err := r.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putToleration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value interface{}) {
	if err := r.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTopologySpreadConstraint",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) PutVolume(value interface{}) {
	if err := r.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		r,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		r,
		"resetAffinity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		r,
		"resetContainer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		r,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		r,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		r,
		"resetHostname",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		r,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		r,
		"resetHostPid",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		r,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		r,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		r,
		"resetNodeName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		r,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		r,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		r,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetRuntimeClassName() {
	_jsii_.InvokeVoid(
		r,
		"resetRuntimeClassName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		r,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		r,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		r,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		r,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		r,
		"resetToleration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		r,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

