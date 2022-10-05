package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecContainerVolumeMountOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountPath() *string
	SetMountPath(val *string)
	MountPathInput() *string
	MountPropagation() *string
	SetMountPropagation(val *string)
	MountPropagationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SubPath() *string
	SetSubPath(val *string)
	SubPathInput() *string
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
	ResetMountPropagation()
	ResetReadOnly()
	ResetSubPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecContainerVolumeMountOutputReference
type jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) MountPropagation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) MountPropagationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) SubPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) SubPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecContainerVolumeMountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobV1SpecTemplateSpecContainerVolumeMountOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecContainerVolumeMountOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerVolumeMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecContainerVolumeMountOutputReference_Override(j JobV1SpecTemplateSpecContainerVolumeMountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerVolumeMountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetMountPath(val *string) {
	if err := j.validateSetMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetMountPropagation(val *string) {
	if err := j.validateSetMountPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPropagation",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetSubPath(val *string) {
	if err := j.validateSetSubPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subPath",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ResetMountPropagation() {
	_jsii_.InvokeVoid(
		j,
		"resetMountPropagation",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		j,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ResetSubPath() {
	_jsii_.InvokeVoid(
		j,
		"resetSubPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerVolumeMountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

