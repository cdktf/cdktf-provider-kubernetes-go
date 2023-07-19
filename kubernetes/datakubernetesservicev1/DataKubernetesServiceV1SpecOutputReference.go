package datakubernetesservicev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/datakubernetesservicev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesServiceV1SpecOutputReference interface {
	cdktf.ComplexObject
	AllocateLoadBalancerNodePorts() cdktf.IResolvable
	ClusterIp() *string
	ClusterIps() *[]*string
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
	ExternalIps() *[]*string
	ExternalName() *string
	ExternalTrafficPolicy() *string
	// Experimental.
	Fqn() *string
	HealthCheckNodePort() *float64
	InternalTrafficPolicy() *string
	InternalValue() *DataKubernetesServiceV1Spec
	SetInternalValue(val *DataKubernetesServiceV1Spec)
	IpFamilies() *[]*string
	IpFamilyPolicy() *string
	LoadBalancerClass() *string
	LoadBalancerIp() *string
	LoadBalancerSourceRanges() *[]*string
	Port() DataKubernetesServiceV1SpecPortList
	PublishNotReadyAddresses() cdktf.IResolvable
	Selector() cdktf.StringMap
	SessionAffinity() *string
	SessionAffinityConfig() DataKubernetesServiceV1SpecSessionAffinityConfigList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesServiceV1SpecOutputReference
type jsiiProxy_DataKubernetesServiceV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) AllocateLoadBalancerNodePorts() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allocateLoadBalancerNodePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ClusterIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ClusterIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ExternalIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ExternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) HealthCheckNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) InternalTrafficPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) InternalValue() *DataKubernetesServiceV1Spec {
	var returns *DataKubernetesServiceV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) IpFamilies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) IpFamilyPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) LoadBalancerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) LoadBalancerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) LoadBalancerSourceRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSourceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) Port() DataKubernetesServiceV1SpecPortList {
	var returns DataKubernetesServiceV1SpecPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) PublishNotReadyAddresses() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publishNotReadyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) Selector() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) SessionAffinityConfig() DataKubernetesServiceV1SpecSessionAffinityConfigList {
	var returns DataKubernetesServiceV1SpecSessionAffinityConfigList
	_jsii_.Get(
		j,
		"sessionAffinityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataKubernetesServiceV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesServiceV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesServiceV1SpecOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesServiceV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesServiceV1.DataKubernetesServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesServiceV1SpecOutputReference_Override(d DataKubernetesServiceV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesServiceV1.DataKubernetesServiceV1SpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference)SetInternalValue(val *DataKubernetesServiceV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesServiceV1SpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesServiceV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

