package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference interface {
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
	ConfigMapKeyRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	FieldRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFrom
	SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFrom)
	ResourceFieldRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef
	SecretKeyRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef)
	PutResourceFieldRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef)
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

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) FieldRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) FieldRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFrom {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRef() ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRefInput() *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutConfigMapKeyRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromConfigMapKeyRef) {
	if err := r.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putConfigMapKeyRef",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutFieldRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef) {
	if err := r.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutResourceFieldRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef) {
	if err := r.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) PutSecretKeyRef(value *ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromSecretKeyRef) {
	if err := r.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecretKeyRef",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		r,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		r,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		r,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

