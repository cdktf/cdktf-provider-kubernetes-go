package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/pod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference interface {
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Divisor() *string
	SetDivisor(val *string)
	DivisorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef
	SetInternalValue(val *PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
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
	ResetDivisor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference
type jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) Divisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"divisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) DivisorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"divisorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) InternalValue() *PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef {
	var returns *PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference_Override(p PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetDivisor(val *string) {
	if err := j.validateSetDivisorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"divisor",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetInternalValue(val *PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ResetDivisor() {
	_jsii_.InvokeVoid(
		p,
		"resetDivisor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRefOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

