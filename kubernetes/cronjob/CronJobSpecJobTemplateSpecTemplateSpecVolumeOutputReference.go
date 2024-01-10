// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/cronjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk
	AzureFile() CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile
	CephFs() CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs
	Cinder() CronJobSpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder
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
	ConfigMap() CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() CronJobSpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi
	DownwardApi() CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi
	EmptyDir() CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir
	Ephemeral() CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral
	Fc() CronJobSpecJobTemplateSpecTemplateSpecVolumeFcOutputReference
	FcInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc
	FlexVolume() CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume
	Flocker() CronJobSpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo
	Glusterfs() CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs
	HostPath() CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi
	Local() CronJobSpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() CronJobSpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte
	Rbd() CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd
	Secret() CronJobSpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs)
	PutCinder(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral)
	PutFc(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi)
	PutLocal(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal)
	PutNfs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd)
	PutSecret(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference
type jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureFile() CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureFileInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) CephFs() CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) CephFsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Cinder() CronJobSpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) CinderInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ConfigMap() CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Csi() CronJobSpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) CsiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) DownwardApi() CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) EmptyDir() CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Ephemeral() CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeralOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) EphemeralInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Fc() CronJobSpecJobTemplateSpecTemplateSpecVolumeFcOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) FcInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlexVolume() CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Flocker() CronJobSpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlockerInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GitRepo() CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GitRepoInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Glusterfs() CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) HostPath() CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) HostPathInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Iscsi() CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) IscsiInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Local() CronJobSpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) LocalInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Nfs() CronJobSpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) NfsInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Projected() CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Quobyte() CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) QuobyteInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Rbd() CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) RbdInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Secret() CronJobSpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) SecretInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) VsphereVolume() CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference_Override(c CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := c.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk) {
	if err := c.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureFile) {
	if err := c.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCephFs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCephFs) {
	if err := c.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCephFs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCinder(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCinder) {
	if err := c.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCinder",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap) {
	if err := c.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCsi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeCsi) {
	if err := c.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCsi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeDownwardApi) {
	if err := c.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir) {
	if err := c.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutEphemeral(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral) {
	if err := c.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFc(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFc) {
	if err := c.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlexVolume) {
	if err := c.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFlocker(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeFlocker) {
	if err := c.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFlocker",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := c.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGitRepo) {
	if err := c.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs) {
	if err := c.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutHostPath(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath) {
	if err := c.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutIscsi(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeIscsi) {
	if err := c.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIscsi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutLocal(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeLocal) {
	if err := c.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocal",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutNfs(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeNfs) {
	if err := c.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *CronJobSpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := c.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *CronJobSpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := c.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := c.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProjected",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeQuobyte) {
	if err := c.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutRbd(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeRbd) {
	if err := c.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRbd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutSecret(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeSecret) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume) {
	if err := c.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		c,
		"resetCephFs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		c,
		"resetCinder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		c,
		"resetCsi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		c,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		c,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		c,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		c,
		"resetFc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		c,
		"resetFlocker",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		c,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		c,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		c,
		"resetHostPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		c,
		"resetIscsi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		c,
		"resetLocal",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		c,
		"resetNfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		c,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		c,
		"resetProjected",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		c,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		c,
		"resetRbd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

