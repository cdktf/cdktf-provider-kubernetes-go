// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/podv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference interface {
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
	InternalValue() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec
	SetInternalValue(val *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec)
	Resources() PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference
	ResourcesInput() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources
	Selector() PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference
	SelectorInput() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutResources(value *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources)
	PutSelector(value *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector)
	ResetSelector()
	ResetStorageClassName()
	ResetVolumeMode()
	ResetVolumeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference
type jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) AccessModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) AccessModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InternalValue() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec {
	var returns *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Resources() PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference {
	var returns PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResourcesInput() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources {
	var returns *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Selector() PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference {
	var returns PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) SelectorInput() *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector {
	var returns *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) StorageClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) VolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeNameInput",
		&returns,
	)
	return returns
}


func NewPodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference_Override(p PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetAccessModes(val *[]*string) {
	if err := j.validateSetAccessModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessModes",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetInternalValue(val *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetStorageClassName(val *string) {
	if err := j.validateSetStorageClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClassName",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetVolumeMode(val *string) {
	if err := j.validateSetVolumeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeMode",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference)SetVolumeName(val *string) {
	if err := j.validateSetVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeName",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) PutResources(value *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecResources) {
	if err := p.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putResources",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) PutSelector(value *PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecSelector) {
	if err := p.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSelector",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetSelector() {
	_jsii_.InvokeVoid(
		p,
		"resetSelector",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetStorageClassName() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageClassName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetVolumeMode() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ResetVolumeName() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

