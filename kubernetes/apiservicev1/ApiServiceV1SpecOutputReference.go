package apiservicev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/apiservicev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiServiceV1SpecOutputReference interface {
	cdktf.ComplexObject
	CaBundle() *string
	SetCaBundle(val *string)
	CaBundleInput() *string
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
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	GroupPriorityMinimum() *float64
	SetGroupPriorityMinimum(val *float64)
	GroupPriorityMinimumInput() *float64
	InsecureSkipTlsVerify() interface{}
	SetInsecureSkipTlsVerify(val interface{})
	InsecureSkipTlsVerifyInput() interface{}
	InternalValue() *ApiServiceV1Spec
	SetInternalValue(val *ApiServiceV1Spec)
	Service() ApiServiceV1SpecServiceOutputReference
	ServiceInput() *ApiServiceV1SpecService
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VersionPriority() *float64
	SetVersionPriority(val *float64)
	VersionPriorityInput() *float64
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
	PutService(value *ApiServiceV1SpecService)
	ResetCaBundle()
	ResetInsecureSkipTlsVerify()
	ResetService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiServiceV1SpecOutputReference
type jsiiProxy_ApiServiceV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) CaBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) CaBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) GroupPriorityMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupPriorityMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) GroupPriorityMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupPriorityMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) InsecureSkipTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) InsecureSkipTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) InternalValue() *ApiServiceV1Spec {
	var returns *ApiServiceV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) Service() ApiServiceV1SpecServiceOutputReference {
	var returns ApiServiceV1SpecServiceOutputReference
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) ServiceInput() *ApiServiceV1SpecService {
	var returns *ApiServiceV1SpecService
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) VersionPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference) VersionPriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionPriorityInput",
		&returns,
	)
	return returns
}


func NewApiServiceV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiServiceV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewApiServiceV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiServiceV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.apiServiceV1.ApiServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiServiceV1SpecOutputReference_Override(a ApiServiceV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.apiServiceV1.ApiServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetCaBundle(val *string) {
	if err := j.validateSetCaBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caBundle",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetGroupPriorityMinimum(val *float64) {
	if err := j.validateSetGroupPriorityMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupPriorityMinimum",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetInsecureSkipTlsVerify(val interface{}) {
	if err := j.validateSetInsecureSkipTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSkipTlsVerify",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetInternalValue(val *ApiServiceV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ApiServiceV1SpecOutputReference)SetVersionPriority(val *float64) {
	if err := j.validateSetVersionPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionPriority",
		val,
	)
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) PutService(value *ApiServiceV1SpecService) {
	if err := a.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) ResetCaBundle() {
	_jsii_.InvokeVoid(
		a,
		"resetCaBundle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) ResetInsecureSkipTlsVerify() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureSkipTlsVerify",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		a,
		"resetService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiServiceV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

