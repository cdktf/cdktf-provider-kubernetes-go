// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v11/daemonset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaemonsetSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() DaemonsetSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *DaemonsetSpecTemplateSpecVolumeAzureDisk
	AzureFile() DaemonsetSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *DaemonsetSpecTemplateSpecVolumeAzureFile
	CephFs() DaemonsetSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *DaemonsetSpecTemplateSpecVolumeCephFs
	Cinder() DaemonsetSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *DaemonsetSpecTemplateSpecVolumeCinder
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
	ConfigMap() DaemonsetSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *DaemonsetSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DaemonsetSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *DaemonsetSpecTemplateSpecVolumeCsi
	DownwardApi() DaemonsetSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *DaemonsetSpecTemplateSpecVolumeDownwardApi
	EmptyDir() DaemonsetSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *DaemonsetSpecTemplateSpecVolumeEmptyDir
	Ephemeral() DaemonsetSpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *DaemonsetSpecTemplateSpecVolumeEphemeral
	Fc() DaemonsetSpecTemplateSpecVolumeFcOutputReference
	FcInput() *DaemonsetSpecTemplateSpecVolumeFc
	FlexVolume() DaemonsetSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *DaemonsetSpecTemplateSpecVolumeFlexVolume
	Flocker() DaemonsetSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *DaemonsetSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DaemonsetSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *DaemonsetSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() DaemonsetSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *DaemonsetSpecTemplateSpecVolumeGitRepo
	Glusterfs() DaemonsetSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *DaemonsetSpecTemplateSpecVolumeGlusterfs
	HostPath() DaemonsetSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *DaemonsetSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() DaemonsetSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *DaemonsetSpecTemplateSpecVolumeIscsi
	Local() DaemonsetSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *DaemonsetSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() DaemonsetSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *DaemonsetSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() DaemonsetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() DaemonsetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *DaemonsetSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() DaemonsetSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() DaemonsetSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *DaemonsetSpecTemplateSpecVolumeQuobyte
	Rbd() DaemonsetSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *DaemonsetSpecTemplateSpecVolumeRbd
	Secret() DaemonsetSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *DaemonsetSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DaemonsetSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *DaemonsetSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *DaemonsetSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *DaemonsetSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *DaemonsetSpecTemplateSpecVolumeCephFs)
	PutCinder(value *DaemonsetSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *DaemonsetSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *DaemonsetSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *DaemonsetSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *DaemonsetSpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *DaemonsetSpecTemplateSpecVolumeEphemeral)
	PutFc(value *DaemonsetSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *DaemonsetSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *DaemonsetSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *DaemonsetSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *DaemonsetSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *DaemonsetSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *DaemonsetSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *DaemonsetSpecTemplateSpecVolumeIscsi)
	PutLocal(value *DaemonsetSpecTemplateSpecVolumeLocal)
	PutNfs(value *DaemonsetSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *DaemonsetSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *DaemonsetSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *DaemonsetSpecTemplateSpecVolumeRbd)
	PutSecret(value *DaemonsetSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *DaemonsetSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for DaemonsetSpecTemplateSpecVolumeOutputReference
type jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AzureDisk() DaemonsetSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *DaemonsetSpecTemplateSpecVolumeAzureDisk {
	var returns *DaemonsetSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AzureFile() DaemonsetSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) AzureFileInput() *DaemonsetSpecTemplateSpecVolumeAzureFile {
	var returns *DaemonsetSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) CephFs() DaemonsetSpecTemplateSpecVolumeCephFsOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) CephFsInput() *DaemonsetSpecTemplateSpecVolumeCephFs {
	var returns *DaemonsetSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Cinder() DaemonsetSpecTemplateSpecVolumeCinderOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) CinderInput() *DaemonsetSpecTemplateSpecVolumeCinder {
	var returns *DaemonsetSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ConfigMap() DaemonsetSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *DaemonsetSpecTemplateSpecVolumeConfigMap {
	var returns *DaemonsetSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Csi() DaemonsetSpecTemplateSpecVolumeCsiOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) CsiInput() *DaemonsetSpecTemplateSpecVolumeCsi {
	var returns *DaemonsetSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) DownwardApi() DaemonsetSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *DaemonsetSpecTemplateSpecVolumeDownwardApi {
	var returns *DaemonsetSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) EmptyDir() DaemonsetSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *DaemonsetSpecTemplateSpecVolumeEmptyDir {
	var returns *DaemonsetSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Ephemeral() DaemonsetSpecTemplateSpecVolumeEphemeralOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) EphemeralInput() *DaemonsetSpecTemplateSpecVolumeEphemeral {
	var returns *DaemonsetSpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Fc() DaemonsetSpecTemplateSpecVolumeFcOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) FcInput() *DaemonsetSpecTemplateSpecVolumeFc {
	var returns *DaemonsetSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) FlexVolume() DaemonsetSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *DaemonsetSpecTemplateSpecVolumeFlexVolume {
	var returns *DaemonsetSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Flocker() DaemonsetSpecTemplateSpecVolumeFlockerOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) FlockerInput() *DaemonsetSpecTemplateSpecVolumeFlocker {
	var returns *DaemonsetSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() DaemonsetSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *DaemonsetSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *DaemonsetSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GitRepo() DaemonsetSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GitRepoInput() *DaemonsetSpecTemplateSpecVolumeGitRepo {
	var returns *DaemonsetSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Glusterfs() DaemonsetSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *DaemonsetSpecTemplateSpecVolumeGlusterfs {
	var returns *DaemonsetSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) HostPath() DaemonsetSpecTemplateSpecVolumeHostPathOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) HostPathInput() *DaemonsetSpecTemplateSpecVolumeHostPath {
	var returns *DaemonsetSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Iscsi() DaemonsetSpecTemplateSpecVolumeIscsiOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) IscsiInput() *DaemonsetSpecTemplateSpecVolumeIscsi {
	var returns *DaemonsetSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Local() DaemonsetSpecTemplateSpecVolumeLocalOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) LocalInput() *DaemonsetSpecTemplateSpecVolumeLocal {
	var returns *DaemonsetSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Nfs() DaemonsetSpecTemplateSpecVolumeNfsOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) NfsInput() *DaemonsetSpecTemplateSpecVolumeNfs {
	var returns *DaemonsetSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() DaemonsetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() DaemonsetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *DaemonsetSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *DaemonsetSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Projected() DaemonsetSpecTemplateSpecVolumeProjectedList {
	var returns DaemonsetSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Quobyte() DaemonsetSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) QuobyteInput() *DaemonsetSpecTemplateSpecVolumeQuobyte {
	var returns *DaemonsetSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Rbd() DaemonsetSpecTemplateSpecVolumeRbdOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) RbdInput() *DaemonsetSpecTemplateSpecVolumeRbd {
	var returns *DaemonsetSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Secret() DaemonsetSpecTemplateSpecVolumeSecretOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) SecretInput() *DaemonsetSpecTemplateSpecVolumeSecret {
	var returns *DaemonsetSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) VsphereVolume() DaemonsetSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns DaemonsetSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *DaemonsetSpecTemplateSpecVolumeVsphereVolume {
	var returns *DaemonsetSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewDaemonsetSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DaemonsetSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonsetSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDaemonsetSpecTemplateSpecVolumeOutputReference_Override(d DaemonsetSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := d.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *DaemonsetSpecTemplateSpecVolumeAzureDisk) {
	if err := d.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *DaemonsetSpecTemplateSpecVolumeAzureFile) {
	if err := d.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutCephFs(value *DaemonsetSpecTemplateSpecVolumeCephFs) {
	if err := d.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCephFs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutCinder(value *DaemonsetSpecTemplateSpecVolumeCinder) {
	if err := d.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCinder",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *DaemonsetSpecTemplateSpecVolumeConfigMap) {
	if err := d.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutCsi(value *DaemonsetSpecTemplateSpecVolumeCsi) {
	if err := d.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *DaemonsetSpecTemplateSpecVolumeDownwardApi) {
	if err := d.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *DaemonsetSpecTemplateSpecVolumeEmptyDir) {
	if err := d.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutEphemeral(value *DaemonsetSpecTemplateSpecVolumeEphemeral) {
	if err := d.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutFc(value *DaemonsetSpecTemplateSpecVolumeFc) {
	if err := d.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *DaemonsetSpecTemplateSpecVolumeFlexVolume) {
	if err := d.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutFlocker(value *DaemonsetSpecTemplateSpecVolumeFlocker) {
	if err := d.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlocker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *DaemonsetSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := d.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *DaemonsetSpecTemplateSpecVolumeGitRepo) {
	if err := d.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *DaemonsetSpecTemplateSpecVolumeGlusterfs) {
	if err := d.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutHostPath(value *DaemonsetSpecTemplateSpecVolumeHostPath) {
	if err := d.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutIscsi(value *DaemonsetSpecTemplateSpecVolumeIscsi) {
	if err := d.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIscsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutLocal(value *DaemonsetSpecTemplateSpecVolumeLocal) {
	if err := d.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutNfs(value *DaemonsetSpecTemplateSpecVolumeNfs) {
	if err := d.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := d.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *DaemonsetSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := d.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := d.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjected",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *DaemonsetSpecTemplateSpecVolumeQuobyte) {
	if err := d.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutRbd(value *DaemonsetSpecTemplateSpecVolumeRbd) {
	if err := d.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRbd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutSecret(value *DaemonsetSpecTemplateSpecVolumeSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *DaemonsetSpecTemplateSpecVolumeVsphereVolume) {
	if err := d.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		d,
		"resetCephFs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		d,
		"resetCinder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		d,
		"resetCsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		d,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		d,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		d,
		"resetFc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		d,
		"resetFlocker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		d,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		d,
		"resetNfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		d,
		"resetProjected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		d,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		d,
		"resetRbd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DaemonsetSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

