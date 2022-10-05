package persistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/persistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference interface {
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
	InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile)
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
	SecretNamespace() *string
	SetSecretNamespace(val *string)
	SecretNamespaceInput() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
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
	ResetReadOnly()
	ResetSecretNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference
type jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) InternalValue() *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) SecretNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) SecretNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference_Override(p PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetSecretName(val *string) {
	if err := j.validateSetSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetSecretNamespace(val *string) {
	if err := j.validateSetSecretNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretNamespace",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ResetSecretNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetSecretNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

