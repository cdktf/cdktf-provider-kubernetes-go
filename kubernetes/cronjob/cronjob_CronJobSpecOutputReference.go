package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/cronjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobSpecOutputReference interface {
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
	InternalValue() *CronJobSpec
	SetInternalValue(val *CronJobSpec)
	JobTemplate() CronJobSpecJobTemplateOutputReference
	JobTemplateInput() *CronJobSpecJobTemplate
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
	PutJobTemplate(value *CronJobSpecJobTemplate)
	ResetConcurrencyPolicy()
	ResetFailedJobsHistoryLimit()
	ResetStartingDeadlineSeconds()
	ResetSuccessfulJobsHistoryLimit()
	ResetSuspend()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobSpecOutputReference
type jsiiProxy_CronJobSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) ConcurrencyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) ConcurrencyPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) FailedJobsHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedJobsHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) FailedJobsHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedJobsHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) InternalValue() *CronJobSpec {
	var returns *CronJobSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) JobTemplate() CronJobSpecJobTemplateOutputReference {
	var returns CronJobSpecJobTemplateOutputReference
	_jsii_.Get(
		j,
		"jobTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) JobTemplateInput() *CronJobSpecJobTemplate {
	var returns *CronJobSpecJobTemplate
	_jsii_.Get(
		j,
		"jobTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) StartingDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) StartingDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) SuccessfulJobsHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulJobsHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) SuccessfulJobsHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successfulJobsHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) Suspend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) SuspendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCronJobSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobSpecOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobSpecOutputReference_Override(c CronJobSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetConcurrencyPolicy(val *string) {
	if err := j.validateSetConcurrencyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetFailedJobsHistoryLimit(val *float64) {
	if err := j.validateSetFailedJobsHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedJobsHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetInternalValue(val *CronJobSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetStartingDeadlineSeconds(val *float64) {
	if err := j.validateSetStartingDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetSuccessfulJobsHistoryLimit(val *float64) {
	if err := j.validateSetSuccessfulJobsHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successfulJobsHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetSuspend(val interface{}) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecOutputReference) PutJobTemplate(value *CronJobSpecJobTemplate) {
	if err := c.validatePutJobTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putJobTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ResetConcurrencyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetConcurrencyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ResetFailedJobsHistoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetFailedJobsHistoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ResetStartingDeadlineSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetStartingDeadlineSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ResetSuccessfulJobsHistoryLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessfulJobsHistoryLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		c,
		"resetSuspend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

