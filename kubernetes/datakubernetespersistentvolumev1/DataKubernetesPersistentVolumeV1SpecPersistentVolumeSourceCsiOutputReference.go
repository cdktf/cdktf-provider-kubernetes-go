// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/datakubernetespersistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference interface {
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
	ControllerExpandSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	ControllerExpandSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef
	ControllerPublishSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	ControllerPublishSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef
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
	InternalValue() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi
	SetInternalValue(val *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi)
	NodePublishSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	NodePublishSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef
	NodeStageSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	NodeStageSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef
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
	VolumeHandle() *string
	SetVolumeHandle(val *string)
	VolumeHandleInput() *string
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
	PutControllerExpandSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef)
	PutControllerPublishSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef)
	PutNodePublishSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef)
	PutNodeStageSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef)
	ResetControllerExpandSecretRef()
	ResetControllerPublishSecretRef()
	ResetFsType()
	ResetNodePublishSecretRef()
	ResetNodeStageSecretRef()
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

// The jsii proxy struct for DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
type jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerExpandSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef
	_jsii_.Get(
		j,
		"controllerExpandSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerPublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef
	_jsii_.Get(
		j,
		"controllerPublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InternalValue() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodePublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef
	_jsii_.Get(
		j,
		"nodePublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRef() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodeStageSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRefInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef
	_jsii_.Get(
		j,
		"nodeStageSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeHandle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeHandleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandleInput",
		&returns,
	)
	return returns
}


func NewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeV1.DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference_Override(d DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeV1.DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetInternalValue(val *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetVolumeAttributes(val *map[string]*string) {
	if err := j.validateSetVolumeAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAttributes",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetVolumeHandle(val *string) {
	if err := j.validateSetVolumeHandleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeHandle",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutControllerExpandSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef) {
	if err := d.validatePutControllerExpandSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putControllerExpandSecretRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutControllerPublishSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef) {
	if err := d.validatePutControllerPublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putControllerPublishSecretRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutNodePublishSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef) {
	if err := d.validatePutNodePublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNodePublishSecretRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutNodeStageSecretRef(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef) {
	if err := d.validatePutNodeStageSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNodeStageSecretRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetControllerExpandSecretRef() {
	_jsii_.InvokeVoid(
		d,
		"resetControllerExpandSecretRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetControllerPublishSecretRef() {
	_jsii_.InvokeVoid(
		d,
		"resetControllerPublishSecretRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		d,
		"resetFsType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetNodePublishSecretRef() {
	_jsii_.InvokeVoid(
		d,
		"resetNodePublishSecretRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetNodeStageSecretRef() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeStageSecretRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetVolumeAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

