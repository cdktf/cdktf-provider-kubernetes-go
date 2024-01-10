// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/persistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference interface {
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
	ControllerExpandSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	ControllerExpandSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef
	ControllerPublishSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	ControllerPublishSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef
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
	InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceCsi
	SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceCsi)
	NodePublishSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	NodePublishSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef
	NodeStageSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	NodeStageSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef
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
	PutControllerExpandSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef)
	PutControllerPublishSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef)
	PutNodePublishSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef)
	PutNodeStageSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef)
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

// The jsii proxy struct for PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
type jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerExpandSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef
	_jsii_.Get(
		j,
		"controllerExpandSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerPublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef
	_jsii_.Get(
		j,
		"controllerPublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceCsi {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodePublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef
	_jsii_.Get(
		j,
		"nodePublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodeStageSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef
	_jsii_.Get(
		j,
		"nodeStageSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeHandle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) VolumeHandleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandleInput",
		&returns,
	)
	return returns
}


func NewPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference_Override(p PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceCsi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetVolumeAttributes(val *map[string]*string) {
	if err := j.validateSetVolumeAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAttributes",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference)SetVolumeHandle(val *string) {
	if err := j.validateSetVolumeHandleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeHandle",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutControllerExpandSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef) {
	if err := p.validatePutControllerExpandSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putControllerExpandSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutControllerPublishSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef) {
	if err := p.validatePutControllerPublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putControllerPublishSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutNodePublishSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef) {
	if err := p.validatePutNodePublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodePublishSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) PutNodeStageSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef) {
	if err := p.validatePutNodeStageSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodeStageSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetControllerExpandSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetControllerExpandSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetControllerPublishSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetControllerPublishSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		p,
		"resetFsType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetNodePublishSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetNodePublishSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetNodeStageSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeStageSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ResetVolumeAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

