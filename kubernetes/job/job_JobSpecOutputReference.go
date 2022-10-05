package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecOutputReference interface {
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
	InternalValue() *JobSpec
	SetInternalValue(val *JobSpec)
	ManualSelector() interface{}
	SetManualSelector(val interface{})
	ManualSelectorInput() interface{}
	Parallelism() *float64
	SetParallelism(val *float64)
	ParallelismInput() *float64
	Selector() JobSpecSelectorOutputReference
	SelectorInput() *JobSpecSelector
	Template() JobSpecTemplateOutputReference
	TemplateInput() *JobSpecTemplate
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
	PutSelector(value *JobSpecSelector)
	PutTemplate(value *JobSpecTemplate)
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

// The jsii proxy struct for JobSpecOutputReference
type jsiiProxy_JobSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) BackoffLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) BackoffLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backoffLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) CompletionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) CompletionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) Completions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) CompletionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) InternalValue() *JobSpec {
	var returns *JobSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ManualSelector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ManualSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) Parallelism() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) ParallelismInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) Selector() JobSpecSelectorOutputReference {
	var returns JobSpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) SelectorInput() *JobSpecSelector {
	var returns *JobSpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) Template() JobSpecTemplateOutputReference {
	var returns JobSpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) TemplateInput() *JobSpecTemplate {
	var returns *JobSpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) TtlSecondsAfterFinished() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlSecondsAfterFinished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) TtlSecondsAfterFinishedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlSecondsAfterFinishedInput",
		&returns,
	)
	return returns
}


func NewJobSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobSpecOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobSpecOutputReference_Override(j JobSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetBackoffLimit(val *float64) {
	if err := j.validateSetBackoffLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backoffLimit",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetCompletionMode(val *string) {
	if err := j.validateSetCompletionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completionMode",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetCompletions(val *float64) {
	if err := j.validateSetCompletionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completions",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetInternalValue(val *JobSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetManualSelector(val interface{}) {
	if err := j.validateSetManualSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualSelector",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetParallelism(val *float64) {
	if err := j.validateSetParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelism",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference)SetTtlSecondsAfterFinished(val *string) {
	if err := j.validateSetTtlSecondsAfterFinishedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttlSecondsAfterFinished",
		val,
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecOutputReference) PutSelector(value *JobSpecSelector) {
	if err := j.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSelector",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecOutputReference) PutTemplate(value *JobSpecTemplate) {
	if err := j.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTemplate",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetBackoffLimit() {
	_jsii_.InvokeVoid(
		j,
		"resetBackoffLimit",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetCompletionMode() {
	_jsii_.InvokeVoid(
		j,
		"resetCompletionMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetCompletions() {
	_jsii_.InvokeVoid(
		j,
		"resetCompletions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetManualSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetManualSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		j,
		"resetParallelism",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) ResetTtlSecondsAfterFinished() {
	_jsii_.InvokeVoid(
		j,
		"resetTtlSecondsAfterFinished",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

