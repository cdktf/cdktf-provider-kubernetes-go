package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference interface {
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
	FsGroup() *string
	SetFsGroup(val *string)
	FsGroupInput() *string
	InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext
	SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext)
	RunAsGroup() *string
	SetRunAsGroup(val *string)
	RunAsGroupInput() *string
	RunAsNonRoot() interface{}
	SetRunAsNonRoot(val interface{})
	RunAsNonRootInput() interface{}
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	SeccompProfile() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference
	SeccompProfileInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile
	SeLinuxOptions() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference
	SeLinuxOptionsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptions
	SupplementalGroups() *[]*float64
	SetSupplementalGroups(val *[]*float64)
	SupplementalGroupsInput() *[]*float64
	Sysctl() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSysctlList
	SysctlInput() interface{}
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
	PutSeccompProfile(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile)
	PutSeLinuxOptions(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptions)
	PutSysctl(value interface{})
	ResetFsGroup()
	ResetRunAsGroup()
	ResetRunAsNonRoot()
	ResetRunAsUser()
	ResetSeccompProfile()
	ResetSeLinuxOptions()
	ResetSupplementalGroups()
	ResetSysctl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) FsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) FsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SeccompProfile() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfileOutputReference
	_jsii_.Get(
		j,
		"seccompProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SeccompProfileInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile
	_jsii_.Get(
		j,
		"seccompProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SeLinuxOptions() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptionsOutputReference
	_jsii_.Get(
		j,
		"seLinuxOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SeLinuxOptionsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptions {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptions
	_jsii_.Get(
		j,
		"seLinuxOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SupplementalGroups() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SupplementalGroupsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"supplementalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) Sysctl() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSysctlList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSysctlList
	_jsii_.Get(
		j,
		"sysctl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) SysctlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sysctlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetFsGroup(val *string) {
	if err := j.validateSetFsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsGroup",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetRunAsGroup(val *string) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetSupplementalGroups(val *[]*float64) {
	if err := j.validateSetSupplementalGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplementalGroups",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) PutSeccompProfile(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeccompProfile) {
	if err := c.validatePutSeccompProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSeccompProfile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) PutSeLinuxOptions(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextSeLinuxOptions) {
	if err := c.validatePutSeLinuxOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSeLinuxOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) PutSysctl(value interface{}) {
	if err := c.validatePutSysctlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSysctl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetFsGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetFsGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		c,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		c,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetSeccompProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetSeccompProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetSeLinuxOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetSeLinuxOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetSupplementalGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetSupplementalGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ResetSysctl() {
	_jsii_.InvokeVoid(
		c,
		"resetSysctl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

