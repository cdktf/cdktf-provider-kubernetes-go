// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecOutputReference interface {
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
	InternalValue() *DeploymentSpec
	SetInternalValue(val *DeploymentSpec)
	MinReadySeconds() *float64
	SetMinReadySeconds(val *float64)
	MinReadySecondsInput() *float64
	Paused() interface{}
	SetPaused(val interface{})
	PausedInput() interface{}
	ProgressDeadlineSeconds() *float64
	SetProgressDeadlineSeconds(val *float64)
	ProgressDeadlineSecondsInput() *float64
	Replicas() *string
	SetReplicas(val *string)
	ReplicasInput() *string
	RevisionHistoryLimit() *float64
	SetRevisionHistoryLimit(val *float64)
	RevisionHistoryLimitInput() *float64
	Selector() DeploymentSpecSelectorOutputReference
	SelectorInput() *DeploymentSpecSelector
	Strategy() DeploymentSpecStrategyOutputReference
	StrategyInput() *DeploymentSpecStrategy
	Template() DeploymentSpecTemplateOutputReference
	TemplateInput() *DeploymentSpecTemplate
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutSelector(value *DeploymentSpecSelector)
	PutStrategy(value *DeploymentSpecStrategy)
	PutTemplate(value *DeploymentSpecTemplate)
	ResetMinReadySeconds()
	ResetPaused()
	ResetProgressDeadlineSeconds()
	ResetReplicas()
	ResetRevisionHistoryLimit()
	ResetSelector()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentSpecOutputReference
type jsiiProxy_DeploymentSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) InternalValue() *DeploymentSpec {
	var returns *DeploymentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) MinReadySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) MinReadySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) ProgressDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"progressDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) ProgressDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"progressDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Replicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) ReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) RevisionHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) RevisionHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Selector() DeploymentSpecSelectorOutputReference {
	var returns DeploymentSpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) SelectorInput() *DeploymentSpecSelector {
	var returns *DeploymentSpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Strategy() DeploymentSpecStrategyOutputReference {
	var returns DeploymentSpecStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) StrategyInput() *DeploymentSpecStrategy {
	var returns *DeploymentSpecStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) Template() DeploymentSpecTemplateOutputReference {
	var returns DeploymentSpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) TemplateInput() *DeploymentSpecTemplate {
	var returns *DeploymentSpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentSpecOutputReference_Override(d DeploymentSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetInternalValue(val *DeploymentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetMinReadySeconds(val *float64) {
	if err := j.validateSetMinReadySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReadySeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetProgressDeadlineSeconds(val *float64) {
	if err := j.validateSetProgressDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"progressDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetReplicas(val *string) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetRevisionHistoryLimit(val *float64) {
	if err := j.validateSetRevisionHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecOutputReference) PutSelector(value *DeploymentSpecSelector) {
	if err := d.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSelector",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) PutStrategy(value *DeploymentSpecStrategy) {
	if err := d.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStrategy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) PutTemplate(value *DeploymentSpecTemplate) {
	if err := d.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTemplate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetMinReadySeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetMinReadySeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetPaused() {
	_jsii_.InvokeVoid(
		d,
		"resetPaused",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetProgressDeadlineSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetProgressDeadlineSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetRevisionHistoryLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetRevisionHistoryLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

