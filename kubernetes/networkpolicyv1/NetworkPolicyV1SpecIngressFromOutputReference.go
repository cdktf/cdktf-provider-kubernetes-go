// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/networkpolicyv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPolicyV1SpecIngressFromOutputReference interface {
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
	IpBlock() NetworkPolicyV1SpecIngressFromIpBlockOutputReference
	IpBlockInput() *NetworkPolicyV1SpecIngressFromIpBlock
	NamespaceSelector() NetworkPolicyV1SpecIngressFromNamespaceSelectorOutputReference
	NamespaceSelectorInput() *NetworkPolicyV1SpecIngressFromNamespaceSelector
	PodSelector() NetworkPolicyV1SpecIngressFromPodSelectorOutputReference
	PodSelectorInput() *NetworkPolicyV1SpecIngressFromPodSelector
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutIpBlock(value *NetworkPolicyV1SpecIngressFromIpBlock)
	PutNamespaceSelector(value *NetworkPolicyV1SpecIngressFromNamespaceSelector)
	PutPodSelector(value *NetworkPolicyV1SpecIngressFromPodSelector)
	ResetIpBlock()
	ResetNamespaceSelector()
	ResetPodSelector()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkPolicyV1SpecIngressFromOutputReference
type jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) IpBlock() NetworkPolicyV1SpecIngressFromIpBlockOutputReference {
	var returns NetworkPolicyV1SpecIngressFromIpBlockOutputReference
	_jsii_.Get(
		j,
		"ipBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) IpBlockInput() *NetworkPolicyV1SpecIngressFromIpBlock {
	var returns *NetworkPolicyV1SpecIngressFromIpBlock
	_jsii_.Get(
		j,
		"ipBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) NamespaceSelector() NetworkPolicyV1SpecIngressFromNamespaceSelectorOutputReference {
	var returns NetworkPolicyV1SpecIngressFromNamespaceSelectorOutputReference
	_jsii_.Get(
		j,
		"namespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) NamespaceSelectorInput() *NetworkPolicyV1SpecIngressFromNamespaceSelector {
	var returns *NetworkPolicyV1SpecIngressFromNamespaceSelector
	_jsii_.Get(
		j,
		"namespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) PodSelector() NetworkPolicyV1SpecIngressFromPodSelectorOutputReference {
	var returns NetworkPolicyV1SpecIngressFromPodSelectorOutputReference
	_jsii_.Get(
		j,
		"podSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) PodSelectorInput() *NetworkPolicyV1SpecIngressFromPodSelector {
	var returns *NetworkPolicyV1SpecIngressFromPodSelector
	_jsii_.Get(
		j,
		"podSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkPolicyV1SpecIngressFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkPolicyV1SpecIngressFromOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkPolicyV1SpecIngressFromOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicyV1.NetworkPolicyV1SpecIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkPolicyV1SpecIngressFromOutputReference_Override(n NetworkPolicyV1SpecIngressFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.networkPolicyV1.NetworkPolicyV1SpecIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) PutIpBlock(value *NetworkPolicyV1SpecIngressFromIpBlock) {
	if err := n.validatePutIpBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putIpBlock",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) PutNamespaceSelector(value *NetworkPolicyV1SpecIngressFromNamespaceSelector) {
	if err := n.validatePutNamespaceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putNamespaceSelector",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) PutPodSelector(value *NetworkPolicyV1SpecIngressFromPodSelector) {
	if err := n.validatePutPodSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPodSelector",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ResetIpBlock() {
	_jsii_.InvokeVoid(
		n,
		"resetIpBlock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ResetNamespaceSelector() {
	_jsii_.InvokeVoid(
		n,
		"resetNamespaceSelector",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ResetPodSelector() {
	_jsii_.InvokeVoid(
		n,
		"resetPodSelector",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

