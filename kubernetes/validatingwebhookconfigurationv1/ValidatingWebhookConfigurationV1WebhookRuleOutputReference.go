package validatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/validatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValidatingWebhookConfigurationV1WebhookRuleOutputReference interface {
	cdktf.ComplexObject
	ApiGroups() *[]*string
	SetApiGroups(val *[]*string)
	ApiGroupsInput() *[]*string
	ApiVersions() *[]*string
	SetApiVersions(val *[]*string)
	ApiVersionsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operations() *[]*string
	SetOperations(val *[]*string)
	OperationsInput() *[]*string
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingWebhookConfigurationV1WebhookRuleOutputReference
type jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ApiGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ApiGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ApiVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ApiVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) Operations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) OperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewValidatingWebhookConfigurationV1WebhookRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ValidatingWebhookConfigurationV1WebhookRuleOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingWebhookConfigurationV1WebhookRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewValidatingWebhookConfigurationV1WebhookRuleOutputReference_Override(v ValidatingWebhookConfigurationV1WebhookRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetApiGroups(val *[]*string) {
	if err := j.validateSetApiGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiGroups",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetApiVersions(val *[]*string) {
	if err := j.validateSetApiVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersions",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetOperations(val *[]*string) {
	if err := j.validateSetOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		v,
		"resetScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

