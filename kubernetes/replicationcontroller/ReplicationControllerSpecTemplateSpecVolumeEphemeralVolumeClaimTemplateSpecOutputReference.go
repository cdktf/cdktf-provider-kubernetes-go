// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/replicationcontroller/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	AccessModes() *[]*string
	SetAccessModes(val *[]*string)
	AccessModesInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec
	SetInternalValue(val *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec)
	Resources() ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference
	ResourcesInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources
	Selector() ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference
	SelectorInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector
	StorageClassName() *string
	SetStorageClassName(val *string)
	StorageClassNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMode() *string
	SetVolumeMode(val *string)
	VolumeModeInput() *string
	VolumeName() *string
	SetVolumeName(val *string)
	VolumeNameInput() *string
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
	PutResources(value *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources)
	PutSelector(value *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector)
	ResetSelector()
	ResetStorageClassName()
	ResetVolumeMode()
	ResetVolumeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference
type jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) AccessModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) AccessModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InternalValue() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec {
	var returns *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Resources() ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResourcesInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources {
	var returns *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Selector() ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) SelectorInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector {
	var returns *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) StorageClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeNameInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference_Override(r ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetAccessModes(val *[]*string) {
	if err := j.validateSetAccessModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessModes",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetInternalValue(val *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetStorageClassName(val *string) {
	if err := j.validateSetStorageClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClassName",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetVolumeMode(val *string) {
	if err := j.validateSetVolumeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeMode",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetVolumeName(val *string) {
	if err := j.validateSetVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeName",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) PutResources(value *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources) {
	if err := r.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResources",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) PutSelector(value *ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector) {
	if err := r.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSelector",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		r,
		"resetSelector",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetStorageClassName() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageClassName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetVolumeMode() {
	_jsii_.InvokeVoid(
		r,
		"resetVolumeMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetVolumeName() {
	_jsii_.InvokeVoid(
		r,
		"resetVolumeName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

