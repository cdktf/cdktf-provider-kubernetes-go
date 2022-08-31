// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/jsii"

	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValidatingWebhookConfigurationWebhookOutputReference interface {
	cdktf.ComplexObject
	AdmissionReviewVersions() *[]*string
	SetAdmissionReviewVersions(val *[]*string)
	AdmissionReviewVersionsInput() *[]*string
	ClientConfig() ValidatingWebhookConfigurationWebhookClientConfigOutputReference
	ClientConfigInput() *ValidatingWebhookConfigurationWebhookClientConfig
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
	NamespaceSelector() ValidatingWebhookConfigurationWebhookNamespaceSelectorOutputReference
	NamespaceSelectorInput() *ValidatingWebhookConfigurationWebhookNamespaceSelector
	ObjectSelector() ValidatingWebhookConfigurationWebhookObjectSelectorOutputReference
	ObjectSelectorInput() *ValidatingWebhookConfigurationWebhookObjectSelector
	Rule() ValidatingWebhookConfigurationWebhookRuleList
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
	PutClientConfig(value *ValidatingWebhookConfigurationWebhookClientConfig)
	PutNamespaceSelector(value *ValidatingWebhookConfigurationWebhookNamespaceSelector)
	PutObjectSelector(value *ValidatingWebhookConfigurationWebhookObjectSelector)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingWebhookConfigurationWebhookOutputReference
type jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) AdmissionReviewVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) AdmissionReviewVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ClientConfig() ValidatingWebhookConfigurationWebhookClientConfigOutputReference {
	var returns ValidatingWebhookConfigurationWebhookClientConfigOutputReference
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ClientConfigInput() *ValidatingWebhookConfigurationWebhookClientConfig {
	var returns *ValidatingWebhookConfigurationWebhookClientConfig
	_jsii_.Get(
		j,
		"clientConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) MatchPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) NamespaceSelector() ValidatingWebhookConfigurationWebhookNamespaceSelectorOutputReference {
	var returns ValidatingWebhookConfigurationWebhookNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) NamespaceSelectorInput() *ValidatingWebhookConfigurationWebhookNamespaceSelector {
	var returns *ValidatingWebhookConfigurationWebhookNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ObjectSelector() ValidatingWebhookConfigurationWebhookObjectSelectorOutputReference {
	var returns ValidatingWebhookConfigurationWebhookObjectSelectorOutputReference
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ObjectSelectorInput() *ValidatingWebhookConfigurationWebhookObjectSelector {
	var returns *ValidatingWebhookConfigurationWebhookObjectSelector
	_jsii_.Get(
		j,
		"objectSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) Rule() ValidatingWebhookConfigurationWebhookRuleList {
	var returns ValidatingWebhookConfigurationWebhookRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) SideEffects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) SideEffectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewValidatingWebhookConfigurationWebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ValidatingWebhookConfigurationWebhookOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingWebhookConfigurationWebhookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.ValidatingWebhookConfigurationWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewValidatingWebhookConfigurationWebhookOutputReference_Override(v ValidatingWebhookConfigurationWebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.ValidatingWebhookConfigurationWebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetAdmissionReviewVersions(val *[]*string) {
	if err := j.validateSetAdmissionReviewVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admissionReviewVersions",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetMatchPolicy(val *string) {
	if err := j.validateSetMatchPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPolicy",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetSideEffects(val *string) {
	if err := j.validateSetSideEffectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sideEffects",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) PutClientConfig(value *ValidatingWebhookConfigurationWebhookClientConfig) {
	if err := v.validatePutClientConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClientConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) PutNamespaceSelector(value *ValidatingWebhookConfigurationWebhookNamespaceSelector) {
	if err := v.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) PutObjectSelector(value *ValidatingWebhookConfigurationWebhookObjectSelector) {
	if err := v.validatePutObjectSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putObjectSelector",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) PutRule(value interface{}) {
	if err := v.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRule",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetAdmissionReviewVersions() {
	_jsii_.InvokeVoid(
		v,
		"resetAdmissionReviewVersions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetMatchPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetMatchPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetObjectSelector() {
	_jsii_.InvokeVoid(
		v,
		"resetObjectSelector",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		v,
		"resetRule",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetSideEffects() {
	_jsii_.InvokeVoid(
		v,
		"resetSideEffects",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationWebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

