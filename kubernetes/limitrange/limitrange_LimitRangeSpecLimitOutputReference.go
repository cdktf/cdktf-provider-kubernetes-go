package limitrange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/limitrange/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LimitRangeSpecLimitOutputReference interface {
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
	Default() *map[string]*string
	SetDefault(val *map[string]*string)
	DefaultInput() *map[string]*string
	DefaultRequest() *map[string]*string
	SetDefaultRequest(val *map[string]*string)
	DefaultRequestInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Max() *map[string]*string
	SetMax(val *map[string]*string)
	MaxInput() *map[string]*string
	MaxLimitRequestRatio() *map[string]*string
	SetMaxLimitRequestRatio(val *map[string]*string)
	MaxLimitRequestRatioInput() *map[string]*string
	Min() *map[string]*string
	SetMin(val *map[string]*string)
	MinInput() *map[string]*string
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
	ResetDefault()
	ResetDefaultRequest()
	ResetMax()
	ResetMaxLimitRequestRatio()
	ResetMin()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LimitRangeSpecLimitOutputReference
type jsiiProxy_LimitRangeSpecLimitOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) Default() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) DefaultInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) DefaultRequest() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) DefaultRequestInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) Max() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) MaxInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) MaxLimitRequestRatio() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"maxLimitRequestRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) MaxLimitRequestRatioInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"maxLimitRequestRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) Min() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) MinInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewLimitRangeSpecLimitOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LimitRangeSpecLimitOutputReference {
	_init_.Initialize()

	if err := validateNewLimitRangeSpecLimitOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LimitRangeSpecLimitOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.limitRange.LimitRangeSpecLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLimitRangeSpecLimitOutputReference_Override(l LimitRangeSpecLimitOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.limitRange.LimitRangeSpecLimitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetDefault(val *map[string]*string) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetDefaultRequest(val *map[string]*string) {
	if err := j.validateSetDefaultRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRequest",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetMax(val *map[string]*string) {
	if err := j.validateSetMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"max",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetMaxLimitRequestRatio(val *map[string]*string) {
	if err := j.validateSetMaxLimitRequestRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLimitRequestRatio",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetMin(val *map[string]*string) {
	if err := j.validateSetMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"min",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LimitRangeSpecLimitOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		l,
		"resetDefault",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetDefaultRequest() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultRequest",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		l,
		"resetMax",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetMaxLimitRequestRatio() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxLimitRequestRatio",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		l,
		"resetMin",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		l,
		"resetType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LimitRangeSpecLimitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

