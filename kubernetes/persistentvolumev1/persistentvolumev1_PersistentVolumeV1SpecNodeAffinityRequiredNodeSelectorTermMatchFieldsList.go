package persistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/persistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList
type jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList {
	_init_.Initialize()

	if err := validateNewPersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList_Override(p PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) Get(index *float64) PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFieldsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

