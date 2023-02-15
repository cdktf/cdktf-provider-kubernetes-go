package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/cronjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Items() CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretItemsList
	ItemsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Optional() interface{}
	SetOptional(val interface{})
	OptionalInput() interface{}
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
	PutItems(value interface{})
	ResetItems()
	ResetName()
	ResetOptional()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference
type jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) Items() CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretItemsList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretItemsList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) Optional() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) OptionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference_Override(c CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetOptional(val interface{}) {
	if err := j.validateSetOptionalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optional",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) PutItems(value interface{}) {
	if err := c.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putItems",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ResetOptional() {
	_jsii_.InvokeVoid(
		c,
		"resetOptional",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecretOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

