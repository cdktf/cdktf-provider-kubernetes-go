package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/podv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference interface {
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
	LabelSelector() PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorList
	LabelSelectorInput() interface{}
	Namespaces() *[]*string
	SetNamespaces(val *[]*string)
	NamespacesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopologyKey() *string
	SetTopologyKey(val *string)
	TopologyKeyInput() *string
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
	PutLabelSelector(value interface{})
	ResetLabelSelector()
	ResetNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference
type jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) LabelSelector() PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorList {
	var returns PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorList
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) LabelSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) Namespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"namespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) NamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"namespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TopologyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) TopologyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyKeyInput",
		&returns,
	)
	return returns
}


func NewPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference_Override(p PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetNamespaces(val *[]*string) {
	if err := j.validateSetNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaces",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference)SetTopologyKey(val *string) {
	if err := j.validateSetTopologyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyKey",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) PutLabelSelector(value interface{}) {
	if err := p.validatePutLabelSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLabelSelector",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ResetLabelSelector() {
	_jsii_.InvokeVoid(
		p,
		"resetLabelSelector",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ResetNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

