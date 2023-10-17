// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/statefulset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() StatefulSetSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *StatefulSetSpecTemplateSpecVolumeAzureDisk
	AzureFile() StatefulSetSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *StatefulSetSpecTemplateSpecVolumeAzureFile
	CephFs() StatefulSetSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *StatefulSetSpecTemplateSpecVolumeCephFs
	Cinder() StatefulSetSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *StatefulSetSpecTemplateSpecVolumeCinder
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
	ConfigMap() StatefulSetSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *StatefulSetSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() StatefulSetSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *StatefulSetSpecTemplateSpecVolumeCsi
	DownwardApi() StatefulSetSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *StatefulSetSpecTemplateSpecVolumeDownwardApi
	EmptyDir() StatefulSetSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *StatefulSetSpecTemplateSpecVolumeEmptyDir
	Ephemeral() StatefulSetSpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *StatefulSetSpecTemplateSpecVolumeEphemeral
	Fc() StatefulSetSpecTemplateSpecVolumeFcOutputReference
	FcInput() *StatefulSetSpecTemplateSpecVolumeFc
	FlexVolume() StatefulSetSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *StatefulSetSpecTemplateSpecVolumeFlexVolume
	Flocker() StatefulSetSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *StatefulSetSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() StatefulSetSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() StatefulSetSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *StatefulSetSpecTemplateSpecVolumeGitRepo
	Glusterfs() StatefulSetSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *StatefulSetSpecTemplateSpecVolumeGlusterfs
	HostPath() StatefulSetSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *StatefulSetSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() StatefulSetSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *StatefulSetSpecTemplateSpecVolumeIscsi
	Local() StatefulSetSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *StatefulSetSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() StatefulSetSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *StatefulSetSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() StatefulSetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() StatefulSetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() StatefulSetSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() StatefulSetSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *StatefulSetSpecTemplateSpecVolumeQuobyte
	Rbd() StatefulSetSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *StatefulSetSpecTemplateSpecVolumeRbd
	Secret() StatefulSetSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *StatefulSetSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() StatefulSetSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *StatefulSetSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *StatefulSetSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *StatefulSetSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *StatefulSetSpecTemplateSpecVolumeCephFs)
	PutCinder(value *StatefulSetSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *StatefulSetSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *StatefulSetSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *StatefulSetSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *StatefulSetSpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *StatefulSetSpecTemplateSpecVolumeEphemeral)
	PutFc(value *StatefulSetSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *StatefulSetSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *StatefulSetSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *StatefulSetSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *StatefulSetSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *StatefulSetSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *StatefulSetSpecTemplateSpecVolumeIscsi)
	PutLocal(value *StatefulSetSpecTemplateSpecVolumeLocal)
	PutNfs(value *StatefulSetSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *StatefulSetSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *StatefulSetSpecTemplateSpecVolumeRbd)
	PutSecret(value *StatefulSetSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *StatefulSetSpecTemplateSpecVolumeVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetConfigMap()
	ResetCsi()
	ResetDownwardApi()
	ResetEmptyDir()
	ResetEphemeral()
	ResetFc()
	ResetFlexVolume()
	ResetFlocker()
	ResetGcePersistentDisk()
	ResetGitRepo()
	ResetGlusterfs()
	ResetHostPath()
	ResetIscsi()
	ResetLocal()
	ResetName()
	ResetNfs()
	ResetPersistentVolumeClaim()
	ResetPhotonPersistentDisk()
	ResetProjected()
	ResetQuobyte()
	ResetRbd()
	ResetSecret()
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

// The jsii proxy struct for StatefulSetSpecTemplateSpecVolumeOutputReference
type jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AzureDisk() StatefulSetSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *StatefulSetSpecTemplateSpecVolumeAzureDisk {
	var returns *StatefulSetSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AzureFile() StatefulSetSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) AzureFileInput() *StatefulSetSpecTemplateSpecVolumeAzureFile {
	var returns *StatefulSetSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) CephFs() StatefulSetSpecTemplateSpecVolumeCephFsOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) CephFsInput() *StatefulSetSpecTemplateSpecVolumeCephFs {
	var returns *StatefulSetSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Cinder() StatefulSetSpecTemplateSpecVolumeCinderOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) CinderInput() *StatefulSetSpecTemplateSpecVolumeCinder {
	var returns *StatefulSetSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ConfigMap() StatefulSetSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *StatefulSetSpecTemplateSpecVolumeConfigMap {
	var returns *StatefulSetSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Csi() StatefulSetSpecTemplateSpecVolumeCsiOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) CsiInput() *StatefulSetSpecTemplateSpecVolumeCsi {
	var returns *StatefulSetSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) DownwardApi() StatefulSetSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *StatefulSetSpecTemplateSpecVolumeDownwardApi {
	var returns *StatefulSetSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) EmptyDir() StatefulSetSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *StatefulSetSpecTemplateSpecVolumeEmptyDir {
	var returns *StatefulSetSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Ephemeral() StatefulSetSpecTemplateSpecVolumeEphemeralOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) EphemeralInput() *StatefulSetSpecTemplateSpecVolumeEphemeral {
	var returns *StatefulSetSpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Fc() StatefulSetSpecTemplateSpecVolumeFcOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) FcInput() *StatefulSetSpecTemplateSpecVolumeFc {
	var returns *StatefulSetSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) FlexVolume() StatefulSetSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *StatefulSetSpecTemplateSpecVolumeFlexVolume {
	var returns *StatefulSetSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Flocker() StatefulSetSpecTemplateSpecVolumeFlockerOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) FlockerInput() *StatefulSetSpecTemplateSpecVolumeFlocker {
	var returns *StatefulSetSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() StatefulSetSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GitRepo() StatefulSetSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GitRepoInput() *StatefulSetSpecTemplateSpecVolumeGitRepo {
	var returns *StatefulSetSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Glusterfs() StatefulSetSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *StatefulSetSpecTemplateSpecVolumeGlusterfs {
	var returns *StatefulSetSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) HostPath() StatefulSetSpecTemplateSpecVolumeHostPathOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) HostPathInput() *StatefulSetSpecTemplateSpecVolumeHostPath {
	var returns *StatefulSetSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Iscsi() StatefulSetSpecTemplateSpecVolumeIscsiOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) IscsiInput() *StatefulSetSpecTemplateSpecVolumeIscsi {
	var returns *StatefulSetSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Local() StatefulSetSpecTemplateSpecVolumeLocalOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) LocalInput() *StatefulSetSpecTemplateSpecVolumeLocal {
	var returns *StatefulSetSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Nfs() StatefulSetSpecTemplateSpecVolumeNfsOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) NfsInput() *StatefulSetSpecTemplateSpecVolumeNfs {
	var returns *StatefulSetSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() StatefulSetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() StatefulSetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Projected() StatefulSetSpecTemplateSpecVolumeProjectedList {
	var returns StatefulSetSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Quobyte() StatefulSetSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) QuobyteInput() *StatefulSetSpecTemplateSpecVolumeQuobyte {
	var returns *StatefulSetSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Rbd() StatefulSetSpecTemplateSpecVolumeRbdOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) RbdInput() *StatefulSetSpecTemplateSpecVolumeRbd {
	var returns *StatefulSetSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Secret() StatefulSetSpecTemplateSpecVolumeSecretOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) SecretInput() *StatefulSetSpecTemplateSpecVolumeSecret {
	var returns *StatefulSetSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) VsphereVolume() StatefulSetSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns StatefulSetSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *StatefulSetSpecTemplateSpecVolumeVsphereVolume {
	var returns *StatefulSetSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewStatefulSetSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StatefulSetSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStatefulSetSpecTemplateSpecVolumeOutputReference_Override(s StatefulSetSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := s.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *StatefulSetSpecTemplateSpecVolumeAzureDisk) {
	if err := s.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *StatefulSetSpecTemplateSpecVolumeAzureFile) {
	if err := s.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutCephFs(value *StatefulSetSpecTemplateSpecVolumeCephFs) {
	if err := s.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCephFs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutCinder(value *StatefulSetSpecTemplateSpecVolumeCinder) {
	if err := s.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCinder",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *StatefulSetSpecTemplateSpecVolumeConfigMap) {
	if err := s.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutCsi(value *StatefulSetSpecTemplateSpecVolumeCsi) {
	if err := s.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *StatefulSetSpecTemplateSpecVolumeDownwardApi) {
	if err := s.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *StatefulSetSpecTemplateSpecVolumeEmptyDir) {
	if err := s.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutEphemeral(value *StatefulSetSpecTemplateSpecVolumeEphemeral) {
	if err := s.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutFc(value *StatefulSetSpecTemplateSpecVolumeFc) {
	if err := s.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *StatefulSetSpecTemplateSpecVolumeFlexVolume) {
	if err := s.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutFlocker(value *StatefulSetSpecTemplateSpecVolumeFlocker) {
	if err := s.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFlocker",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := s.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *StatefulSetSpecTemplateSpecVolumeGitRepo) {
	if err := s.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *StatefulSetSpecTemplateSpecVolumeGlusterfs) {
	if err := s.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutHostPath(value *StatefulSetSpecTemplateSpecVolumeHostPath) {
	if err := s.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHostPath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutIscsi(value *StatefulSetSpecTemplateSpecVolumeIscsi) {
	if err := s.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIscsi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutLocal(value *StatefulSetSpecTemplateSpecVolumeLocal) {
	if err := s.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLocal",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutNfs(value *StatefulSetSpecTemplateSpecVolumeNfs) {
	if err := s.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNfs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := s.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := s.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := s.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProjected",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *StatefulSetSpecTemplateSpecVolumeQuobyte) {
	if err := s.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutRbd(value *StatefulSetSpecTemplateSpecVolumeRbd) {
	if err := s.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRbd",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutSecret(value *StatefulSetSpecTemplateSpecVolumeSecret) {
	if err := s.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecret",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *StatefulSetSpecTemplateSpecVolumeVsphereVolume) {
	if err := s.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		s,
		"resetCephFs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		s,
		"resetCinder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		s,
		"resetCsi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		s,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		s,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		s,
		"resetFc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		s,
		"resetFlocker",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		s,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		s,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		s,
		"resetHostPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		s,
		"resetIscsi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		s,
		"resetLocal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		s,
		"resetNfs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		s,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		s,
		"resetProjected",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		s,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		s,
		"resetRbd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

