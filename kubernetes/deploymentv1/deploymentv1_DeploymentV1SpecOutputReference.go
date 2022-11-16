package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecOutputReference interface {
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
	InternalValue() *DeploymentV1Spec
	SetInternalValue(val *DeploymentV1Spec)
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
	Selector() DeploymentV1SpecSelectorOutputReference
	SelectorInput() *DeploymentV1SpecSelector
	Strategy() DeploymentV1SpecStrategyOutputReference
	StrategyInput() *DeploymentV1SpecStrategy
	Template() DeploymentV1SpecTemplateOutputReference
	TemplateInput() *DeploymentV1SpecTemplate
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
	PutSelector(value *DeploymentV1SpecSelector)
	PutStrategy(value *DeploymentV1SpecStrategy)
	PutTemplate(value *DeploymentV1SpecTemplate)
	ResetMinReadySeconds()
	ResetPaused()
	ResetProgressDeadlineSeconds()
	ResetReplicas()
	ResetRevisionHistoryLimit()
	ResetSelector()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentV1SpecOutputReference
type jsiiProxy_DeploymentV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) InternalValue() *DeploymentV1Spec {
	var returns *DeploymentV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) MinReadySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) MinReadySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReadySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) ProgressDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"progressDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) ProgressDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"progressDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Replicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) ReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) RevisionHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) RevisionHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Selector() DeploymentV1SpecSelectorOutputReference {
	var returns DeploymentV1SpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) SelectorInput() *DeploymentV1SpecSelector {
	var returns *DeploymentV1SpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Strategy() DeploymentV1SpecStrategyOutputReference {
	var returns DeploymentV1SpecStrategyOutputReference
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) StrategyInput() *DeploymentV1SpecStrategy {
	var returns *DeploymentV1SpecStrategy
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) Template() DeploymentV1SpecTemplateOutputReference {
	var returns DeploymentV1SpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) TemplateInput() *DeploymentV1SpecTemplate {
	var returns *DeploymentV1SpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecOutputReference_Override(d DeploymentV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetInternalValue(val *DeploymentV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetMinReadySeconds(val *float64) {
	if err := j.validateSetMinReadySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReadySeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetProgressDeadlineSeconds(val *float64) {
	if err := j.validateSetProgressDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"progressDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetReplicas(val *string) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetRevisionHistoryLimit(val *float64) {
	if err := j.validateSetRevisionHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) PutSelector(value *DeploymentV1SpecSelector) {
	if err := d.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSelector",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) PutStrategy(value *DeploymentV1SpecStrategy) {
	if err := d.validatePutStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStrategy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) PutTemplate(value *DeploymentV1SpecTemplate) {
	if err := d.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTemplate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetMinReadySeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetMinReadySeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetPaused() {
	_jsii_.InvokeVoid(
		d,
		"resetPaused",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetProgressDeadlineSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetProgressDeadlineSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetRevisionHistoryLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetRevisionHistoryLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		d,
		"resetSelector",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

