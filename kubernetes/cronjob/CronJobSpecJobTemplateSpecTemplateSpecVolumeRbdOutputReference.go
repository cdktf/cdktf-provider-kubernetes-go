// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/cronjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference interface {
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
	InternalValue() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd
	SetInternalValue(val *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd)
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
	SecretRef() CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRefOutputReference
	SecretRefInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef
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
	PutSecretRef(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef)
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

// The jsii proxy struct for CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference
type jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) CephMonitors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cephMonitors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) CephMonitorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cephMonitorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) FsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) FsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) InternalValue() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) Keyring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) KeyringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RadosUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radosUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RadosUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radosUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RbdImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RbdImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RbdPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) RbdPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rbdPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) SecretRef() CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRefOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRefOutputReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) SecretRefInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef
	_jsii_.Get(
		j,
		"secretRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference_Override(c CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetCephMonitors(val *[]*string) {
	if err := j.validateSetCephMonitorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cephMonitors",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetFsType(val *string) {
	if err := j.validateSetFsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsType",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetInternalValue(val *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetKeyring(val *string) {
	if err := j.validateSetKeyringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetRadosUser(val *string) {
	if err := j.validateSetRadosUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radosUser",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetRbdImage(val *string) {
	if err := j.validateSetRbdImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rbdImage",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetRbdPool(val *string) {
	if err := j.validateSetRbdPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rbdPool",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) PutSecretRef(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef) {
	if err := c.validatePutSecretRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretRef",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetFsType() {
	_jsii_.InvokeVoid(
		c,
		"resetFsType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetKeyring() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyring",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetRadosUser() {
	_jsii_.InvokeVoid(
		c,
		"resetRadosUser",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetRbdPool() {
	_jsii_.InvokeVoid(
		c,
		"resetRbdPool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		c,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ResetSecretRef() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretRef",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

