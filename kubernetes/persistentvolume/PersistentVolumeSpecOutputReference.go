package persistentvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/persistentvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeSpecOutputReference interface {
	cdktf.ComplexObject
	AccessModes() *[]*string
	SetAccessModes(val *[]*string)
	AccessModesInput() *[]*string
	Capacity() *map[string]*string
	SetCapacity(val *map[string]*string)
	CapacityInput() *map[string]*string
	ClaimRef() PersistentVolumeSpecClaimRefOutputReference
	ClaimRefInput() *PersistentVolumeSpecClaimRef
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountOptions() *[]*string
	SetMountOptions(val *[]*string)
	MountOptionsInput() *[]*string
	NodeAffinity() PersistentVolumeSpecNodeAffinityOutputReference
	NodeAffinityInput() *PersistentVolumeSpecNodeAffinity
	PersistentVolumeReclaimPolicy() *string
	SetPersistentVolumeReclaimPolicy(val *string)
	PersistentVolumeReclaimPolicyInput() *string
	PersistentVolumeSource() PersistentVolumeSpecPersistentVolumeSourceOutputReference
	PersistentVolumeSourceInput() *PersistentVolumeSpecPersistentVolumeSource
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
	PutClaimRef(value *PersistentVolumeSpecClaimRef)
	PutNodeAffinity(value *PersistentVolumeSpecNodeAffinity)
	PutPersistentVolumeSource(value *PersistentVolumeSpecPersistentVolumeSource)
	ResetClaimRef()
	ResetMountOptions()
	ResetNodeAffinity()
	ResetPersistentVolumeReclaimPolicy()
	ResetStorageClassName()
	ResetVolumeMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeSpecOutputReference
type jsiiProxy_PersistentVolumeSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) AccessModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) AccessModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) Capacity() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) CapacityInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) ClaimRef() PersistentVolumeSpecClaimRefOutputReference {
	var returns PersistentVolumeSpecClaimRefOutputReference
	_jsii_.Get(
		j,
		"claimRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) ClaimRefInput() *PersistentVolumeSpecClaimRef {
	var returns *PersistentVolumeSpecClaimRef
	_jsii_.Get(
		j,
		"claimRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) MountOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) MountOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) NodeAffinity() PersistentVolumeSpecNodeAffinityOutputReference {
	var returns PersistentVolumeSpecNodeAffinityOutputReference
	_jsii_.Get(
		j,
		"nodeAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) NodeAffinityInput() *PersistentVolumeSpecNodeAffinity {
	var returns *PersistentVolumeSpecNodeAffinity
	_jsii_.Get(
		j,
		"nodeAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) PersistentVolumeReclaimPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentVolumeReclaimPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) PersistentVolumeReclaimPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentVolumeReclaimPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) PersistentVolumeSource() PersistentVolumeSpecPersistentVolumeSourceOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) PersistentVolumeSourceInput() *PersistentVolumeSpecPersistentVolumeSource {
	var returns *PersistentVolumeSpecPersistentVolumeSource
	_jsii_.Get(
		j,
		"persistentVolumeSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) StorageClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) StorageClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) VolumeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference) VolumeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeModeInput",
		&returns,
	)
	return returns
}


func NewPersistentVolumeSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PersistentVolumeSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeSpecOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPersistentVolumeSpecOutputReference_Override(p PersistentVolumeSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetAccessModes(val *[]*string) {
	if err := j.validateSetAccessModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessModes",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetCapacity(val *map[string]*string) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetMountOptions(val *[]*string) {
	if err := j.validateSetMountOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountOptions",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetPersistentVolumeReclaimPolicy(val *string) {
	if err := j.validateSetPersistentVolumeReclaimPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentVolumeReclaimPolicy",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetStorageClassName(val *string) {
	if err := j.validateSetStorageClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClassName",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecOutputReference)SetVolumeMode(val *string) {
	if err := j.validateSetVolumeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeMode",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) PutClaimRef(value *PersistentVolumeSpecClaimRef) {
	if err := p.validatePutClaimRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putClaimRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) PutNodeAffinity(value *PersistentVolumeSpecNodeAffinity) {
	if err := p.validatePutNodeAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNodeAffinity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) PutPersistentVolumeSource(value *PersistentVolumeSpecPersistentVolumeSource) {
	if err := p.validatePutPersistentVolumeSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPersistentVolumeSource",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetClaimRef() {
	_jsii_.InvokeVoid(
		p,
		"resetClaimRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetMountOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetMountOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetNodeAffinity() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeAffinity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetPersistentVolumeReclaimPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetPersistentVolumeReclaimPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetStorageClassName() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageClassName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ResetVolumeMode() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersistentVolumeSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

