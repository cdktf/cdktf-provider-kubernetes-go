// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mutatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/mutatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MutatingWebhookConfigurationV1WebhookClientConfigOutputReference interface {
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
	InternalValue() *MutatingWebhookConfigurationV1WebhookClientConfig
	SetInternalValue(val *MutatingWebhookConfigurationV1WebhookClientConfig)
	Service() MutatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference
	ServiceInput() *MutatingWebhookConfigurationV1WebhookClientConfigService
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutService(value *MutatingWebhookConfigurationV1WebhookClientConfigService)
	ResetCaBundle()
	ResetService()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MutatingWebhookConfigurationV1WebhookClientConfigOutputReference
type jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) CaBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) CaBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) InternalValue() *MutatingWebhookConfigurationV1WebhookClientConfig {
	var returns *MutatingWebhookConfigurationV1WebhookClientConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) Service() MutatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference {
	var returns MutatingWebhookConfigurationV1WebhookClientConfigServiceOutputReference
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ServiceInput() *MutatingWebhookConfigurationV1WebhookClientConfigService {
	var returns *MutatingWebhookConfigurationV1WebhookClientConfigService
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewMutatingWebhookConfigurationV1WebhookClientConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MutatingWebhookConfigurationV1WebhookClientConfigOutputReference {
	_init_.Initialize()

	if err := validateNewMutatingWebhookConfigurationV1WebhookClientConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfigurationV1.MutatingWebhookConfigurationV1WebhookClientConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMutatingWebhookConfigurationV1WebhookClientConfigOutputReference_Override(m MutatingWebhookConfigurationV1WebhookClientConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.mutatingWebhookConfigurationV1.MutatingWebhookConfigurationV1WebhookClientConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetCaBundle(val *string) {
	if err := j.validateSetCaBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caBundle",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetInternalValue(val *MutatingWebhookConfigurationV1WebhookClientConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) PutService(value *MutatingWebhookConfigurationV1WebhookClientConfigService) {
	if err := m.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putService",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetCaBundle() {
	_jsii_.InvokeVoid(
		m,
		"resetCaBundle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		m,
		"resetService",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MutatingWebhookConfigurationV1WebhookClientConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

