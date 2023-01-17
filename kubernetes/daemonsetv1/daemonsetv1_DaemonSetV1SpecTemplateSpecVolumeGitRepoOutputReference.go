package daemonsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/daemonsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference interface {
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
	Directory() *string
	SetDirectory(val *string)
	DirectoryInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DaemonSetV1SpecTemplateSpecVolumeGitRepo
	SetInternalValue(val *DaemonSetV1SpecTemplateSpecVolumeGitRepo)
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
	Revision() *string
	SetRevision(val *string)
	RevisionInput() *string
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
	ResetDirectory()
	ResetRepository()
	ResetRevision()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference
type jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) DirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) InternalValue() *DaemonSetV1SpecTemplateSpecVolumeGitRepo {
	var returns *DaemonSetV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference_Override(d DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetDirectory(val *string) {
	if err := j.validateSetDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directory",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetInternalValue(val *DaemonSetV1SpecTemplateSpecVolumeGitRepo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ResetDirectory() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		d,
		"resetRepository",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ResetRevision() {
	_jsii_.InvokeVoid(
		d,
		"resetRevision",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

