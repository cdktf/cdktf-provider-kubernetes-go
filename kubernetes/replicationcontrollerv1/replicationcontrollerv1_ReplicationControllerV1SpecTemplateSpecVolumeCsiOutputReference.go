package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference interface {
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
	Driver() *string
	SetDriver(val *string)
	DriverInput() *string
	// Experimental.
	Fqn() *string
	FsType() *string
	SetFsType(val *string)
	FsTypeInput() *string
	InternalValue() *ReplicationControllerV1SpecTemplateSpecVolumeCsi
	SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecVolumeCsi)
	NodePublishSecretRef() ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference
	NodePublishSecretRefInput() *ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRef
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeAttributes() *map[string]*string
	SetVolumeAttributes(val *map[string]*string)
	VolumeAttributesInput() *map[string]*string
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
	PutNodePublishSecretRef(value *ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRef)
	ResetFsType()
	ResetNodePublishSecretRef()
	ResetReadOnly()
	ResetVolumeAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) InternalValue() *ReplicationControllerV1SpecTemplateSpecVolumeCsi {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) NodePublishSecretRef() ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodePublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) NodePublishSecretRefInput() *ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRef {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRef
	_jsii_.Get(
		j,
		"nodePublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) VolumeAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) VolumeAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributesInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecVolumeCsi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference)SetVolumeAttributes(val *map[string]*string) {
	if err := j.validateSetVolumeAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAttributes",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) PutNodePublishSecretRef(value *ReplicationControllerV1SpecTemplateSpecVolumeCsiNodePublishSecretRef) {
	if err := r.validatePutNodePublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNodePublishSecretRef",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		r,
		"resetFsType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ResetNodePublishSecretRef() {
	_jsii_.InvokeVoid(
		r,
		"resetNodePublishSecretRef",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		r,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ResetVolumeAttributes() {
	_jsii_.InvokeVoid(
		r,
		"resetVolumeAttributes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

