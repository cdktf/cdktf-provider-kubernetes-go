package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList interface {
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
	Get(index *float64) ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList_Override(r ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) Get(index *float64) ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeaderList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
