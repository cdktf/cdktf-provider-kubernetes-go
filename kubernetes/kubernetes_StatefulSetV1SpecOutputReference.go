// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/jsii"

	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetV1SpecOutputReference interface {
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
	InternalValue() *StatefulSetV1Spec
	SetInternalValue(val *StatefulSetV1Spec)
	PodManagementPolicy() *string
	SetPodManagementPolicy(val *string)
	PodManagementPolicyInput() *string
	Replicas() *string
	SetReplicas(val *string)
	ReplicasInput() *string
	RevisionHistoryLimit() *float64
	SetRevisionHistoryLimit(val *float64)
	RevisionHistoryLimitInput() *float64
	Selector() StatefulSetV1SpecSelectorOutputReference
	SelectorInput() *StatefulSetV1SpecSelector
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	Template() StatefulSetV1SpecTemplateOutputReference
	TemplateInput() *StatefulSetV1SpecTemplate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateStrategy() StatefulSetV1SpecUpdateStrategyList
	UpdateStrategyInput() interface{}
	VolumeClaimTemplate() StatefulSetV1SpecVolumeClaimTemplateList
	VolumeClaimTemplateInput() interface{}
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
	PutSelector(value *StatefulSetV1SpecSelector)
	PutTemplate(value *StatefulSetV1SpecTemplate)
	PutUpdateStrategy(value interface{})
	PutVolumeClaimTemplate(value interface{})
	ResetPodManagementPolicy()
	ResetReplicas()
	ResetRevisionHistoryLimit()
	ResetUpdateStrategy()
	ResetVolumeClaimTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulSetV1SpecOutputReference
type jsiiProxy_StatefulSetV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) InternalValue() *StatefulSetV1Spec {
	var returns *StatefulSetV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) PodManagementPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podManagementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) PodManagementPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podManagementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) Replicas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) ReplicasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) RevisionHistoryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) RevisionHistoryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revisionHistoryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) Selector() StatefulSetV1SpecSelectorOutputReference {
	var returns StatefulSetV1SpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) SelectorInput() *StatefulSetV1SpecSelector {
	var returns *StatefulSetV1SpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) Template() StatefulSetV1SpecTemplateOutputReference {
	var returns StatefulSetV1SpecTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) TemplateInput() *StatefulSetV1SpecTemplate {
	var returns *StatefulSetV1SpecTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) UpdateStrategy() StatefulSetV1SpecUpdateStrategyList {
	var returns StatefulSetV1SpecUpdateStrategyList
	_jsii_.Get(
		j,
		"updateStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) UpdateStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) VolumeClaimTemplate() StatefulSetV1SpecVolumeClaimTemplateList {
	var returns StatefulSetV1SpecVolumeClaimTemplateList
	_jsii_.Get(
		j,
		"volumeClaimTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference) VolumeClaimTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeClaimTemplateInput",
		&returns,
	)
	return returns
}


func NewStatefulSetV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.StatefulSetV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetV1SpecOutputReference_Override(s StatefulSetV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.StatefulSetV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetInternalValue(val *StatefulSetV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetPodManagementPolicy(val *string) {
	if err := j.validateSetPodManagementPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podManagementPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetReplicas(val *string) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetRevisionHistoryLimit(val *float64) {
	if err := j.validateSetRevisionHistoryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionHistoryLimit",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) PutSelector(value *StatefulSetV1SpecSelector) {
	if err := s.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSelector",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) PutTemplate(value *StatefulSetV1SpecTemplate) {
	if err := s.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) PutUpdateStrategy(value interface{}) {
	if err := s.validatePutUpdateStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpdateStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) PutVolumeClaimTemplate(value interface{}) {
	if err := s.validatePutVolumeClaimTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVolumeClaimTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ResetPodManagementPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPodManagementPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicas",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ResetRevisionHistoryLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetRevisionHistoryLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ResetUpdateStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdateStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ResetVolumeClaimTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeClaimTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

