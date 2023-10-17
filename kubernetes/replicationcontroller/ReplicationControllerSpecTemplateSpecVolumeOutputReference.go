// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/replicationcontroller/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() ReplicationControllerSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *ReplicationControllerSpecTemplateSpecVolumeAzureDisk
	AzureFile() ReplicationControllerSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *ReplicationControllerSpecTemplateSpecVolumeAzureFile
	CephFs() ReplicationControllerSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *ReplicationControllerSpecTemplateSpecVolumeCephFs
	Cinder() ReplicationControllerSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *ReplicationControllerSpecTemplateSpecVolumeCinder
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
	ConfigMap() ReplicationControllerSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *ReplicationControllerSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() ReplicationControllerSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *ReplicationControllerSpecTemplateSpecVolumeCsi
	DownwardApi() ReplicationControllerSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *ReplicationControllerSpecTemplateSpecVolumeDownwardApi
	EmptyDir() ReplicationControllerSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *ReplicationControllerSpecTemplateSpecVolumeEmptyDir
	Ephemeral() ReplicationControllerSpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeral
	Fc() ReplicationControllerSpecTemplateSpecVolumeFcOutputReference
	FcInput() *ReplicationControllerSpecTemplateSpecVolumeFc
	FlexVolume() ReplicationControllerSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *ReplicationControllerSpecTemplateSpecVolumeFlexVolume
	Flocker() ReplicationControllerSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *ReplicationControllerSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() ReplicationControllerSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *ReplicationControllerSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() ReplicationControllerSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *ReplicationControllerSpecTemplateSpecVolumeGitRepo
	Glusterfs() ReplicationControllerSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *ReplicationControllerSpecTemplateSpecVolumeGlusterfs
	HostPath() ReplicationControllerSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *ReplicationControllerSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() ReplicationControllerSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *ReplicationControllerSpecTemplateSpecVolumeIscsi
	Local() ReplicationControllerSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *ReplicationControllerSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() ReplicationControllerSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *ReplicationControllerSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() ReplicationControllerSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() ReplicationControllerSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *ReplicationControllerSpecTemplateSpecVolumeQuobyte
	Rbd() ReplicationControllerSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *ReplicationControllerSpecTemplateSpecVolumeRbd
	Secret() ReplicationControllerSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *ReplicationControllerSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() ReplicationControllerSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *ReplicationControllerSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *ReplicationControllerSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *ReplicationControllerSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *ReplicationControllerSpecTemplateSpecVolumeCephFs)
	PutCinder(value *ReplicationControllerSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *ReplicationControllerSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *ReplicationControllerSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *ReplicationControllerSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *ReplicationControllerSpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *ReplicationControllerSpecTemplateSpecVolumeEphemeral)
	PutFc(value *ReplicationControllerSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *ReplicationControllerSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *ReplicationControllerSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *ReplicationControllerSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *ReplicationControllerSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *ReplicationControllerSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *ReplicationControllerSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *ReplicationControllerSpecTemplateSpecVolumeIscsi)
	PutLocal(value *ReplicationControllerSpecTemplateSpecVolumeLocal)
	PutNfs(value *ReplicationControllerSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *ReplicationControllerSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *ReplicationControllerSpecTemplateSpecVolumeRbd)
	PutSecret(value *ReplicationControllerSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *ReplicationControllerSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for ReplicationControllerSpecTemplateSpecVolumeOutputReference
type jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AzureDisk() ReplicationControllerSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *ReplicationControllerSpecTemplateSpecVolumeAzureDisk {
	var returns *ReplicationControllerSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AzureFile() ReplicationControllerSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) AzureFileInput() *ReplicationControllerSpecTemplateSpecVolumeAzureFile {
	var returns *ReplicationControllerSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) CephFs() ReplicationControllerSpecTemplateSpecVolumeCephFsOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) CephFsInput() *ReplicationControllerSpecTemplateSpecVolumeCephFs {
	var returns *ReplicationControllerSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Cinder() ReplicationControllerSpecTemplateSpecVolumeCinderOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) CinderInput() *ReplicationControllerSpecTemplateSpecVolumeCinder {
	var returns *ReplicationControllerSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ConfigMap() ReplicationControllerSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *ReplicationControllerSpecTemplateSpecVolumeConfigMap {
	var returns *ReplicationControllerSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Csi() ReplicationControllerSpecTemplateSpecVolumeCsiOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) CsiInput() *ReplicationControllerSpecTemplateSpecVolumeCsi {
	var returns *ReplicationControllerSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) DownwardApi() ReplicationControllerSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *ReplicationControllerSpecTemplateSpecVolumeDownwardApi {
	var returns *ReplicationControllerSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) EmptyDir() ReplicationControllerSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *ReplicationControllerSpecTemplateSpecVolumeEmptyDir {
	var returns *ReplicationControllerSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Ephemeral() ReplicationControllerSpecTemplateSpecVolumeEphemeralOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) EphemeralInput() *ReplicationControllerSpecTemplateSpecVolumeEphemeral {
	var returns *ReplicationControllerSpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Fc() ReplicationControllerSpecTemplateSpecVolumeFcOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) FcInput() *ReplicationControllerSpecTemplateSpecVolumeFc {
	var returns *ReplicationControllerSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) FlexVolume() ReplicationControllerSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *ReplicationControllerSpecTemplateSpecVolumeFlexVolume {
	var returns *ReplicationControllerSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Flocker() ReplicationControllerSpecTemplateSpecVolumeFlockerOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) FlockerInput() *ReplicationControllerSpecTemplateSpecVolumeFlocker {
	var returns *ReplicationControllerSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() ReplicationControllerSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *ReplicationControllerSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *ReplicationControllerSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GitRepo() ReplicationControllerSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GitRepoInput() *ReplicationControllerSpecTemplateSpecVolumeGitRepo {
	var returns *ReplicationControllerSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Glusterfs() ReplicationControllerSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *ReplicationControllerSpecTemplateSpecVolumeGlusterfs {
	var returns *ReplicationControllerSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) HostPath() ReplicationControllerSpecTemplateSpecVolumeHostPathOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) HostPathInput() *ReplicationControllerSpecTemplateSpecVolumeHostPath {
	var returns *ReplicationControllerSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Iscsi() ReplicationControllerSpecTemplateSpecVolumeIscsiOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) IscsiInput() *ReplicationControllerSpecTemplateSpecVolumeIscsi {
	var returns *ReplicationControllerSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Local() ReplicationControllerSpecTemplateSpecVolumeLocalOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) LocalInput() *ReplicationControllerSpecTemplateSpecVolumeLocal {
	var returns *ReplicationControllerSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Nfs() ReplicationControllerSpecTemplateSpecVolumeNfsOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) NfsInput() *ReplicationControllerSpecTemplateSpecVolumeNfs {
	var returns *ReplicationControllerSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Projected() ReplicationControllerSpecTemplateSpecVolumeProjectedList {
	var returns ReplicationControllerSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Quobyte() ReplicationControllerSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) QuobyteInput() *ReplicationControllerSpecTemplateSpecVolumeQuobyte {
	var returns *ReplicationControllerSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Rbd() ReplicationControllerSpecTemplateSpecVolumeRbdOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) RbdInput() *ReplicationControllerSpecTemplateSpecVolumeRbd {
	var returns *ReplicationControllerSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Secret() ReplicationControllerSpecTemplateSpecVolumeSecretOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) SecretInput() *ReplicationControllerSpecTemplateSpecVolumeSecret {
	var returns *ReplicationControllerSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) VsphereVolume() ReplicationControllerSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns ReplicationControllerSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *ReplicationControllerSpecTemplateSpecVolumeVsphereVolume {
	var returns *ReplicationControllerSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReplicationControllerSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReplicationControllerSpecTemplateSpecVolumeOutputReference_Override(r ReplicationControllerSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationController.ReplicationControllerSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *ReplicationControllerSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := r.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *ReplicationControllerSpecTemplateSpecVolumeAzureDisk) {
	if err := r.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *ReplicationControllerSpecTemplateSpecVolumeAzureFile) {
	if err := r.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutCephFs(value *ReplicationControllerSpecTemplateSpecVolumeCephFs) {
	if err := r.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCephFs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutCinder(value *ReplicationControllerSpecTemplateSpecVolumeCinder) {
	if err := r.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCinder",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *ReplicationControllerSpecTemplateSpecVolumeConfigMap) {
	if err := r.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutCsi(value *ReplicationControllerSpecTemplateSpecVolumeCsi) {
	if err := r.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCsi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *ReplicationControllerSpecTemplateSpecVolumeDownwardApi) {
	if err := r.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *ReplicationControllerSpecTemplateSpecVolumeEmptyDir) {
	if err := r.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutEphemeral(value *ReplicationControllerSpecTemplateSpecVolumeEphemeral) {
	if err := r.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutFc(value *ReplicationControllerSpecTemplateSpecVolumeFc) {
	if err := r.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFc",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *ReplicationControllerSpecTemplateSpecVolumeFlexVolume) {
	if err := r.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutFlocker(value *ReplicationControllerSpecTemplateSpecVolumeFlocker) {
	if err := r.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFlocker",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *ReplicationControllerSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := r.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *ReplicationControllerSpecTemplateSpecVolumeGitRepo) {
	if err := r.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *ReplicationControllerSpecTemplateSpecVolumeGlusterfs) {
	if err := r.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutHostPath(value *ReplicationControllerSpecTemplateSpecVolumeHostPath) {
	if err := r.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHostPath",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutIscsi(value *ReplicationControllerSpecTemplateSpecVolumeIscsi) {
	if err := r.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIscsi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutLocal(value *ReplicationControllerSpecTemplateSpecVolumeLocal) {
	if err := r.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLocal",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutNfs(value *ReplicationControllerSpecTemplateSpecVolumeNfs) {
	if err := r.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNfs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := r.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *ReplicationControllerSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := r.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := r.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putProjected",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *ReplicationControllerSpecTemplateSpecVolumeQuobyte) {
	if err := r.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutRbd(value *ReplicationControllerSpecTemplateSpecVolumeRbd) {
	if err := r.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRbd",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutSecret(value *ReplicationControllerSpecTemplateSpecVolumeSecret) {
	if err := r.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *ReplicationControllerSpecTemplateSpecVolumeVsphereVolume) {
	if err := r.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		r,
		"resetCephFs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		r,
		"resetCinder",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		r,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		r,
		"resetCsi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		r,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		r,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		r,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		r,
		"resetFc",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		r,
		"resetFlocker",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		r,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		r,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		r,
		"resetHostPath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		r,
		"resetIscsi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		r,
		"resetLocal",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		r,
		"resetNfs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		r,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		r,
		"resetProjected",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		r,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		r,
		"resetRbd",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

