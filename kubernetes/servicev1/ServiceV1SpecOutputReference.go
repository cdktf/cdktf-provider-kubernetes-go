// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/servicev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceV1SpecOutputReference interface {
	cdktf.ComplexObject
	AllocateLoadBalancerNodePorts() interface{}
	SetAllocateLoadBalancerNodePorts(val interface{})
	AllocateLoadBalancerNodePortsInput() interface{}
	ClusterIp() *string
	SetClusterIp(val *string)
	ClusterIpInput() *string
	ClusterIps() *[]*string
	SetClusterIps(val *[]*string)
	ClusterIpsInput() *[]*string
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
	ExternalIps() *[]*string
	SetExternalIps(val *[]*string)
	ExternalIpsInput() *[]*string
	ExternalName() *string
	SetExternalName(val *string)
	ExternalNameInput() *string
	ExternalTrafficPolicy() *string
	SetExternalTrafficPolicy(val *string)
	ExternalTrafficPolicyInput() *string
	// Experimental.
	Fqn() *string
	HealthCheckNodePort() *float64
	SetHealthCheckNodePort(val *float64)
	HealthCheckNodePortInput() *float64
	InternalTrafficPolicy() *string
	SetInternalTrafficPolicy(val *string)
	InternalTrafficPolicyInput() *string
	InternalValue() *ServiceV1Spec
	SetInternalValue(val *ServiceV1Spec)
	IpFamilies() *[]*string
	SetIpFamilies(val *[]*string)
	IpFamiliesInput() *[]*string
	IpFamilyPolicy() *string
	SetIpFamilyPolicy(val *string)
	IpFamilyPolicyInput() *string
	LoadBalancerClass() *string
	SetLoadBalancerClass(val *string)
	LoadBalancerClassInput() *string
	LoadBalancerIp() *string
	SetLoadBalancerIp(val *string)
	LoadBalancerIpInput() *string
	LoadBalancerSourceRanges() *[]*string
	SetLoadBalancerSourceRanges(val *[]*string)
	LoadBalancerSourceRangesInput() *[]*string
	Port() ServiceV1SpecPortList
	PortInput() interface{}
	PublishNotReadyAddresses() interface{}
	SetPublishNotReadyAddresses(val interface{})
	PublishNotReadyAddressesInput() interface{}
	Selector() *map[string]*string
	SetSelector(val *map[string]*string)
	SelectorInput() *map[string]*string
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityConfig() ServiceV1SpecSessionAffinityConfigOutputReference
	SessionAffinityConfigInput() *ServiceV1SpecSessionAffinityConfig
	SessionAffinityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutPort(value interface{})
	PutSessionAffinityConfig(value *ServiceV1SpecSessionAffinityConfig)
	ResetAllocateLoadBalancerNodePorts()
	ResetClusterIp()
	ResetClusterIps()
	ResetExternalIps()
	ResetExternalName()
	ResetExternalTrafficPolicy()
	ResetHealthCheckNodePort()
	ResetInternalTrafficPolicy()
	ResetIpFamilies()
	ResetIpFamilyPolicy()
	ResetLoadBalancerClass()
	ResetLoadBalancerIp()
	ResetLoadBalancerSourceRanges()
	ResetPort()
	ResetPublishNotReadyAddresses()
	ResetSelector()
	ResetSessionAffinity()
	ResetSessionAffinityConfig()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceV1SpecOutputReference
type jsiiProxy_ServiceV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) AllocateLoadBalancerNodePorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocateLoadBalancerNodePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) AllocateLoadBalancerNodePortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocateLoadBalancerNodePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ClusterIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ClusterIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ClusterIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ClusterIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) ExternalTrafficPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTrafficPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) HealthCheckNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) HealthCheckNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) InternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) InternalTrafficPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTrafficPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) InternalValue() *ServiceV1Spec {
	var returns *ServiceV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) IpFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) IpFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) IpFamilyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) IpFamilyPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerSourceRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSourceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) LoadBalancerSourceRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSourceRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) Port() ServiceV1SpecPortList {
	var returns ServiceV1SpecPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) PublishNotReadyAddresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishNotReadyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) PublishNotReadyAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishNotReadyAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) Selector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) SelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) SessionAffinityConfig() ServiceV1SpecSessionAffinityConfigOutputReference {
	var returns ServiceV1SpecSessionAffinityConfigOutputReference
	_jsii_.Get(
		j,
		"sessionAffinityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) SessionAffinityConfigInput() *ServiceV1SpecSessionAffinityConfig {
	var returns *ServiceV1SpecSessionAffinityConfig
	_jsii_.Get(
		j,
		"sessionAffinityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceV1SpecOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewServiceV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewServiceV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.serviceV1.ServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceV1SpecOutputReference_Override(s ServiceV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.serviceV1.ServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetAllocateLoadBalancerNodePorts(val interface{}) {
	if err := j.validateSetAllocateLoadBalancerNodePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocateLoadBalancerNodePorts",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetClusterIp(val *string) {
	if err := j.validateSetClusterIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIp",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetClusterIps(val *[]*string) {
	if err := j.validateSetClusterIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIps",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetExternalIps(val *[]*string) {
	if err := j.validateSetExternalIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIps",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetExternalTrafficPolicy(val *string) {
	if err := j.validateSetExternalTrafficPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTrafficPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetHealthCheckNodePort(val *float64) {
	if err := j.validateSetHealthCheckNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckNodePort",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetInternalTrafficPolicy(val *string) {
	if err := j.validateSetInternalTrafficPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalTrafficPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetInternalValue(val *ServiceV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetIpFamilies(val *[]*string) {
	if err := j.validateSetIpFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamilies",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetIpFamilyPolicy(val *string) {
	if err := j.validateSetIpFamilyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamilyPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetLoadBalancerClass(val *string) {
	if err := j.validateSetLoadBalancerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerClass",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetLoadBalancerIp(val *string) {
	if err := j.validateSetLoadBalancerIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerIp",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetLoadBalancerSourceRanges(val *[]*string) {
	if err := j.validateSetLoadBalancerSourceRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerSourceRanges",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetPublishNotReadyAddresses(val interface{}) {
	if err := j.validateSetPublishNotReadyAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishNotReadyAddresses",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetSelector(val *map[string]*string) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceV1SpecOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) PutPort(value interface{}) {
	if err := s.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPort",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) PutSessionAffinityConfig(value *ServiceV1SpecSessionAffinityConfig) {
	if err := s.validatePutSessionAffinityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSessionAffinityConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetAllocateLoadBalancerNodePorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllocateLoadBalancerNodePorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetClusterIp() {
	_jsii_.InvokeVoid(
		s,
		"resetClusterIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetClusterIps() {
	_jsii_.InvokeVoid(
		s,
		"resetClusterIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetExternalIps() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetExternalName() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetExternalTrafficPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalTrafficPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetHealthCheckNodePort() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckNodePort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetInternalTrafficPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetInternalTrafficPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetIpFamilies() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFamilies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetIpFamilyPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFamilyPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetLoadBalancerClass() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetLoadBalancerIp() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetLoadBalancerSourceRanges() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerSourceRanges",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetPublishNotReadyAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetPublishNotReadyAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetSessionAffinityConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionAffinityConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

