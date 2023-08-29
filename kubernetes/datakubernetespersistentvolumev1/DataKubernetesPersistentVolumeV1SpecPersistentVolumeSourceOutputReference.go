// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/datakubernetespersistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore
	AzureDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference
	AzureDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk
	AzureFile() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference
	AzureFileInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	CephFs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference
	CephFsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs
	Cinder() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference
	CinderInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder
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
	Csi() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
	CsiInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi
	Fc() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference
	FcInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc
	FlexVolume() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference
	FlexVolumeInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume
	Flocker() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference
	FlockerInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference
	GcePersistentDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk
	Glusterfs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference
	GlusterfsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs
	HostPath() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference
	HostPathInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath
	InternalValue() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource
	SetInternalValue(val *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource)
	Iscsi() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference
	IscsiInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi
	Local() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference
	LocalInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal
	Nfs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference
	NfsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs
	PhotonPersistentDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk
	Quobyte() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference
	QuobyteInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte
	Rbd() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference
	RbdInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference
	VsphereVolumeInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume
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
	PutAwsElasticBlockStore(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore)
	PutAzureDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk)
	PutAzureFile(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile)
	PutCephFs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs)
	PutCinder(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder)
	PutCsi(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi)
	PutFc(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc)
	PutFlexVolume(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume)
	PutFlocker(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker)
	PutGcePersistentDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk)
	PutGlusterfs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs)
	PutHostPath(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath)
	PutIscsi(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi)
	PutLocal(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal)
	PutNfs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs)
	PutPhotonPersistentDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk)
	PutQuobyte(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte)
	PutRbd(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd)
	PutVsphereVolume(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetCsi()
	ResetFc()
	ResetFlexVolume()
	ResetFlocker()
	ResetGcePersistentDisk()
	ResetGlusterfs()
	ResetHostPath()
	ResetIscsi()
	ResetLocal()
	ResetNfs()
	ResetPhotonPersistentDisk()
	ResetQuobyte()
	ResetRbd()
	ResetVsphereVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference
type jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AwsElasticBlockStore() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AwsElasticBlockStoreInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureFile() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureFileInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CephFs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CephFsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Cinder() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CinderInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Csi() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CsiInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Fc() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FcInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlexVolume() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlexVolumeInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Flocker() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlockerInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GcePersistentDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GcePersistentDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Glusterfs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GlusterfsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) HostPath() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) HostPathInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InternalValue() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Iscsi() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) IscsiInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Local() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) LocalInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Nfs() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) NfsInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PhotonPersistentDisk() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PhotonPersistentDiskInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Quobyte() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) QuobyteInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Rbd() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) RbdInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) VsphereVolume() DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference {
	var returns DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) VsphereVolumeInput() *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume {
	var returns *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeV1.DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference_Override(d DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPersistentVolumeV1.DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetInternalValue(val *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAwsElasticBlockStore(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore) {
	if err := d.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAzureDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureDisk) {
	if err := d.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAzureFile(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceAzureFile) {
	if err := d.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCephFs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCephFs) {
	if err := d.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCephFs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCinder(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCinder) {
	if err := d.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCinder",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCsi(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceCsi) {
	if err := d.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFc(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFc) {
	if err := d.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFlexVolume(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlexVolume) {
	if err := d.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFlocker(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceFlocker) {
	if err := d.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlocker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutGcePersistentDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk) {
	if err := d.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutGlusterfs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceGlusterfs) {
	if err := d.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutHostPath(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceHostPath) {
	if err := d.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutIscsi(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceIscsi) {
	if err := d.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIscsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutLocal(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceLocal) {
	if err := d.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutNfs(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceNfs) {
	if err := d.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutPhotonPersistentDisk(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk) {
	if err := d.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutQuobyte(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceQuobyte) {
	if err := d.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutRbd(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd) {
	if err := d.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRbd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutVsphereVolume(value *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume) {
	if err := d.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		d,
		"resetCephFs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		d,
		"resetCinder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		d,
		"resetCsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		d,
		"resetFc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		d,
		"resetFlocker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		d,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		d,
		"resetNfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		d,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		d,
		"resetRbd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

