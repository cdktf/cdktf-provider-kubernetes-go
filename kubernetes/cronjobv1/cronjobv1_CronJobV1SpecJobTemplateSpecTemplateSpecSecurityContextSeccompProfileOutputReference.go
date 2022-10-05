package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference interface {
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
	InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile
	SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile)
	LocalhostProfile() *string
	SetLocalhostProfile(val *string)
	LocalhostProfileInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetLocalhostProfile()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) LocalhostProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localhostProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) LocalhostProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localhostProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetLocalhostProfile(val *string) {
	if err := j.validateSetLocalhostProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localhostProfile",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ResetLocalhostProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalhostProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

