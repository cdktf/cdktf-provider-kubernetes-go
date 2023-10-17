// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecOutputReference interface {
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
	ConcurrencyPolicy() *string
	SetConcurrencyPolicy(val *string)
	ConcurrencyPolicyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FailedJobsHistoryLimit() *float64
	SetFailedJobsHistoryLimit(val *float64)
	FailedJobsHistoryLimitInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *CronJobV1Spec
	SetInternalValue(val *CronJobV1Spec)
	JobTemplate() CronJobV1SpecJobTemplateOutputReference
	JobTemplateInput() *CronJobV1SpecJobTemplate
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	StartingDeadlineSeconds() *float64
	SetStartingDeadlineSeconds(val *float64)
	StartingDeadlineSecondsInput() *float64
	SuccessfulJobsHistoryLimit() *float64
	SetSuccessfulJobsHistoryLimit(val *float64)
	SuccessfulJobsHistoryLimitInput() *float64
	Suspend() interface{}
	SetSuspend(val interface{})
	SuspendInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	PutJobTemplate(value *CronJobV1SpecJobTemplate)
	ResetConcurrencyPolicy()
	ResetFailedJobsHistoryLimit()
	ResetStartingDeadlineSeconds()
	ResetSuccessfulJobsHistoryLimit()
	ResetSuspend()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobV1SpecOutputReference
type jsiiProxy_CronJobV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) ConcurrencyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) ConcurrencyPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) FailedJobsHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedJobsHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) FailedJobsHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedJobsHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) InternalValue() *CronJobV1Spec {
	var returns *CronJobV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) JobTemplate() CronJobV1SpecJobTemplateOutputReference {
	var returns CronJobV1SpecJobTemplateOutputReference
	_jsii_.Get(
		j,
		"jobTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) JobTemplateInput() *CronJobV1SpecJobTemplate {
	var returns *CronJobV1SpecJobTemplate
	_jsii_.Get(
		j,
		"jobTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) StartingDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) StartingDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) SuccessfulJobsHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulJobsHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) SuccessfulJobsHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulJobsHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) Suspend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) SuspendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobV1SpecOutputReference_Override(c CronJobV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetConcurrencyPolicy(val *string) {
	if err := j.validateSetConcurrencyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetFailedJobsHistoryLimit(val *float64) {
	if err := j.validateSetFailedJobsHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedJobsHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetInternalValue(val *CronJobV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetStartingDeadlineSeconds(val *float64) {
	if err := j.validateSetStartingDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetSuccessfulJobsHistoryLimit(val *float64) {
	if err := j.validateSetSuccessfulJobsHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successfulJobsHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetSuspend(val interface{}) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) PutJobTemplate(value *CronJobV1SpecJobTemplate) {
	if err := c.validatePutJobTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putJobTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetConcurrencyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetConcurrencyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetFailedJobsHistoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetFailedJobsHistoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetStartingDeadlineSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStartingDeadlineSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetSuccessfulJobsHistoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessfulJobsHistoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		c,
		"resetSuspend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

