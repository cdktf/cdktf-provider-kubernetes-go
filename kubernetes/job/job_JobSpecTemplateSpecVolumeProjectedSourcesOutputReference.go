package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecVolumeProjectedSourcesOutputReference interface {
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
	ConfigMap() JobSpecTemplateSpecVolumeProjectedSourcesConfigMapList
	ConfigMapInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DownwardApi() JobSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	DownwardApiInput() *JobSpecTemplateSpecVolumeProjectedSourcesDownwardApi
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Secret() JobSpecTemplateSpecVolumeProjectedSourcesSecretList
	SecretInput() interface{}
	ServiceAccountToken() JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	ServiceAccountTokenInput() *JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
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
	PutConfigMap(value interface{})
	PutDownwardApi(value *JobSpecTemplateSpecVolumeProjectedSourcesDownwardApi)
	PutSecret(value interface{})
	PutServiceAccountToken(value *JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken)
	ResetConfigMap()
	ResetDownwardApi()
	ResetSecret()
	ResetServiceAccountToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobSpecTemplateSpecVolumeProjectedSourcesOutputReference
type jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMap() JobSpecTemplateSpecVolumeProjectedSourcesConfigMapList {
	var returns JobSpecTemplateSpecVolumeProjectedSourcesConfigMapList
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApi() JobSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference {
	var returns JobSpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApiInput() *JobSpecTemplateSpecVolumeProjectedSourcesDownwardApi {
	var returns *JobSpecTemplateSpecVolumeProjectedSourcesDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) Secret() JobSpecTemplateSpecVolumeProjectedSourcesSecretList {
	var returns JobSpecTemplateSpecVolumeProjectedSourcesSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountToken() JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference {
	var returns JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	_jsii_.Get(
		j,
		"serviceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountTokenInput() *JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken {
	var returns *JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
	_jsii_.Get(
		j,
		"serviceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecVolumeProjectedSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecVolumeProjectedSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecVolumeProjectedSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecVolumeProjectedSourcesOutputReference_Override(j JobSpecTemplateSpecVolumeProjectedSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutConfigMap(value interface{}) {
	if err := j.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutDownwardApi(value *JobSpecTemplateSpecVolumeProjectedSourcesDownwardApi) {
	if err := j.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutSecret(value interface{}) {
	if err := j.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecret",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) PutServiceAccountToken(value *JobSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken) {
	if err := j.validatePutServiceAccountTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putServiceAccountToken",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		j,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		j,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		j,
		"resetSecret",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetServiceAccountToken() {
	_jsii_.InvokeVoid(
		j,
		"resetServiceAccountToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeProjectedSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

