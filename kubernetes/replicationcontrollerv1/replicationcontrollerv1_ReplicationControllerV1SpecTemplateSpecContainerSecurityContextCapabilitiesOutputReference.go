package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference interface {
	cdktf.ComplexObject
	Add() *[]*string
	SetAdd(val *[]*string)
	AddInput() *[]*string
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
	Drop() *[]*string
	SetDrop(val *[]*string)
	DropInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities
	SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities)
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
	ResetAdd()
	ResetDrop()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) Add() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"add",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) AddInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) Drop() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) DropInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetAdd(val *[]*string) {
	if err := j.validateSetAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"add",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetDrop(val *[]*string) {
	if err := j.validateSetDropParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drop",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ResetAdd() {
	_jsii_.InvokeVoid(
		r,
		"resetAdd",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ResetDrop() {
	_jsii_.InvokeVoid(
		r,
		"resetDrop",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

