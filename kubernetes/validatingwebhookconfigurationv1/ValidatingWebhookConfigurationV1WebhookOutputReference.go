// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/validatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValidatingWebhookConfigurationV1WebhookOutputReference interface {
	cdktf.ComplexObject
	AdmissionReviewVersions() *[]*string
	SetAdmissionReviewVersions(val *[]*string)
	AdmissionReviewVersionsInput() *[]*string
	ClientConfig() ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference
	ClientConfigInput() *ValidatingWebhookConfigurationV1WebhookClientConfig
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
	NamespaceSelector() ValidatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference
	NamespaceSelectorInput() *ValidatingWebhookConfigurationV1WebhookNamespaceSelector
	ObjectSelector() ValidatingWebhookConfigurationV1WebhookObjectSelectorOutputReference
	ObjectSelectorInput() *ValidatingWebhookConfigurationV1WebhookObjectSelector
	Rule() ValidatingWebhookConfigurationV1WebhookRuleList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutClientConfig(value *ValidatingWebhookConfigurationV1WebhookClientConfig)
	PutNamespaceSelector(value *ValidatingWebhookConfigurationV1WebhookNamespaceSelector)
	PutObjectSelector(value *ValidatingWebhookConfigurationV1WebhookObjectSelector)
	PutRule(value interface{})
	ResetAdmissionReviewVersions()
	ResetFailurePolicy()
	ResetMatchPolicy()
	ResetNamespaceSelector()
	ResetObjectSelector()
	ResetRule()
	ResetSideEffects()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingWebhookConfigurationV1WebhookOutputReference
type jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) AdmissionReviewVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) AdmissionReviewVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ClientConfig() ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference {
	var returns ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ClientConfigInput() *ValidatingWebhookConfigurationV1WebhookClientConfig {
	var returns *ValidatingWebhookConfigurationV1WebhookClientConfig
	_jsii_.Get(
		j,
		"clientConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) MatchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) NamespaceSelector() ValidatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference {
	var returns ValidatingWebhookConfigurationV1WebhookNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) NamespaceSelectorInput() *ValidatingWebhookConfigurationV1WebhookNamespaceSelector {
	var returns *ValidatingWebhookConfigurationV1WebhookNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ObjectSelector() ValidatingWebhookConfigurationV1WebhookObjectSelectorOutputReference {
	var returns ValidatingWebhookConfigurationV1WebhookObjectSelectorOutputReference
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ObjectSelectorInput() *ValidatingWebhookConfigurationV1WebhookObjectSelector {
	var returns *ValidatingWebhookConfigurationV1WebhookObjectSelector
	_jsii_.Get(
		j,
		"objectSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) Rule() ValidatingWebhookConfigurationV1WebhookRuleList {
	var returns ValidatingWebhookConfigurationV1WebhookRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) SideEffects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) SideEffectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewValidatingWebhookConfigurationV1WebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ValidatingWebhookConfigurationV1WebhookOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingWebhookConfigurationV1WebhookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewValidatingWebhookConfigurationV1WebhookOutputReference_Override(v ValidatingWebhookConfigurationV1WebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetAdmissionReviewVersions(val *[]*string) {
	if err := j.validateSetAdmissionReviewVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admissionReviewVersions",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetMatchPolicy(val *string) {
	if err := j.validateSetMatchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetSideEffects(val *string) {
	if err := j.validateSetSideEffectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sideEffects",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) PutClientConfig(value *ValidatingWebhookConfigurationV1WebhookClientConfig) {
	if err := v.validatePutClientConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClientConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) PutNamespaceSelector(value *ValidatingWebhookConfigurationV1WebhookNamespaceSelector) {
	if err := v.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) PutObjectSelector(value *ValidatingWebhookConfigurationV1WebhookObjectSelector) {
	if err := v.validatePutObjectSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putObjectSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) PutRule(value interface{}) {
	if err := v.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRule",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetAdmissionReviewVersions() {
	_jsii_.InvokeVoid(
		v,
		"resetAdmissionReviewVersions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetMatchPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetMatchPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetObjectSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetObjectSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		v,
		"resetRule",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetSideEffects() {
	_jsii_.InvokeVoid(
		v,
		"resetSideEffects",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

