package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference interface {
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
	InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity
	SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity)
	NodeAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityOutputReference
	NodeAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity
	PodAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinityOutputReference
	PodAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity
	PodAntiAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinityOutputReference
	PodAntiAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity
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
	PutNodeAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity)
	PutPodAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity)
	PutPodAntiAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity)
	ResetNodeAffinity()
	ResetPodAffinity()
	ResetPodAntiAffinity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) NodeAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityOutputReference
	_jsii_.Get(
		j,
		"nodeAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) NodeAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity
	_jsii_.Get(
		j,
		"nodeAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PodAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinityOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinityOutputReference
	_jsii_.Get(
		j,
		"podAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PodAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity
	_jsii_.Get(
		j,
		"podAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PodAntiAffinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinityOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinityOutputReference
	_jsii_.Get(
		j,
		"podAntiAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PodAntiAffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity
	_jsii_.Get(
		j,
		"podAntiAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference)SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PutNodeAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity) {
	if err := c.validatePutNodeAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeAffinity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PutPodAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity) {
	if err := c.validatePutPodAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPodAffinity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) PutPodAntiAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity) {
	if err := c.validatePutPodAntiAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPodAntiAffinity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ResetNodeAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ResetPodAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetPodAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ResetPodAntiAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetPodAntiAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

