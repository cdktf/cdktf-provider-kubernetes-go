package mutatingwebhookconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v3/mutatingwebhookconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MutatingWebhookConfigurationWebhookOutputReference interface {
	cdktf.ComplexObject
	AdmissionReviewVersions() *[]*string
	SetAdmissionReviewVersions(val *[]*string)
	AdmissionReviewVersionsInput() *[]*string
	ClientConfig() MutatingWebhookConfigurationWebhookClientConfigOutputReference
	ClientConfigInput() *MutatingWebhookConfigurationWebhookClientConfig
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
	NamespaceSelector() MutatingWebhookConfigurationWebhookNamespaceSelectorOutputReference
	NamespaceSelectorInput() *MutatingWebhookConfigurationWebhookNamespaceSelector
	ObjectSelector() MutatingWebhookConfigurationWebhookObjectSelectorOutputReference
	ObjectSelectorInput() *MutatingWebhookConfigurationWebhookObjectSelector
	ReinvocationPolicy() *string
	SetReinvocationPolicy(val *string)
	ReinvocationPolicyInput() *string
	Rule() MutatingWebhookConfigurationWebhookRuleList
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
	PutClientConfig(value *MutatingWebhookConfigurationWebhookClientConfig)
	PutNamespaceSelector(value *MutatingWebhookConfigurationWebhookNamespaceSelector)
	PutObjectSelector(value *MutatingWebhookConfigurationWebhookObjectSelector)
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

// The jsii proxy struct for MutatingWebhookConfigurationWebhookOutputReference
type jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) AdmissionReviewVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) AdmissionReviewVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ClientConfig() MutatingWebhookConfigurationWebhookClientConfigOutputReference {
	var returns MutatingWebhookConfigurationWebhookClientConfigOutputReference
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ClientConfigInput() *MutatingWebhookConfigurationWebhookClientConfig {
	var returns *MutatingWebhookConfigurationWebhookClientConfig
	_jsii_.Get(
		j,
		"clientConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) MatchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) NamespaceSelector() MutatingWebhookConfigurationWebhookNamespaceSelectorOutputReference {
	var returns MutatingWebhookConfigurationWebhookNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) NamespaceSelectorInput() *MutatingWebhookConfigurationWebhookNamespaceSelector {
	var returns *MutatingWebhookConfigurationWebhookNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ObjectSelector() MutatingWebhookConfigurationWebhookObjectSelectorOutputReference {
	var returns MutatingWebhookConfigurationWebhookObjectSelectorOutputReference
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ObjectSelectorInput() *MutatingWebhookConfigurationWebhookObjectSelector {
	var returns *MutatingWebhookConfigurationWebhookObjectSelector
	_jsii_.Get(
		j,
		"objectSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ReinvocationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reinvocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ReinvocationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reinvocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) Rule() MutatingWebhookConfigurationWebhookRuleList {
	var returns MutatingWebhookConfigurationWebhookRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) SideEffects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) SideEffectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewMutatingWebhookConfigurationWebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MutatingWebhookConfigurationWebhookOutputReference {
	_init_.Initialize()

	if err := validateNewMutatingWebhookConfigurationWebhookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfiguration.MutatingWebhookConfigurationWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMutatingWebhookConfigurationWebhookOutputReference_Override(m MutatingWebhookConfigurationWebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfiguration.MutatingWebhookConfigurationWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetAdmissionReviewVersions(val *[]*string) {
	if err := j.validateSetAdmissionReviewVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admissionReviewVersions",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetMatchPolicy(val *string) {
	if err := j.validateSetMatchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetReinvocationPolicy(val *string) {
	if err := j.validateSetReinvocationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reinvocationPolicy",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetSideEffects(val *string) {
	if err := j.validateSetSideEffectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sideEffects",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) PutClientConfig(value *MutatingWebhookConfigurationWebhookClientConfig) {
	if err := m.validatePutClientConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) PutNamespaceSelector(value *MutatingWebhookConfigurationWebhookNamespaceSelector) {
	if err := m.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) PutObjectSelector(value *MutatingWebhookConfigurationWebhookObjectSelector) {
	if err := m.validatePutObjectSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putObjectSelector",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) PutRule(value interface{}) {
	if err := m.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetAdmissionReviewVersions() {
	_jsii_.InvokeVoid(
		m,
		"resetAdmissionReviewVersions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetMatchPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetMatchPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetObjectSelector() {
	_jsii_.InvokeVoid(
		m,
		"resetObjectSelector",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetReinvocationPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetReinvocationPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		m,
		"resetRule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetSideEffects() {
	_jsii_.InvokeVoid(
		m,
		"resetSideEffects",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

