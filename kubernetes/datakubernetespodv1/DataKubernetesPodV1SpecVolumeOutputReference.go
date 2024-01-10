// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespodv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/datakubernetespodv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPodV1SpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DataKubernetesPodV1SpecVolumeAwsElasticBlockStoreList
	AzureDisk() DataKubernetesPodV1SpecVolumeAzureDiskList
	AzureFile() DataKubernetesPodV1SpecVolumeAzureFileList
	CephFs() DataKubernetesPodV1SpecVolumeCephFsList
	Cinder() DataKubernetesPodV1SpecVolumeCinderList
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
	ConfigMap() DataKubernetesPodV1SpecVolumeConfigMapList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DataKubernetesPodV1SpecVolumeCsiList
	DownwardApi() DataKubernetesPodV1SpecVolumeDownwardApiList
	EmptyDir() DataKubernetesPodV1SpecVolumeEmptyDirList
	Ephemeral() DataKubernetesPodV1SpecVolumeEphemeralList
	Fc() DataKubernetesPodV1SpecVolumeFcList
	FlexVolume() DataKubernetesPodV1SpecVolumeFlexVolumeList
	Flocker() DataKubernetesPodV1SpecVolumeFlockerList
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DataKubernetesPodV1SpecVolumeGcePersistentDiskList
	GitRepo() DataKubernetesPodV1SpecVolumeGitRepoList
	Glusterfs() DataKubernetesPodV1SpecVolumeGlusterfsList
	HostPath() DataKubernetesPodV1SpecVolumeHostPathList
	InternalValue() *DataKubernetesPodV1SpecVolume
	SetInternalValue(val *DataKubernetesPodV1SpecVolume)
	Iscsi() DataKubernetesPodV1SpecVolumeIscsiList
	Local() DataKubernetesPodV1SpecVolumeLocalList
	Name() *string
	Nfs() DataKubernetesPodV1SpecVolumeNfsList
	PersistentVolumeClaim() DataKubernetesPodV1SpecVolumePersistentVolumeClaimList
	PhotonPersistentDisk() DataKubernetesPodV1SpecVolumePhotonPersistentDiskList
	Projected() DataKubernetesPodV1SpecVolumeProjectedList
	Quobyte() DataKubernetesPodV1SpecVolumeQuobyteList
	Rbd() DataKubernetesPodV1SpecVolumeRbdList
	Secret() DataKubernetesPodV1SpecVolumeSecretList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DataKubernetesPodV1SpecVolumeVsphereVolumeList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPodV1SpecVolumeOutputReference
type jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) AwsElasticBlockStore() DataKubernetesPodV1SpecVolumeAwsElasticBlockStoreList {
	var returns DataKubernetesPodV1SpecVolumeAwsElasticBlockStoreList
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) AzureDisk() DataKubernetesPodV1SpecVolumeAzureDiskList {
	var returns DataKubernetesPodV1SpecVolumeAzureDiskList
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) AzureFile() DataKubernetesPodV1SpecVolumeAzureFileList {
	var returns DataKubernetesPodV1SpecVolumeAzureFileList
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) CephFs() DataKubernetesPodV1SpecVolumeCephFsList {
	var returns DataKubernetesPodV1SpecVolumeCephFsList
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Cinder() DataKubernetesPodV1SpecVolumeCinderList {
	var returns DataKubernetesPodV1SpecVolumeCinderList
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) ConfigMap() DataKubernetesPodV1SpecVolumeConfigMapList {
	var returns DataKubernetesPodV1SpecVolumeConfigMapList
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Csi() DataKubernetesPodV1SpecVolumeCsiList {
	var returns DataKubernetesPodV1SpecVolumeCsiList
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) DownwardApi() DataKubernetesPodV1SpecVolumeDownwardApiList {
	var returns DataKubernetesPodV1SpecVolumeDownwardApiList
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) EmptyDir() DataKubernetesPodV1SpecVolumeEmptyDirList {
	var returns DataKubernetesPodV1SpecVolumeEmptyDirList
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Ephemeral() DataKubernetesPodV1SpecVolumeEphemeralList {
	var returns DataKubernetesPodV1SpecVolumeEphemeralList
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Fc() DataKubernetesPodV1SpecVolumeFcList {
	var returns DataKubernetesPodV1SpecVolumeFcList
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) FlexVolume() DataKubernetesPodV1SpecVolumeFlexVolumeList {
	var returns DataKubernetesPodV1SpecVolumeFlexVolumeList
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Flocker() DataKubernetesPodV1SpecVolumeFlockerList {
	var returns DataKubernetesPodV1SpecVolumeFlockerList
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GcePersistentDisk() DataKubernetesPodV1SpecVolumeGcePersistentDiskList {
	var returns DataKubernetesPodV1SpecVolumeGcePersistentDiskList
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GitRepo() DataKubernetesPodV1SpecVolumeGitRepoList {
	var returns DataKubernetesPodV1SpecVolumeGitRepoList
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Glusterfs() DataKubernetesPodV1SpecVolumeGlusterfsList {
	var returns DataKubernetesPodV1SpecVolumeGlusterfsList
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) HostPath() DataKubernetesPodV1SpecVolumeHostPathList {
	var returns DataKubernetesPodV1SpecVolumeHostPathList
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) InternalValue() *DataKubernetesPodV1SpecVolume {
	var returns *DataKubernetesPodV1SpecVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Iscsi() DataKubernetesPodV1SpecVolumeIscsiList {
	var returns DataKubernetesPodV1SpecVolumeIscsiList
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Local() DataKubernetesPodV1SpecVolumeLocalList {
	var returns DataKubernetesPodV1SpecVolumeLocalList
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Nfs() DataKubernetesPodV1SpecVolumeNfsList {
	var returns DataKubernetesPodV1SpecVolumeNfsList
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) PersistentVolumeClaim() DataKubernetesPodV1SpecVolumePersistentVolumeClaimList {
	var returns DataKubernetesPodV1SpecVolumePersistentVolumeClaimList
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) PhotonPersistentDisk() DataKubernetesPodV1SpecVolumePhotonPersistentDiskList {
	var returns DataKubernetesPodV1SpecVolumePhotonPersistentDiskList
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Projected() DataKubernetesPodV1SpecVolumeProjectedList {
	var returns DataKubernetesPodV1SpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Quobyte() DataKubernetesPodV1SpecVolumeQuobyteList {
	var returns DataKubernetesPodV1SpecVolumeQuobyteList
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Rbd() DataKubernetesPodV1SpecVolumeRbdList {
	var returns DataKubernetesPodV1SpecVolumeRbdList
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Secret() DataKubernetesPodV1SpecVolumeSecretList {
	var returns DataKubernetesPodV1SpecVolumeSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) VsphereVolume() DataKubernetesPodV1SpecVolumeVsphereVolumeList {
	var returns DataKubernetesPodV1SpecVolumeVsphereVolumeList
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}


func NewDataKubernetesPodV1SpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesPodV1SpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPodV1SpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPodV1.DataKubernetesPodV1SpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPodV1SpecVolumeOutputReference_Override(d DataKubernetesPodV1SpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPodV1.DataKubernetesPodV1SpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference)SetInternalValue(val *DataKubernetesPodV1SpecVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPodV1SpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

