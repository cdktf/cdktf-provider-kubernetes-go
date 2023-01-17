package replicationcontroller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/replicationcontroller/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference interface {
	cdktf.ComplexObject
	ApiVersion() *string
	SetApiVersion(val *string)
	ApiVersionInput() *string
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
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRef
	SetInternalValue(val *ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRef)
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
	ResetApiVersion()
	ResetFieldPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference
type jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) InternalValue() *ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRef {
	var returns *ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRef
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference_Override(r ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetApiVersion(val *string) {
	if err := j.validateSetApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetInternalValue(val *ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRef) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ResetApiVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetApiVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		r,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeDownwardApiItemsFieldRefOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

