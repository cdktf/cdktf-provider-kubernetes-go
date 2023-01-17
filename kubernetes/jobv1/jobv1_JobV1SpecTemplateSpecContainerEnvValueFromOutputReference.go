package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecContainerEnvValueFromOutputReference interface {
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
	ConfigMapKeyRef() JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() JobV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	FieldRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *JobV1SpecTemplateSpecContainerEnvValueFrom
	SetInternalValue(val *JobV1SpecTemplateSpecContainerEnvValueFrom)
	ResourceFieldRef() JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef
	SecretKeyRef() JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *JobV1SpecTemplateSpecContainerEnvValueFromFieldRef)
	PutResourceFieldRef(value *JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef)
	ResetConfigMapKeyRef()
	ResetFieldRef()
	ResetResourceFieldRef()
	ResetSecretKeyRef()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecContainerEnvValueFromOutputReference
type jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRef() JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference {
	var returns JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef {
	var returns *JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) FieldRef() JobV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference {
	var returns JobV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) FieldRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromFieldRef {
	var returns *JobV1SpecTemplateSpecContainerEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) InternalValue() *JobV1SpecTemplateSpecContainerEnvValueFrom {
	var returns *JobV1SpecTemplateSpecContainerEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRef() JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference {
	var returns JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef {
	var returns *JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRef() JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference {
	var returns JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRefInput() *JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef {
	var returns *JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecContainerEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobV1SpecTemplateSpecContainerEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecContainerEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecContainerEnvValueFromOutputReference_Override(j JobV1SpecTemplateSpecContainerEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetInternalValue(val *JobV1SpecTemplateSpecContainerEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutConfigMapKeyRef(value *JobV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef) {
	if err := j.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfigMapKeyRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutFieldRef(value *JobV1SpecTemplateSpecContainerEnvValueFromFieldRef) {
	if err := j.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutResourceFieldRef(value *JobV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef) {
	if err := j.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutSecretKeyRef(value *JobV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef) {
	if err := j.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecretKeyRef",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		j,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		j,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		j,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		j,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

