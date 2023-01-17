package horizontalpodautoscalerv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/horizontalpodautoscalerv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HorizontalPodAutoscalerV2SpecMetricObjectOutputReference interface {
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
	DescribedObject() HorizontalPodAutoscalerV2SpecMetricObjectDescribedObjectOutputReference
	DescribedObjectInput() *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject
	// Experimental.
	Fqn() *string
	InternalValue() *HorizontalPodAutoscalerV2SpecMetricObject
	SetInternalValue(val *HorizontalPodAutoscalerV2SpecMetricObject)
	Metric() HorizontalPodAutoscalerV2SpecMetricObjectMetricOutputReference
	MetricInput() *HorizontalPodAutoscalerV2SpecMetricObjectMetric
	Target() HorizontalPodAutoscalerV2SpecMetricObjectTargetOutputReference
	TargetInput() *HorizontalPodAutoscalerV2SpecMetricObjectTarget
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
	PutDescribedObject(value *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject)
	PutMetric(value *HorizontalPodAutoscalerV2SpecMetricObjectMetric)
	PutTarget(value *HorizontalPodAutoscalerV2SpecMetricObjectTarget)
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerV2SpecMetricObjectOutputReference
type jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) DescribedObject() HorizontalPodAutoscalerV2SpecMetricObjectDescribedObjectOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricObjectDescribedObjectOutputReference
	_jsii_.Get(
		j,
		"describedObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) DescribedObjectInput() *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject {
	var returns *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject
	_jsii_.Get(
		j,
		"describedObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) InternalValue() *HorizontalPodAutoscalerV2SpecMetricObject {
	var returns *HorizontalPodAutoscalerV2SpecMetricObject
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) Metric() HorizontalPodAutoscalerV2SpecMetricObjectMetricOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricObjectMetricOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) MetricInput() *HorizontalPodAutoscalerV2SpecMetricObjectMetric {
	var returns *HorizontalPodAutoscalerV2SpecMetricObjectMetric
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) Target() HorizontalPodAutoscalerV2SpecMetricObjectTargetOutputReference {
	var returns HorizontalPodAutoscalerV2SpecMetricObjectTargetOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) TargetInput() *HorizontalPodAutoscalerV2SpecMetricObjectTarget {
	var returns *HorizontalPodAutoscalerV2SpecMetricObjectTarget
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerV2SpecMetricObjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HorizontalPodAutoscalerV2SpecMetricObjectOutputReference {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerV2SpecMetricObjectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecMetricObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerV2SpecMetricObjectOutputReference_Override(h HorizontalPodAutoscalerV2SpecMetricObjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscalerV2.HorizontalPodAutoscalerV2SpecMetricObjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference)SetInternalValue(val *HorizontalPodAutoscalerV2SpecMetricObject) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) PutDescribedObject(value *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject) {
	if err := h.validatePutDescribedObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDescribedObject",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) PutMetric(value *HorizontalPodAutoscalerV2SpecMetricObjectMetric) {
	if err := h.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putMetric",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) PutTarget(value *HorizontalPodAutoscalerV2SpecMetricObjectTarget) {
	if err := h.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTarget",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		h,
		"resetTarget",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2SpecMetricObjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

