package statefulsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/statefulsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList interface {
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
	Get(index *float64) StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList
type jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewStatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList {
	_init_.Initialize()

	if err := validateNewStatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewStatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList_Override(s StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) Get(index *float64) StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerLifecyclePreStopTcpSocketList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

