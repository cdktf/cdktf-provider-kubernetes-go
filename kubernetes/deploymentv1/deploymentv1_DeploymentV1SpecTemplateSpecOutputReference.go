package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() DeploymentV1SpecTemplateSpecAffinityOutputReference
	AffinityInput() *DeploymentV1SpecTemplateSpecAffinity
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
	Container() DeploymentV1SpecTemplateSpecContainerList
	ContainerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() DeploymentV1SpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *DeploymentV1SpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() interface{}
	SetEnableServiceLinks(val interface{})
	EnableServiceLinksInput() interface{}
	// Experimental.
	Fqn() *string
	HostAliases() DeploymentV1SpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() DeploymentV1SpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() interface{}
	InitContainer() DeploymentV1SpecTemplateSpecInitContainerList
	InitContainerInput() interface{}
	InternalValue() *DeploymentV1SpecTemplateSpec
	SetInternalValue(val *DeploymentV1SpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() DeploymentV1SpecTemplateSpecReadinessGateList
	ReadinessGateInput() interface{}
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	RuntimeClassName() *string
	SetRuntimeClassName(val *string)
	RuntimeClassNameInput() *string
	SecurityContext() DeploymentV1SpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *DeploymentV1SpecTemplateSpecSecurityContext
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
	Toleration() DeploymentV1SpecTemplateSpecTolerationList
	TolerationInput() interface{}
	TopologySpreadConstraint() DeploymentV1SpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() interface{}
	Volume() DeploymentV1SpecTemplateSpecVolumeList
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
	PutAffinity(value *DeploymentV1SpecTemplateSpecAffinity)
	PutContainer(value interface{})
	PutDnsConfig(value *DeploymentV1SpecTemplateSpecDnsConfig)
	PutHostAliases(value interface{})
	PutImagePullSecrets(value interface{})
	PutInitContainer(value interface{})
	PutReadinessGate(value interface{})
	PutSecurityContext(value *DeploymentV1SpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for DeploymentV1SpecTemplateSpecOutputReference
type jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Affinity() DeploymentV1SpecTemplateSpecAffinityOutputReference {
	var returns DeploymentV1SpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) AffinityInput() *DeploymentV1SpecTemplateSpecAffinity {
	var returns *DeploymentV1SpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) AutomountServiceAccountToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Container() DeploymentV1SpecTemplateSpecContainerList {
	var returns DeploymentV1SpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) DnsConfig() DeploymentV1SpecTemplateSpecDnsConfigOutputReference {
	var returns DeploymentV1SpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) DnsConfigInput() *DeploymentV1SpecTemplateSpecDnsConfig {
	var returns *DeploymentV1SpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) EnableServiceLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) EnableServiceLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostAliases() DeploymentV1SpecTemplateSpecHostAliasesList {
	var returns DeploymentV1SpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostIpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostIpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostPid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) HostPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ImagePullSecrets() DeploymentV1SpecTemplateSpecImagePullSecretsList {
	var returns DeploymentV1SpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ImagePullSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) InitContainer() DeploymentV1SpecTemplateSpecInitContainerList {
	var returns DeploymentV1SpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) InitContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) InternalValue() *DeploymentV1SpecTemplateSpec {
	var returns *DeploymentV1SpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ReadinessGate() DeploymentV1SpecTemplateSpecReadinessGateList {
	var returns DeploymentV1SpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ReadinessGateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) RuntimeClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) RuntimeClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) SecurityContext() DeploymentV1SpecTemplateSpecSecurityContextOutputReference {
	var returns DeploymentV1SpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) SecurityContextInput() *DeploymentV1SpecTemplateSpecSecurityContext {
	var returns *DeploymentV1SpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ShareProcessNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ShareProcessNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Toleration() DeploymentV1SpecTemplateSpecTolerationList {
	var returns DeploymentV1SpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TolerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TopologySpreadConstraint() DeploymentV1SpecTemplateSpecTopologySpreadConstraintList {
	var returns DeploymentV1SpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) TopologySpreadConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Volume() DeploymentV1SpecTemplateSpecVolumeList {
	var returns DeploymentV1SpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentV1SpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecTemplateSpecOutputReference_Override(d DeploymentV1SpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetAutomountServiceAccountToken(val interface{}) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetEnableServiceLinks(val interface{}) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetHostIpc(val interface{}) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetHostPid(val interface{}) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetInternalValue(val *DeploymentV1SpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetRuntimeClassName(val *string) {
	if err := j.validateSetRuntimeClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeClassName",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetShareProcessNamespace(val interface{}) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutAffinity(value *DeploymentV1SpecTemplateSpecAffinity) {
	if err := d.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAffinity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutContainer(value interface{}) {
	if err := d.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContainer",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutDnsConfig(value *DeploymentV1SpecTemplateSpecDnsConfig) {
	if err := d.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDnsConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutHostAliases(value interface{}) {
	if err := d.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostAliases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutImagePullSecrets(value interface{}) {
	if err := d.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImagePullSecrets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutInitContainer(value interface{}) {
	if err := d.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitContainer",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutReadinessGate(value interface{}) {
	if err := d.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadinessGate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutSecurityContext(value *DeploymentV1SpecTemplateSpecSecurityContext) {
	if err := d.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutToleration(value interface{}) {
	if err := d.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToleration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value interface{}) {
	if err := d.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTopologySpreadConstraint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) PutVolume(value interface{}) {
	if err := d.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		d,
		"resetAffinity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		d,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		d,
		"resetContainer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		d,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		d,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		d,
		"resetHostname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		d,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		d,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		d,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		d,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetRuntimeClassName() {
	_jsii_.InvokeVoid(
		d,
		"resetRuntimeClassName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		d,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		d,
		"resetToleration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

