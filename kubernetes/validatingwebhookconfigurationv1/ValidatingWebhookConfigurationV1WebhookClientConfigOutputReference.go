// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/validatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference interface {
	cdktf.ComplexObject
	CaBundle() *string
	SetCaBundle(val *string)
	CaBundleInput() *string
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
	InternalValue() *ValidatingWebhookConfigurationV1WebhookClientConfig
	SetInternalValue(val *ValidatingWebhookConfigurationV1WebhookClientConfig)
	Service() ValidatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference
	ServiceInput() *ValidatingWebhookConfigurationV1WebhookClientConfigService
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutService(value *ValidatingWebhookConfigurationV1WebhookClientConfigService)
	ResetCaBundle()
	ResetService()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference
type jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) CaBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) CaBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) InternalValue() *ValidatingWebhookConfigurationV1WebhookClientConfig {
	var returns *ValidatingWebhookConfigurationV1WebhookClientConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) Service() ValidatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference {
	var returns ValidatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ServiceInput() *ValidatingWebhookConfigurationV1WebhookClientConfigService {
	var returns *ValidatingWebhookConfigurationV1WebhookClientConfigService
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewValidatingWebhookConfigurationV1WebhookClientConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference {
	_init_.Initialize()

	if err := validateNewValidatingWebhookConfigurationV1WebhookClientConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewValidatingWebhookConfigurationV1WebhookClientConfigOutputReference_Override(v ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.validatingWebhookConfigurationV1.ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetCaBundle(val *string) {
	if err := j.validateSetCaBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caBundle",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetInternalValue(val *ValidatingWebhookConfigurationV1WebhookClientConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) PutService(value *ValidatingWebhookConfigurationV1WebhookClientConfigService) {
	if err := v.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putService",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetCaBundle() {
	_jsii_.InvokeVoid(
		v,
		"resetCaBundle",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		v,
		"resetService",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		v,
		"resetUrl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookClientConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

