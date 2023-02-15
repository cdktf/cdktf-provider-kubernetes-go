package env

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/env/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnvEnvValueFromOutputReference interface {
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
	ConfigMapKeyRef() EnvEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *EnvEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() EnvEnvValueFromFieldRefOutputReference
	FieldRefInput() *EnvEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *EnvEnvValueFrom
	SetInternalValue(val *EnvEnvValueFrom)
	ResourceFieldRef() EnvEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *EnvEnvValueFromResourceFieldRef
	SecretKeyRef() EnvEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *EnvEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *EnvEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *EnvEnvValueFromFieldRef)
	PutResourceFieldRef(value *EnvEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *EnvEnvValueFromSecretKeyRef)
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

// The jsii proxy struct for EnvEnvValueFromOutputReference
type jsiiProxy_EnvEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ConfigMapKeyRef() EnvEnvValueFromConfigMapKeyRefOutputReference {
	var returns EnvEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ConfigMapKeyRefInput() *EnvEnvValueFromConfigMapKeyRef {
	var returns *EnvEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) FieldRef() EnvEnvValueFromFieldRefOutputReference {
	var returns EnvEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) FieldRefInput() *EnvEnvValueFromFieldRef {
	var returns *EnvEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) InternalValue() *EnvEnvValueFrom {
	var returns *EnvEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ResourceFieldRef() EnvEnvValueFromResourceFieldRefOutputReference {
	var returns EnvEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) ResourceFieldRefInput() *EnvEnvValueFromResourceFieldRef {
	var returns *EnvEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) SecretKeyRef() EnvEnvValueFromSecretKeyRefOutputReference {
	var returns EnvEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) SecretKeyRefInput() *EnvEnvValueFromSecretKeyRef {
	var returns *EnvEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEnvEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnvEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewEnvEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.env.EnvEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnvEnvValueFromOutputReference_Override(e EnvEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.env.EnvEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference)SetInternalValue(val *EnvEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnvEnvValueFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) PutConfigMapKeyRef(value *EnvEnvValueFromConfigMapKeyRef) {
	if err := e.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConfigMapKeyRef",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) PutFieldRef(value *EnvEnvValueFromFieldRef) {
	if err := e.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) PutResourceFieldRef(value *EnvEnvValueFromResourceFieldRef) {
	if err := e.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) PutSecretKeyRef(value *EnvEnvValueFromSecretKeyRef) {
	if err := e.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecretKeyRef",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		e,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

