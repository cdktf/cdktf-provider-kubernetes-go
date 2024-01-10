// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesmutatingwebhookconfigurationv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/datakubernetesmutatingwebhookconfigurationv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference interface {
	cdktf.ComplexObject
	AdmissionReviewVersions() *[]*string
	ClientConfig() DataKubernetesMutatingWebhookConfigurationV1WebhookClientConfigList
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataKubernetesMutatingWebhookConfigurationV1Webhook
	SetInternalValue(val *DataKubernetesMutatingWebhookConfigurationV1Webhook)
	MatchPolicy() *string
	Name() *string
	NamespaceSelector() DataKubernetesMutatingWebhookConfigurationV1WebhookNamespaceSelectorList
	ObjectSelector() DataKubernetesMutatingWebhookConfigurationV1WebhookObjectSelectorList
	ReinvocationPolicy() *string
	Rule() DataKubernetesMutatingWebhookConfigurationV1WebhookRuleList
	SideEffects() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference
type jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) AdmissionReviewVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admissionReviewVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ClientConfig() DataKubernetesMutatingWebhookConfigurationV1WebhookClientConfigList {
	var returns DataKubernetesMutatingWebhookConfigurationV1WebhookClientConfigList
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) InternalValue() *DataKubernetesMutatingWebhookConfigurationV1Webhook {
	var returns *DataKubernetesMutatingWebhookConfigurationV1Webhook
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) MatchPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) NamespaceSelector() DataKubernetesMutatingWebhookConfigurationV1WebhookNamespaceSelectorList {
	var returns DataKubernetesMutatingWebhookConfigurationV1WebhookNamespaceSelectorList
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ObjectSelector() DataKubernetesMutatingWebhookConfigurationV1WebhookObjectSelectorList {
	var returns DataKubernetesMutatingWebhookConfigurationV1WebhookObjectSelectorList
	_jsii_.Get(
		j,
		"objectSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ReinvocationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reinvocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) Rule() DataKubernetesMutatingWebhookConfigurationV1WebhookRuleList {
	var returns DataKubernetesMutatingWebhookConfigurationV1WebhookRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) SideEffects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sideEffects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}


func NewDataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesMutatingWebhookConfigurationV1WebhookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesMutatingWebhookConfigurationV1.DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference_Override(d DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesMutatingWebhookConfigurationV1.DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference)SetInternalValue(val *DataKubernetesMutatingWebhookConfigurationV1Webhook) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesMutatingWebhookConfigurationV1WebhookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

