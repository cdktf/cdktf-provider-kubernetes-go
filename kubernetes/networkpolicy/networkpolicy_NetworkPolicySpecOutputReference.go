package networkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/networkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPolicySpecOutputReference interface {
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
	Egress() NetworkPolicySpecEgressList
	EgressInput() interface{}
	// Experimental.
	Fqn() *string
	Ingress() NetworkPolicySpecIngressList
	IngressInput() interface{}
	InternalValue() *NetworkPolicySpec
	SetInternalValue(val *NetworkPolicySpec)
	PodSelector() NetworkPolicySpecPodSelectorOutputReference
	PodSelectorInput() *NetworkPolicySpecPodSelector
	PolicyTypes() *[]*string
	SetPolicyTypes(val *[]*string)
	PolicyTypesInput() *[]*string
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
	PutEgress(value interface{})
	PutIngress(value interface{})
	PutPodSelector(value *NetworkPolicySpecPodSelector)
	ResetEgress()
	ResetIngress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicySpecOutputReference
type jsiiProxy_NetworkPolicySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) Egress() NetworkPolicySpecEgressList {
	var returns NetworkPolicySpecEgressList
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) EgressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) Ingress() NetworkPolicySpecIngressList {
	var returns NetworkPolicySpecIngressList
	_jsii_.Get(
		j,
		"ingress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) IngressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) InternalValue() *NetworkPolicySpec {
	var returns *NetworkPolicySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) PodSelector() NetworkPolicySpecPodSelectorOutputReference {
	var returns NetworkPolicySpecPodSelectorOutputReference
	_jsii_.Get(
		j,
		"podSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) PodSelectorInput() *NetworkPolicySpecPodSelector {
	var returns *NetworkPolicySpecPodSelector
	_jsii_.Get(
		j,
		"podSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) PolicyTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) PolicyTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkPolicySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkPolicySpecOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkPolicySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPolicySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicy.NetworkPolicySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkPolicySpecOutputReference_Override(n NetworkPolicySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicy.NetworkPolicySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetInternalValue(val *NetworkPolicySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetPolicyTypes(val *[]*string) {
	if err := j.validateSetPolicyTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyTypes",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) PutEgress(value interface{}) {
	if err := n.validatePutEgressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putEgress",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) PutIngress(value interface{}) {
	if err := n.validatePutIngressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putIngress",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) PutPodSelector(value *NetworkPolicySpecPodSelector) {
	if err := n.validatePutPodSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPodSelector",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) ResetEgress() {
	_jsii_.InvokeVoid(
		n,
		"resetEgress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) ResetIngress() {
	_jsii_.InvokeVoid(
		n,
		"resetIngress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

