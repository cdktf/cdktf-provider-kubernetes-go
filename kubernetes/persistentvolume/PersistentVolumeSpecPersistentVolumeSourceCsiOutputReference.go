// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/persistentvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference interface {
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
	ControllerExpandSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	ControllerExpandSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef
	ControllerPublishSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	ControllerPublishSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef
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
	InternalValue() *PersistentVolumeSpecPersistentVolumeSourceCsi
	SetInternalValue(val *PersistentVolumeSpecPersistentVolumeSourceCsi)
	NodePublishSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	NodePublishSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef
	NodeStageSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	NodeStageSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef
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
	PutControllerExpandSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef)
	PutControllerPublishSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef)
	PutNodePublishSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef)
	PutNodeStageSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef)
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

// The jsii proxy struct for PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference
type jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerExpandSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ControllerExpandSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef
	_jsii_.Get(
		j,
		"controllerExpandSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"controllerPublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ControllerPublishSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef
	_jsii_.Get(
		j,
		"controllerPublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) InternalValue() *PersistentVolumeSpecPersistentVolumeSourceCsi {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodePublishSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) NodePublishSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef
	_jsii_.Get(
		j,
		"nodePublishSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRef() PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRefOutputReference
	_jsii_.Get(
		j,
		"nodeStageSecretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) NodeStageSecretRefInput() *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef
	_jsii_.Get(
		j,
		"nodeStageSecretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) VolumeAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) VolumeAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) VolumeHandle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) VolumeHandleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeHandleInput",
		&returns,
	)
	return returns
}


func NewPersistentVolumeSpecPersistentVolumeSourceCsiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeSpecPersistentVolumeSourceCsiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeSpecPersistentVolumeSourceCsiOutputReference_Override(p PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetInternalValue(val *PersistentVolumeSpecPersistentVolumeSourceCsi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetVolumeAttributes(val *map[string]*string) {
	if err := j.validateSetVolumeAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAttributes",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference)SetVolumeHandle(val *string) {
	if err := j.validateSetVolumeHandleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeHandle",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) PutControllerExpandSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef) {
	if err := p.validatePutControllerExpandSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putControllerExpandSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) PutControllerPublishSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef) {
	if err := p.validatePutControllerPublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putControllerPublishSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) PutNodePublishSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef) {
	if err := p.validatePutNodePublishSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodePublishSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) PutNodeStageSecretRef(value *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef) {
	if err := p.validatePutNodeStageSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodeStageSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetControllerExpandSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetControllerExpandSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetControllerPublishSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetControllerPublishSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		p,
		"resetFsType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetNodePublishSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetNodePublishSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetNodeStageSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeStageSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ResetVolumeAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

