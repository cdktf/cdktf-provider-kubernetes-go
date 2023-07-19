package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/pod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodSpecVolumeRbdOutputReference interface {
	cdktf.ComplexObject
	CephMonitors() *[]*string
	SetCephMonitors(val *[]*string)
	CephMonitorsInput() *[]*string
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
	FsType() *string
	SetFsType(val *string)
	FsTypeInput() *string
	InternalValue() *PodSpecVolumeRbd
	SetInternalValue(val *PodSpecVolumeRbd)
	Keyring() *string
	SetKeyring(val *string)
	KeyringInput() *string
	RadosUser() *string
	SetRadosUser(val *string)
	RadosUserInput() *string
	RbdImage() *string
	SetRbdImage(val *string)
	RbdImageInput() *string
	RbdPool() *string
	SetRbdPool(val *string)
	RbdPoolInput() *string
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SecretRef() PodSpecVolumeRbdSecretRefOutputReference
	SecretRefInput() *PodSpecVolumeRbdSecretRef
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
	PutSecretRef(value *PodSpecVolumeRbdSecretRef)
	ResetFsType()
	ResetKeyring()
	ResetRadosUser()
	ResetRbdPool()
	ResetReadOnly()
	ResetSecretRef()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodSpecVolumeRbdOutputReference
type jsiiProxy_PodSpecVolumeRbdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) CephMonitors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cephMonitors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) CephMonitorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cephMonitorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) InternalValue() *PodSpecVolumeRbd {
	var returns *PodSpecVolumeRbd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) Keyring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) KeyringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RadosUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radosUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RadosUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radosUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RbdImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RbdImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RbdPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) RbdPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) SecretRef() PodSpecVolumeRbdSecretRefOutputReference {
	var returns PodSpecVolumeRbdSecretRefOutputReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) SecretRefInput() *PodSpecVolumeRbdSecretRef {
	var returns *PodSpecVolumeRbdSecretRef
	_jsii_.Get(
		j,
		"secretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPodSpecVolumeRbdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSpecVolumeRbdOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecVolumeRbdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecVolumeRbdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeRbdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSpecVolumeRbdOutputReference_Override(p PodSpecVolumeRbdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeRbdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetCephMonitors(val *[]*string) {
	if err := j.validateSetCephMonitorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cephMonitors",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetInternalValue(val *PodSpecVolumeRbd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetKeyring(val *string) {
	if err := j.validateSetKeyringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetRadosUser(val *string) {
	if err := j.validateSetRadosUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radosUser",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetRbdImage(val *string) {
	if err := j.validateSetRbdImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rbdImage",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetRbdPool(val *string) {
	if err := j.validateSetRbdPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rbdPool",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeRbdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) PutSecretRef(value *PodSpecVolumeRbdSecretRef) {
	if err := p.validatePutSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		p,
		"resetFsType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetKeyring() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyring",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetRadosUser() {
	_jsii_.InvokeVoid(
		p,
		"resetRadosUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetRbdPool() {
	_jsii_.InvokeVoid(
		p,
		"resetRbdPool",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ResetSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PodSpecVolumeRbdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

