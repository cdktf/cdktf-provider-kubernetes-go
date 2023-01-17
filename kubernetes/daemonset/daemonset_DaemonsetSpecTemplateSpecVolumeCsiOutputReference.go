package daemonset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/daemonset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaemonsetSpecTemplateSpecVolumeCsiOutputReference interface {
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
	InternalValue() *DaemonsetSpecTemplateSpecVolumeCsi
	SetInternalValue(val *DaemonsetSpecTemplateSpecVolumeCsi)
	NodePublishSecretRef() DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference
	NodePublishSecretRefInput() *DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRef
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
	PutNodePublishSecretRef(value *DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRef)
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

// The jsii proxy struct for DaemonsetSpecTemplateSpecVolumeCsiOutputReference
type jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) InternalValue() *DaemonsetSpecTemplateSpecVolumeCsi {
	var returns *DaemonsetSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) NodePublishSecretRef() DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodePublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) NodePublishSecretRefInput() *DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRef {
	var returns *DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRef
	_jsii_.Get(
		j,
		"nodePublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) VolumeAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) VolumeAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributesInput",
		&returns,
	)
	return returns
}


func NewDaemonsetSpecTemplateSpecVolumeCsiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DaemonsetSpecTemplateSpecVolumeCsiOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonsetSpecTemplateSpecVolumeCsiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDaemonsetSpecTemplateSpecVolumeCsiOutputReference_Override(d DaemonsetSpecTemplateSpecVolumeCsiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetInternalValue(val *DaemonsetSpecTemplateSpecVolumeCsi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference)SetVolumeAttributes(val *map[string]*string) {
	if err := j.validateSetVolumeAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAttributes",
		val,
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) PutNodePublishSecretRef(value *DaemonsetSpecTemplateSpecVolumeCsiNodePublishSecretRef) {
	if err := d.validatePutNodePublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNodePublishSecretRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		d,
		"resetFsType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ResetNodePublishSecretRef() {
	_jsii_.InvokeVoid(
		d,
		"resetNodePublishSecretRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ResetVolumeAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeCsiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

