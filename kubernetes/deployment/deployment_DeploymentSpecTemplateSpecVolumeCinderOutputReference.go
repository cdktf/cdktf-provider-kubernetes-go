package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecTemplateSpecVolumeCinderOutputReference interface {
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
	FsType() *string
	SetFsType(val *string)
	FsTypeInput() *string
	InternalValue() *DeploymentSpecTemplateSpecVolumeCinder
	SetInternalValue(val *DeploymentSpecTemplateSpecVolumeCinder)
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
	VolumeId() *string
	SetVolumeId(val *string)
	VolumeIdInput() *string
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
	ResetFsType()
	ResetReadOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentSpecTemplateSpecVolumeCinderOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) InternalValue() *DeploymentSpecTemplateSpecVolumeCinder {
	var returns *DeploymentSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}


func NewDeploymentSpecTemplateSpecVolumeCinderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DeploymentSpecTemplateSpecVolumeCinderOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecVolumeCinderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeCinderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecVolumeCinderOutputReference_Override(d DeploymentSpecTemplateSpecVolumeCinderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeCinderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetInternalValue(val *DeploymentSpecTemplateSpecVolumeCinder) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference)SetVolumeId(val *string) {
	if err := j.validateSetVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		d,
		"resetFsType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeCinderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

