package datakubernetespod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v3/datakubernetespod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList
type jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList {
	_init_.Initialize()

	if err := validateNewDataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList_Override(d DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) Get(index *float64) DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPodSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

