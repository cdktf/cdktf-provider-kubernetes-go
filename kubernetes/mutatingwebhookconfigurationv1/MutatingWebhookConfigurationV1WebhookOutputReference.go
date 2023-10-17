// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mutatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/mutatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MutatingWebhookConfigurationV1WebhookOutputReference interface {
	cdktf.ComplexObject
	AdmissionReviewVersions() *[]*string
	SetAdmissionReviewVersions(val *[]*string)
	AdmissionReviewVersionsInput() *[]*string
	ClientConfig() MutatingWebhookConfigurationV1WebhookClientConfigOutputReference
	ClientConfigInput() *MutatingWebhookConfigurationV1WebhookClientConfig
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
	FailurePolicy() *string
	SetFailurePolicy(val *string)
	FailurePolicyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchPolicy() *string
	SetMatchPolicy(val *string)
	MatchPolicyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceSelector() MutatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference
	NamespaceSelectorInput() *MutatingWebhookConfigurationV1WebhookNamespaceSelector
	ObjectSelector() MutatingWebhookConfigurationV1WebhookObjectSelectorOutputReference
	ObjectSelectorInput() *MutatingWebhookConfigurationV1WebhookObjectSelector
	ReinvocationPolicy() *string
	SetReinvocationPolicy(val *string)
	ReinvocationPolicyInput() *string
	Rule() MutatingWebhookConfigurationV1WebhookRuleList
	RuleInput() interface{}
	SideEffects() *string
	SetSideEffects(val *string)
	SideEffectsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutClientConfig(value *MutatingWebhookConfigurationV1WebhookClientConfig)
	PutNamespaceSelector(value *MutatingWebhookConfigurationV1WebhookNamespaceSelector)
	PutObjectSelector(value *MutatingWebhookConfigurationV1WebhookObjectSelector)
	PutRule(value interface{})
	ResetAdmissionReviewVersions()
	ResetFailurePolicy()
	ResetMatchPolicy()
	ResetNamespaceSelector()
	ResetObjectSelector()
	ResetReinvocationPolicy()
	ResetRule()
	ResetSideEffects()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MutatingWebhookConfigurationV1WebhookOutputReference
type jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) AdmissionReviewVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) AdmissionReviewVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ClientConfig() MutatingWebhookConfigurationV1WebhookClientConfigOutputReference {
	var returns MutatingWebhookConfigurationV1WebhookClientConfigOutputReference
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ClientConfigInput() *MutatingWebhookConfigurationV1WebhookClientConfig {
	var returns *MutatingWebhookConfigurationV1WebhookClientConfig
	_jsii_.Get(
		j,
		"clientConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) MatchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) NamespaceSelector() MutatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference {
	var returns MutatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) NamespaceSelectorInput() *MutatingWebhookConfigurationV1WebhookNamespaceSelector {
	var returns *MutatingWebhookConfigurationV1WebhookNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ObjectSelector() MutatingWebhookConfigurationV1WebhookObjectSelectorOutputReference {
	var returns MutatingWebhookConfigurationV1WebhookObjectSelectorOutputReference
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ObjectSelectorInput() *MutatingWebhookConfigurationV1WebhookObjectSelector {
	var returns *MutatingWebhookConfigurationV1WebhookObjectSelector
	_jsii_.Get(
		j,
		"objectSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ReinvocationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reinvocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ReinvocationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reinvocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) Rule() MutatingWebhookConfigurationV1WebhookRuleList {
	var returns MutatingWebhookConfigurationV1WebhookRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) SideEffects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) SideEffectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewMutatingWebhookConfigurationV1WebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MutatingWebhookConfigurationV1WebhookOutputReference {
	_init_.Initialize()

	if err := validateNewMutatingWebhookConfigurationV1WebhookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfigurationV1.MutatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMutatingWebhookConfigurationV1WebhookOutputReference_Override(m MutatingWebhookConfigurationV1WebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfigurationV1.MutatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetAdmissionReviewVersions(val *[]*string) {
	if err := j.validateSetAdmissionReviewVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admissionReviewVersions",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetMatchPolicy(val *string) {
	if err := j.validateSetMatchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetReinvocationPolicy(val *string) {
	if err := j.validateSetReinvocationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reinvocationPolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetSideEffects(val *string) {
	if err := j.validateSetSideEffectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sideEffects",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) PutClientConfig(value *MutatingWebhookConfigurationV1WebhookClientConfig) {
	if err := m.validatePutClientConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) PutNamespaceSelector(value *MutatingWebhookConfigurationV1WebhookNamespaceSelector) {
	if err := m.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) PutObjectSelector(value *MutatingWebhookConfigurationV1WebhookObjectSelector) {
	if err := m.validatePutObjectSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putObjectSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) PutRule(value interface{}) {
	if err := m.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetAdmissionReviewVersions() {
	_jsii_.InvokeVoid(
		m,
		"resetAdmissionReviewVersions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetMatchPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetMatchPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetObjectSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetObjectSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetReinvocationPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetReinvocationPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		m,
		"resetRule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetSideEffects() {
	_jsii_.InvokeVoid(
		m,
		"resetSideEffects",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

