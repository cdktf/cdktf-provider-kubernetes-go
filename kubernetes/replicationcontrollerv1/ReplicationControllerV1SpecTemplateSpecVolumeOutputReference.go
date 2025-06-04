// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() ReplicationControllerV1SpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumeAzureDisk
	AzureFile() ReplicationControllerV1SpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *ReplicationControllerV1SpecTemplateSpecVolumeAzureFile
	CephFs() ReplicationControllerV1SpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *ReplicationControllerV1SpecTemplateSpecVolumeCephFs
	Cinder() ReplicationControllerV1SpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *ReplicationControllerV1SpecTemplateSpecVolumeCinder
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
	ConfigMap() ReplicationControllerV1SpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *ReplicationControllerV1SpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *ReplicationControllerV1SpecTemplateSpecVolumeCsi
	DownwardApi() ReplicationControllerV1SpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *ReplicationControllerV1SpecTemplateSpecVolumeDownwardApi
	EmptyDir() ReplicationControllerV1SpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *ReplicationControllerV1SpecTemplateSpecVolumeEmptyDir
	Ephemeral() ReplicationControllerV1SpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *ReplicationControllerV1SpecTemplateSpecVolumeEphemeral
	Fc() ReplicationControllerV1SpecTemplateSpecVolumeFcOutputReference
	FcInput() *ReplicationControllerV1SpecTemplateSpecVolumeFc
	FlexVolume() ReplicationControllerV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *ReplicationControllerV1SpecTemplateSpecVolumeFlexVolume
	Flocker() ReplicationControllerV1SpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *ReplicationControllerV1SpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() ReplicationControllerV1SpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *ReplicationControllerV1SpecTemplateSpecVolumeGitRepo
	Glusterfs() ReplicationControllerV1SpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *ReplicationControllerV1SpecTemplateSpecVolumeGlusterfs
	HostPath() ReplicationControllerV1SpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *ReplicationControllerV1SpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() ReplicationControllerV1SpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *ReplicationControllerV1SpecTemplateSpecVolumeIscsi
	Local() ReplicationControllerV1SpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *ReplicationControllerV1SpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() ReplicationControllerV1SpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *ReplicationControllerV1SpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDisk
	Projected() ReplicationControllerV1SpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() ReplicationControllerV1SpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *ReplicationControllerV1SpecTemplateSpecVolumeQuobyte
	Rbd() ReplicationControllerV1SpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *ReplicationControllerV1SpecTemplateSpecVolumeRbd
	Secret() ReplicationControllerV1SpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *ReplicationControllerV1SpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *ReplicationControllerV1SpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *ReplicationControllerV1SpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *ReplicationControllerV1SpecTemplateSpecVolumeCephFs)
	PutCinder(value *ReplicationControllerV1SpecTemplateSpecVolumeCinder)
	PutConfigMap(value *ReplicationControllerV1SpecTemplateSpecVolumeConfigMap)
	PutCsi(value *ReplicationControllerV1SpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *ReplicationControllerV1SpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *ReplicationControllerV1SpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *ReplicationControllerV1SpecTemplateSpecVolumeEphemeral)
	PutFc(value *ReplicationControllerV1SpecTemplateSpecVolumeFc)
	PutFlexVolume(value *ReplicationControllerV1SpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *ReplicationControllerV1SpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *ReplicationControllerV1SpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *ReplicationControllerV1SpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *ReplicationControllerV1SpecTemplateSpecVolumeHostPath)
	PutIscsi(value *ReplicationControllerV1SpecTemplateSpecVolumeIscsi)
	PutLocal(value *ReplicationControllerV1SpecTemplateSpecVolumeLocal)
	PutNfs(value *ReplicationControllerV1SpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *ReplicationControllerV1SpecTemplateSpecVolumeQuobyte)
	PutRbd(value *ReplicationControllerV1SpecTemplateSpecVolumeRbd)
	PutSecret(value *ReplicationControllerV1SpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecVolumeOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AzureDisk() ReplicationControllerV1SpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AzureDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumeAzureDisk {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AzureFile() ReplicationControllerV1SpecTemplateSpecVolumeAzureFileOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) AzureFileInput() *ReplicationControllerV1SpecTemplateSpecVolumeAzureFile {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) CephFs() ReplicationControllerV1SpecTemplateSpecVolumeCephFsOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) CephFsInput() *ReplicationControllerV1SpecTemplateSpecVolumeCephFs {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Cinder() ReplicationControllerV1SpecTemplateSpecVolumeCinderOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) CinderInput() *ReplicationControllerV1SpecTemplateSpecVolumeCinder {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ConfigMap() ReplicationControllerV1SpecTemplateSpecVolumeConfigMapOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ConfigMapInput() *ReplicationControllerV1SpecTemplateSpecVolumeConfigMap {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Csi() ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) CsiInput() *ReplicationControllerV1SpecTemplateSpecVolumeCsi {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) DownwardApi() ReplicationControllerV1SpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) DownwardApiInput() *ReplicationControllerV1SpecTemplateSpecVolumeDownwardApi {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) EmptyDir() ReplicationControllerV1SpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) EmptyDirInput() *ReplicationControllerV1SpecTemplateSpecVolumeEmptyDir {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Ephemeral() ReplicationControllerV1SpecTemplateSpecVolumeEphemeralOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) EphemeralInput() *ReplicationControllerV1SpecTemplateSpecVolumeEphemeral {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Fc() ReplicationControllerV1SpecTemplateSpecVolumeFcOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) FcInput() *ReplicationControllerV1SpecTemplateSpecVolumeFc {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) FlexVolume() ReplicationControllerV1SpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *ReplicationControllerV1SpecTemplateSpecVolumeFlexVolume {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Flocker() ReplicationControllerV1SpecTemplateSpecVolumeFlockerOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) FlockerInput() *ReplicationControllerV1SpecTemplateSpecVolumeFlocker {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GcePersistentDisk() ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDisk {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GitRepo() ReplicationControllerV1SpecTemplateSpecVolumeGitRepoOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GitRepoInput() *ReplicationControllerV1SpecTemplateSpecVolumeGitRepo {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Glusterfs() ReplicationControllerV1SpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GlusterfsInput() *ReplicationControllerV1SpecTemplateSpecVolumeGlusterfs {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) HostPath() ReplicationControllerV1SpecTemplateSpecVolumeHostPathOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) HostPathInput() *ReplicationControllerV1SpecTemplateSpecVolumeHostPath {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Iscsi() ReplicationControllerV1SpecTemplateSpecVolumeIscsiOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) IscsiInput() *ReplicationControllerV1SpecTemplateSpecVolumeIscsi {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Local() ReplicationControllerV1SpecTemplateSpecVolumeLocalOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) LocalInput() *ReplicationControllerV1SpecTemplateSpecVolumeLocal {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Nfs() ReplicationControllerV1SpecTemplateSpecVolumeNfsOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) NfsInput() *ReplicationControllerV1SpecTemplateSpecVolumeNfs {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Projected() ReplicationControllerV1SpecTemplateSpecVolumeProjectedList {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Quobyte() ReplicationControllerV1SpecTemplateSpecVolumeQuobyteOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) QuobyteInput() *ReplicationControllerV1SpecTemplateSpecVolumeQuobyte {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Rbd() ReplicationControllerV1SpecTemplateSpecVolumeRbdOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) RbdInput() *ReplicationControllerV1SpecTemplateSpecVolumeRbd {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Secret() ReplicationControllerV1SpecTemplateSpecVolumeSecretOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) SecretInput() *ReplicationControllerV1SpecTemplateSpecVolumeSecret {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) VsphereVolume() ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolume {
	var returns *ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReplicationControllerV1SpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecVolumeOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *ReplicationControllerV1SpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := r.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *ReplicationControllerV1SpecTemplateSpecVolumeAzureDisk) {
	if err := r.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutAzureFile(value *ReplicationControllerV1SpecTemplateSpecVolumeAzureFile) {
	if err := r.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutCephFs(value *ReplicationControllerV1SpecTemplateSpecVolumeCephFs) {
	if err := r.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCephFs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutCinder(value *ReplicationControllerV1SpecTemplateSpecVolumeCinder) {
	if err := r.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCinder",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutConfigMap(value *ReplicationControllerV1SpecTemplateSpecVolumeConfigMap) {
	if err := r.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutCsi(value *ReplicationControllerV1SpecTemplateSpecVolumeCsi) {
	if err := r.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCsi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *ReplicationControllerV1SpecTemplateSpecVolumeDownwardApi) {
	if err := r.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *ReplicationControllerV1SpecTemplateSpecVolumeEmptyDir) {
	if err := r.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutEphemeral(value *ReplicationControllerV1SpecTemplateSpecVolumeEphemeral) {
	if err := r.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutFc(value *ReplicationControllerV1SpecTemplateSpecVolumeFc) {
	if err := r.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFc",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *ReplicationControllerV1SpecTemplateSpecVolumeFlexVolume) {
	if err := r.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutFlocker(value *ReplicationControllerV1SpecTemplateSpecVolumeFlocker) {
	if err := r.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putFlocker",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *ReplicationControllerV1SpecTemplateSpecVolumeGcePersistentDisk) {
	if err := r.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutGitRepo(value *ReplicationControllerV1SpecTemplateSpecVolumeGitRepo) {
	if err := r.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *ReplicationControllerV1SpecTemplateSpecVolumeGlusterfs) {
	if err := r.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutHostPath(value *ReplicationControllerV1SpecTemplateSpecVolumeHostPath) {
	if err := r.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHostPath",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutIscsi(value *ReplicationControllerV1SpecTemplateSpecVolumeIscsi) {
	if err := r.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIscsi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutLocal(value *ReplicationControllerV1SpecTemplateSpecVolumeLocal) {
	if err := r.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLocal",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutNfs(value *ReplicationControllerV1SpecTemplateSpecVolumeNfs) {
	if err := r.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNfs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *ReplicationControllerV1SpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := r.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *ReplicationControllerV1SpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := r.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := r.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putProjected",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutQuobyte(value *ReplicationControllerV1SpecTemplateSpecVolumeQuobyte) {
	if err := r.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutRbd(value *ReplicationControllerV1SpecTemplateSpecVolumeRbd) {
	if err := r.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRbd",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutSecret(value *ReplicationControllerV1SpecTemplateSpecVolumeSecret) {
	if err := r.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *ReplicationControllerV1SpecTemplateSpecVolumeVsphereVolume) {
	if err := r.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		r,
		"resetCephFs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		r,
		"resetCinder",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		r,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		r,
		"resetCsi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		r,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		r,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		r,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		r,
		"resetFc",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		r,
		"resetFlocker",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		r,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		r,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		r,
		"resetHostPath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		r,
		"resetIscsi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		r,
		"resetLocal",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		r,
		"resetNfs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		r,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		r,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		r,
		"resetProjected",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		r,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		r,
		"resetRbd",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

