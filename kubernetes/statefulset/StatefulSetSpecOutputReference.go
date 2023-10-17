// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/statefulset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetSpecOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *StatefulSetSpec
	SetInternalValue(val *StatefulSetSpec)
	PodManagementPolicy() *string
	SetPodManagementPolicy(val *string)
	PodManagementPolicyInput() *string
	Replicas() *string
	SetReplicas(val *string)
	ReplicasInput() *string
	RevisionHistoryLimit() *float64
	SetRevisionHistoryLimit(val *float64)
	RevisionHistoryLimitInput() *float64
	Selector() StatefulSetSpecSelectorOutputReference
	SelectorInput() *StatefulSetSpecSelector
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	Template() StatefulSetSpecTemplateOutputReference
	TemplateInput() *StatefulSetSpecTemplate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateStrategy() StatefulSetSpecUpdateStrategyList
	UpdateStrategyInput() interface{}
	VolumeClaimTemplate() StatefulSetSpecVolumeClaimTemplateList
	VolumeClaimTemplateInput() interface{}
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
	PutSelector(value *StatefulSetSpecSelector)
	PutTemplate(value *StatefulSetSpecTemplate)
	PutUpdateStrategy(value interface{})
	PutVolumeClaimTemplate(value interface{})
	ResetPodManagementPolicy()
	ResetReplicas()
	ResetRevisionHistoryLimit()
	ResetUpdateStrategy()
	ResetVolumeClaimTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulSetSpecOutputReference
type jsiiProxy_StatefulSetSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) InternalValue() *StatefulSetSpec {
	var returns *StatefulSetSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) PodManagementPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podManagementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) PodManagementPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podManagementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) Replicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) ReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) RevisionHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) RevisionHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) Selector() StatefulSetSpecSelectorOutputReference {
	var returns StatefulSetSpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) SelectorInput() *StatefulSetSpecSelector {
	var returns *StatefulSetSpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) Template() StatefulSetSpecTemplateOutputReference {
	var returns StatefulSetSpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) TemplateInput() *StatefulSetSpecTemplate {
	var returns *StatefulSetSpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) UpdateStrategy() StatefulSetSpecUpdateStrategyList {
	var returns StatefulSetSpecUpdateStrategyList
	_jsii_.Get(
		j,
		"updateStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) UpdateStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) VolumeClaimTemplate() StatefulSetSpecVolumeClaimTemplateList {
	var returns StatefulSetSpecVolumeClaimTemplateList
	_jsii_.Get(
		j,
		"volumeClaimTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecOutputReference) VolumeClaimTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeClaimTemplateInput",
		&returns,
	)
	return returns
}


func NewStatefulSetSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetSpecOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetSpecOutputReference_Override(s StatefulSetSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetInternalValue(val *StatefulSetSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetPodManagementPolicy(val *string) {
	if err := j.validateSetPodManagementPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podManagementPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetReplicas(val *string) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetRevisionHistoryLimit(val *float64) {
	if err := j.validateSetRevisionHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) PutSelector(value *StatefulSetSpecSelector) {
	if err := s.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSelector",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) PutTemplate(value *StatefulSetSpecTemplate) {
	if err := s.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) PutUpdateStrategy(value interface{}) {
	if err := s.validatePutUpdateStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpdateStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) PutVolumeClaimTemplate(value interface{}) {
	if err := s.validatePutVolumeClaimTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVolumeClaimTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ResetPodManagementPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPodManagementPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicas",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ResetRevisionHistoryLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetRevisionHistoryLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ResetUpdateStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdateStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) ResetVolumeClaimTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeClaimTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulSetSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

