package networkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/networkpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPolicySpecIngressFromOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpBlock() NetworkPolicySpecIngressFromIpBlockOutputReference
	IpBlockInput() *NetworkPolicySpecIngressFromIpBlock
	NamespaceSelector() NetworkPolicySpecIngressFromNamespaceSelectorOutputReference
	NamespaceSelectorInput() *NetworkPolicySpecIngressFromNamespaceSelector
	PodSelector() NetworkPolicySpecIngressFromPodSelectorOutputReference
	PodSelectorInput() *NetworkPolicySpecIngressFromPodSelector
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
	PutIpBlock(value *NetworkPolicySpecIngressFromIpBlock)
	PutNamespaceSelector(value *NetworkPolicySpecIngressFromNamespaceSelector)
	PutPodSelector(value *NetworkPolicySpecIngressFromPodSelector)
	ResetIpBlock()
	ResetNamespaceSelector()
	ResetPodSelector()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicySpecIngressFromOutputReference
type jsiiProxy_NetworkPolicySpecIngressFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) IpBlock() NetworkPolicySpecIngressFromIpBlockOutputReference {
	var returns NetworkPolicySpecIngressFromIpBlockOutputReference
	_jsii_.Get(
		j,
		"ipBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) IpBlockInput() *NetworkPolicySpecIngressFromIpBlock {
	var returns *NetworkPolicySpecIngressFromIpBlock
	_jsii_.Get(
		j,
		"ipBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) NamespaceSelector() NetworkPolicySpecIngressFromNamespaceSelectorOutputReference {
	var returns NetworkPolicySpecIngressFromNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) NamespaceSelectorInput() *NetworkPolicySpecIngressFromNamespaceSelector {
	var returns *NetworkPolicySpecIngressFromNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) PodSelector() NetworkPolicySpecIngressFromPodSelectorOutputReference {
	var returns NetworkPolicySpecIngressFromPodSelectorOutputReference
	_jsii_.Get(
		j,
		"podSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) PodSelectorInput() *NetworkPolicySpecIngressFromPodSelector {
	var returns *NetworkPolicySpecIngressFromPodSelector
	_jsii_.Get(
		j,
		"podSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkPolicySpecIngressFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkPolicySpecIngressFromOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkPolicySpecIngressFromOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPolicySpecIngressFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicy.NetworkPolicySpecIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkPolicySpecIngressFromOutputReference_Override(n NetworkPolicySpecIngressFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicy.NetworkPolicySpecIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) PutIpBlock(value *NetworkPolicySpecIngressFromIpBlock) {
	if err := n.validatePutIpBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putIpBlock",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) PutNamespaceSelector(value *NetworkPolicySpecIngressFromNamespaceSelector) {
	if err := n.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) PutPodSelector(value *NetworkPolicySpecIngressFromPodSelector) {
	if err := n.validatePutPodSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPodSelector",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ResetIpBlock() {
	_jsii_.InvokeVoid(
		n,
		"resetIpBlock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		n,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ResetPodSelector() {
	_jsii_.InvokeVoid(
		n,
		"resetPodSelector",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkPolicySpecIngressFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

