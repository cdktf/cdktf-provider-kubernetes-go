// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/service/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceSpecOutputReference interface {
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
	InternalValue() *ServiceSpec
	SetInternalValue(val *ServiceSpec)
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
	Port() ServiceSpecPortList
	PortInput() interface{}
	PublishNotReadyAddresses() interface{}
	SetPublishNotReadyAddresses(val interface{})
	PublishNotReadyAddressesInput() interface{}
	Selector() *map[string]*string
	SetSelector(val *map[string]*string)
	SelectorInput() *map[string]*string
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityConfig() ServiceSpecSessionAffinityConfigOutputReference
	SessionAffinityConfigInput() *ServiceSpecSessionAffinityConfig
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
	PutSessionAffinityConfig(value *ServiceSpecSessionAffinityConfig)
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

// The jsii proxy struct for ServiceSpecOutputReference
type jsiiProxy_ServiceSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceSpecOutputReference) AllocateLoadBalancerNodePorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocateLoadBalancerNodePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) AllocateLoadBalancerNodePortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocateLoadBalancerNodePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ClusterIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ClusterIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ClusterIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ClusterIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) ExternalTrafficPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTrafficPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) HealthCheckNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) HealthCheckNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) InternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) InternalTrafficPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTrafficPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) InternalValue() *ServiceSpec {
	var returns *ServiceSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) IpFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) IpFamiliesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) IpFamilyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) IpFamilyPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerSourceRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSourceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) LoadBalancerSourceRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSourceRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) Port() ServiceSpecPortList {
	var returns ServiceSpecPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) PortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) PublishNotReadyAddresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishNotReadyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) PublishNotReadyAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishNotReadyAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) Selector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) SelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) SessionAffinityConfig() ServiceSpecSessionAffinityConfigOutputReference {
	var returns ServiceSpecSessionAffinityConfigOutputReference
	_jsii_.Get(
		j,
		"sessionAffinityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) SessionAffinityConfigInput() *ServiceSpecSessionAffinityConfig {
	var returns *ServiceSpecSessionAffinityConfig
	_jsii_.Get(
		j,
		"sessionAffinityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceSpecOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewServiceSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceSpecOutputReference {
	_init_.Initialize()

	if err := validateNewServiceSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.service.ServiceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceSpecOutputReference_Override(s ServiceSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.service.ServiceSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetAllocateLoadBalancerNodePorts(val interface{}) {
	if err := j.validateSetAllocateLoadBalancerNodePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocateLoadBalancerNodePorts",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetClusterIp(val *string) {
	if err := j.validateSetClusterIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIp",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetClusterIps(val *[]*string) {
	if err := j.validateSetClusterIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIps",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetExternalIps(val *[]*string) {
	if err := j.validateSetExternalIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIps",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetExternalTrafficPolicy(val *string) {
	if err := j.validateSetExternalTrafficPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTrafficPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetHealthCheckNodePort(val *float64) {
	if err := j.validateSetHealthCheckNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckNodePort",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetInternalTrafficPolicy(val *string) {
	if err := j.validateSetInternalTrafficPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalTrafficPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetInternalValue(val *ServiceSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetIpFamilies(val *[]*string) {
	if err := j.validateSetIpFamiliesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamilies",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetIpFamilyPolicy(val *string) {
	if err := j.validateSetIpFamilyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamilyPolicy",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetLoadBalancerClass(val *string) {
	if err := j.validateSetLoadBalancerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerClass",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetLoadBalancerIp(val *string) {
	if err := j.validateSetLoadBalancerIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerIp",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetLoadBalancerSourceRanges(val *[]*string) {
	if err := j.validateSetLoadBalancerSourceRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerSourceRanges",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetPublishNotReadyAddresses(val interface{}) {
	if err := j.validateSetPublishNotReadyAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishNotReadyAddresses",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetSelector(val *map[string]*string) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceSpecOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceSpecOutputReference) PutPort(value interface{}) {
	if err := s.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPort",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) PutSessionAffinityConfig(value *ServiceSpecSessionAffinityConfig) {
	if err := s.validatePutSessionAffinityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSessionAffinityConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetAllocateLoadBalancerNodePorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllocateLoadBalancerNodePorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetClusterIp() {
	_jsii_.InvokeVoid(
		s,
		"resetClusterIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetClusterIps() {
	_jsii_.InvokeVoid(
		s,
		"resetClusterIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetExternalIps() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetExternalName() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetExternalTrafficPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalTrafficPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetHealthCheckNodePort() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckNodePort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetInternalTrafficPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetInternalTrafficPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetIpFamilies() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFamilies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetIpFamilyPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFamilyPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetLoadBalancerClass() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetLoadBalancerIp() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetLoadBalancerSourceRanges() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadBalancerSourceRanges",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetPublishNotReadyAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetPublishNotReadyAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetSessionAffinityConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionAffinityConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServiceSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

