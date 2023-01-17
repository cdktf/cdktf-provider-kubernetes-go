package persistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/persistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceCephFs
	SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceCephFs)
	Monitors() *[]*string
	SetMonitors(val *[]*string)
	MonitorsInput() *[]*string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SecretFile() *string
	SetSecretFile(val *string)
	SecretFileInput() *string
	SecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRefOutputReference
	SecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef)
	ResetPath()
	ResetReadOnly()
	ResetSecretFile()
	ResetSecretRef()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference
type jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceCephFs {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCephFs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) Monitors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) MonitorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) SecretFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) SecretFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) SecretRef() PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRefOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRefOutputReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) SecretRefInput() *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef
	_jsii_.Get(
		j,
		"secretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference_Override(p PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceCephFs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetMonitors(val *[]*string) {
	if err := j.validateSetMonitorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitors",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetSecretFile(val *string) {
	if err := j.validateSetSecretFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretFile",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) PutSecretRef(value *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef) {
	if err := p.validatePutSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecretRef",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		p,
		"resetPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ResetSecretFile() {
	_jsii_.InvokeVoid(
		p,
		"resetSecretFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ResetSecretRef() {
	_jsii_.InvokeVoid(
		p,
		"resetSecretRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		p,
		"resetUser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

