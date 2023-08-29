// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	BackoffLimit() *float64
	SetBackoffLimit(val *float64)
	BackoffLimitInput() *float64
	CompletionMode() *string
	SetCompletionMode(val *string)
	CompletionModeInput() *string
	Completions() *float64
	SetCompletions(val *float64)
	CompletionsInput() *float64
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
	InternalValue() *JobV1Spec
	SetInternalValue(val *JobV1Spec)
	ManualSelector() interface{}
	SetManualSelector(val interface{})
	ManualSelectorInput() interface{}
	Parallelism() *float64
	SetParallelism(val *float64)
	ParallelismInput() *float64
	Selector() JobV1SpecSelectorOutputReference
	SelectorInput() *JobV1SpecSelector
	Template() JobV1SpecTemplateOutputReference
	TemplateInput() *JobV1SpecTemplate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TtlSecondsAfterFinished() *string
	SetTtlSecondsAfterFinished(val *string)
	TtlSecondsAfterFinishedInput() *string
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
	PutSelector(value *JobV1SpecSelector)
	PutTemplate(value *JobV1SpecTemplate)
	ResetActiveDeadlineSeconds()
	ResetBackoffLimit()
	ResetCompletionMode()
	ResetCompletions()
	ResetManualSelector()
	ResetParallelism()
	ResetSelector()
	ResetTtlSecondsAfterFinished()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecOutputReference
type jsiiProxy_JobV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) BackoffLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) BackoffLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backoffLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) CompletionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) CompletionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) Completions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) CompletionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) InternalValue() *JobV1Spec {
	var returns *JobV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ManualSelector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ManualSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) Parallelism() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ParallelismInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) Selector() JobV1SpecSelectorOutputReference {
	var returns JobV1SpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) SelectorInput() *JobV1SpecSelector {
	var returns *JobV1SpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) Template() JobV1SpecTemplateOutputReference {
	var returns JobV1SpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) TemplateInput() *JobV1SpecTemplate {
	var returns *JobV1SpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) TtlSecondsAfterFinished() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlSecondsAfterFinished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) TtlSecondsAfterFinishedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlSecondsAfterFinishedInput",
		&returns,
	)
	return returns
}


func NewJobV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecOutputReference_Override(j JobV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetBackoffLimit(val *float64) {
	if err := j.validateSetBackoffLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backoffLimit",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetCompletionMode(val *string) {
	if err := j.validateSetCompletionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completionMode",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetCompletions(val *float64) {
	if err := j.validateSetCompletionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completions",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetInternalValue(val *JobV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetManualSelector(val interface{}) {
	if err := j.validateSetManualSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualSelector",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetParallelism(val *float64) {
	if err := j.validateSetParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelism",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference)SetTtlSecondsAfterFinished(val *string) {
	if err := j.validateSetTtlSecondsAfterFinishedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttlSecondsAfterFinished",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) PutSelector(value *JobV1SpecSelector) {
	if err := j.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSelector",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) PutTemplate(value *JobV1SpecTemplate) {
	if err := j.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTemplate",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetBackoffLimit() {
	_jsii_.InvokeVoid(
		j,
		"resetBackoffLimit",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetCompletionMode() {
	_jsii_.InvokeVoid(
		j,
		"resetCompletionMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetCompletions() {
	_jsii_.InvokeVoid(
		j,
		"resetCompletions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetManualSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetManualSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		j,
		"resetParallelism",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) ResetTtlSecondsAfterFinished() {
	_jsii_.InvokeVoid(
		j,
		"resetTtlSecondsAfterFinished",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

