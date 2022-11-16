package horizontalpodautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/horizontalpodautoscaler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HorizontalPodAutoscalerSpecMetricOutputReference interface {
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
	ContainerResource() HorizontalPodAutoscalerSpecMetricContainerResourceOutputReference
	ContainerResourceInput() *HorizontalPodAutoscalerSpecMetricContainerResource
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	External() HorizontalPodAutoscalerSpecMetricExternalOutputReference
	ExternalInput() *HorizontalPodAutoscalerSpecMetricExternal
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Object() HorizontalPodAutoscalerSpecMetricObjectOutputReference
	ObjectInput() *HorizontalPodAutoscalerSpecMetricObject
	Pods() HorizontalPodAutoscalerSpecMetricPodsOutputReference
	PodsInput() *HorizontalPodAutoscalerSpecMetricPods
	Resource() HorizontalPodAutoscalerSpecMetricResourceOutputReference
	ResourceInput() *HorizontalPodAutoscalerSpecMetricResource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutContainerResource(value *HorizontalPodAutoscalerSpecMetricContainerResource)
	PutExternal(value *HorizontalPodAutoscalerSpecMetricExternal)
	PutObject(value *HorizontalPodAutoscalerSpecMetricObject)
	PutPods(value *HorizontalPodAutoscalerSpecMetricPods)
	PutResource(value *HorizontalPodAutoscalerSpecMetricResource)
	ResetContainerResource()
	ResetExternal()
	ResetObject()
	ResetPods()
	ResetResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizontalPodAutoscalerSpecMetricOutputReference
type jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ContainerResource() HorizontalPodAutoscalerSpecMetricContainerResourceOutputReference {
	var returns HorizontalPodAutoscalerSpecMetricContainerResourceOutputReference
	_jsii_.Get(
		j,
		"containerResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ContainerResourceInput() *HorizontalPodAutoscalerSpecMetricContainerResource {
	var returns *HorizontalPodAutoscalerSpecMetricContainerResource
	_jsii_.Get(
		j,
		"containerResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) External() HorizontalPodAutoscalerSpecMetricExternalOutputReference {
	var returns HorizontalPodAutoscalerSpecMetricExternalOutputReference
	_jsii_.Get(
		j,
		"external",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ExternalInput() *HorizontalPodAutoscalerSpecMetricExternal {
	var returns *HorizontalPodAutoscalerSpecMetricExternal
	_jsii_.Get(
		j,
		"externalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Object() HorizontalPodAutoscalerSpecMetricObjectOutputReference {
	var returns HorizontalPodAutoscalerSpecMetricObjectOutputReference
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ObjectInput() *HorizontalPodAutoscalerSpecMetricObject {
	var returns *HorizontalPodAutoscalerSpecMetricObject
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Pods() HorizontalPodAutoscalerSpecMetricPodsOutputReference {
	var returns HorizontalPodAutoscalerSpecMetricPodsOutputReference
	_jsii_.Get(
		j,
		"pods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PodsInput() *HorizontalPodAutoscalerSpecMetricPods {
	var returns *HorizontalPodAutoscalerSpecMetricPods
	_jsii_.Get(
		j,
		"podsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Resource() HorizontalPodAutoscalerSpecMetricResourceOutputReference {
	var returns HorizontalPodAutoscalerSpecMetricResourceOutputReference
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResourceInput() *HorizontalPodAutoscalerSpecMetricResource {
	var returns *HorizontalPodAutoscalerSpecMetricResource
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewHorizontalPodAutoscalerSpecMetricOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HorizontalPodAutoscalerSpecMetricOutputReference {
	_init_.Initialize()

	if err := validateNewHorizontalPodAutoscalerSpecMetricOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscaler.HorizontalPodAutoscalerSpecMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHorizontalPodAutoscalerSpecMetricOutputReference_Override(h HorizontalPodAutoscalerSpecMetricOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.horizontalPodAutoscaler.HorizontalPodAutoscalerSpecMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PutContainerResource(value *HorizontalPodAutoscalerSpecMetricContainerResource) {
	if err := h.validatePutContainerResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putContainerResource",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PutExternal(value *HorizontalPodAutoscalerSpecMetricExternal) {
	if err := h.validatePutExternalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExternal",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PutObject(value *HorizontalPodAutoscalerSpecMetricObject) {
	if err := h.validatePutObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putObject",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PutPods(value *HorizontalPodAutoscalerSpecMetricPods) {
	if err := h.validatePutPodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putPods",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) PutResource(value *HorizontalPodAutoscalerSpecMetricResource) {
	if err := h.validatePutResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putResource",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResetContainerResource() {
	_jsii_.InvokeVoid(
		h,
		"resetContainerResource",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResetExternal() {
	_jsii_.InvokeVoid(
		h,
		"resetExternal",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResetObject() {
	_jsii_.InvokeVoid(
		h,
		"resetObject",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResetPods() {
	_jsii_.InvokeVoid(
		h,
		"resetPods",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		h,
		"resetResource",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

