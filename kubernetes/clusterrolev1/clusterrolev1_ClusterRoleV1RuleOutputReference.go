package clusterrolev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/clusterrolev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterRoleV1RuleOutputReference interface {
	cdktf.ComplexObject
	ApiGroups() *[]*string
	SetApiGroups(val *[]*string)
	ApiGroupsInput() *[]*string
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
	NonResourceUrls() *[]*string
	SetNonResourceUrls(val *[]*string)
	NonResourceUrlsInput() *[]*string
	ResourceNames() *[]*string
	SetResourceNames(val *[]*string)
	ResourceNamesInput() *[]*string
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Verbs() *[]*string
	SetVerbs(val *[]*string)
	VerbsInput() *[]*string
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
	ResetApiGroups()
	ResetNonResourceUrls()
	ResetResourceNames()
	ResetResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterRoleV1RuleOutputReference
type jsiiProxy_ClusterRoleV1RuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ApiGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ApiGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) NonResourceUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonResourceUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) NonResourceUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonResourceUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ResourceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ResourceNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) Verbs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"verbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference) VerbsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"verbsInput",
		&returns,
	)
	return returns
}


func NewClusterRoleV1RuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClusterRoleV1RuleOutputReference {
	_init_.Initialize()

	if err := validateNewClusterRoleV1RuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterRoleV1RuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.clusterRoleV1.ClusterRoleV1RuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClusterRoleV1RuleOutputReference_Override(c ClusterRoleV1RuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.clusterRoleV1.ClusterRoleV1RuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetApiGroups(val *[]*string) {
	if err := j.validateSetApiGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiGroups",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetNonResourceUrls(val *[]*string) {
	if err := j.validateSetNonResourceUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonResourceUrls",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetResourceNames(val *[]*string) {
	if err := j.validateSetResourceNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceNames",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClusterRoleV1RuleOutputReference)SetVerbs(val *[]*string) {
	if err := j.validateSetVerbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verbs",
		val,
	)
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ResetApiGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetApiGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ResetNonResourceUrls() {
	_jsii_.InvokeVoid(
		c,
		"resetNonResourceUrls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ResetResourceNames() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterRoleV1RuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

