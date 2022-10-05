package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference interface {
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
	ConfigMap() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesConfigMapList
	ConfigMapInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DownwardApi() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	DownwardApiInput() *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Secret() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesSecretList
	SecretInput() interface{}
	ServiceAccountToken() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	ServiceAccountTokenInput() *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
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
	PutDownwardApi(value *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi)
	PutSecret(value interface{})
	PutServiceAccountToken(value *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken)
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

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMap() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesConfigMapList {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesConfigMapList
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ConfigMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApi() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) DownwardApiInput() *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) Secret() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesSecretList {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountToken() ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountTokenOutputReference
	_jsii_.Get(
		j,
		"serviceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ServiceAccountTokenInput() *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken
	_jsii_.Get(
		j,
		"serviceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) PutConfigMap(value interface{}) {
	if err := r.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) PutDownwardApi(value *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi) {
	if err := r.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) PutSecret(value interface{}) {
	if err := r.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) PutServiceAccountToken(value *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken) {
	if err := r.validatePutServiceAccountTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServiceAccountToken",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		r,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		r,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ResetServiceAccountToken() {
	_jsii_.InvokeVoid(
		r,
		"resetServiceAccountToken",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

