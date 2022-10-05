package daemonset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/daemonset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference interface {
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
	FieldRef() DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference
	FieldRefInput() *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	ResourceFieldRef() DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference
	ResourceFieldRefInput() *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef
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
	PutFieldRef(value *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef)
	PutResourceFieldRef(value *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef)
	ResetFieldRef()
	ResetMode()
	ResetResourceFieldRef()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference
type jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) FieldRef() DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) FieldRefInput() *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef {
	var returns *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ResourceFieldRef() DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ResourceFieldRefInput() *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef {
	var returns *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference_Override(d DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) PutFieldRef(value *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef) {
	if err := d.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) PutResourceFieldRef(value *DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef) {
	if err := d.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResourceFieldRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

