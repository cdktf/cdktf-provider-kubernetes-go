package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobV1SpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() JobV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() JobV1SpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *JobV1SpecTemplateSpecVolumeAzureDisk
	AzureFile() JobV1SpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *JobV1SpecTemplateSpecVolumeAzureFile
	CephFs() JobV1SpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *JobV1SpecTemplateSpecVolumeCephFs
	Cinder() JobV1SpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *JobV1SpecTemplateSpecVolumeCinder
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
	ConfigMap() JobV1SpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *JobV1SpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() JobV1SpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *JobV1SpecTemplateSpecVolumeCsi
	DownwardApi() JobV1SpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *JobV1SpecTemplateSpecVolumeDownwardApi
	EmptyDir() JobV1SpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *JobV1SpecTemplateSpecVolumeEmptyDir
	Fc() JobV1SpecTemplateSpecVolumeFcOutputReference
	FcInput() *JobV1SpecTemplateSpecVolumeFc
	FlexVolume() JobV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *JobV1SpecTemplateSpecVolumeFlexVolume
	Flocker() JobV1SpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *JobV1SpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() JobV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *JobV1SpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() JobV1SpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *JobV1SpecTemplateSpecVolumeGitRepo
	Glusterfs() JobV1SpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *JobV1SpecTemplateSpecVolumeGlusterfs
	HostPath() JobV1SpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *JobV1SpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() JobV1SpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *JobV1SpecTemplateSpecVolumeIscsi
	Local() JobV1SpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *JobV1SpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() JobV1SpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *JobV1SpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() JobV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *JobV1SpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() JobV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *JobV1SpecTemplateSpecVolumePhotonPersistentDisk
	Projected() JobV1SpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() JobV1SpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *JobV1SpecTemplateSpecVolumeQuobyte
	Rbd() JobV1SpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *JobV1SpecTemplateSpecVolumeRbd
	Secret() JobV1SpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *JobV1SpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() JobV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *JobV1SpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *JobV1SpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *JobV1SpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *JobV1SpecTemplateSpecVolumeCephFs)
	PutCinder(value *JobV1SpecTemplateSpecVolumeCinder)
	PutConfigMap(value *JobV1SpecTemplateSpecVolumeConfigMap)
	PutCsi(value *JobV1SpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *JobV1SpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *JobV1SpecTemplateSpecVolumeEmptyDir)
	PutFc(value *JobV1SpecTemplateSpecVolumeFc)
	PutFlexVolume(value *JobV1SpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *JobV1SpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *JobV1SpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *JobV1SpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *JobV1SpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *JobV1SpecTemplateSpecVolumeHostPath)
	PutIscsi(value *JobV1SpecTemplateSpecVolumeIscsi)
	PutLocal(value *JobV1SpecTemplateSpecVolumeLocal)
	PutNfs(value *JobV1SpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *JobV1SpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *JobV1SpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *JobV1SpecTemplateSpecVolumeQuobyte)
	PutRbd(value *JobV1SpecTemplateSpecVolumeRbd)
	PutSecret(value *JobV1SpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *JobV1SpecTemplateSpecVolumeVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetConfigMap()
	ResetCsi()
	ResetDownwardApi()
	ResetEmptyDir()
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

// The jsii proxy struct for JobV1SpecTemplateSpecVolumeOutputReference
type jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() JobV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns JobV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AzureDisk() JobV1SpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns JobV1SpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AzureDiskInput() *JobV1SpecTemplateSpecVolumeAzureDisk {
	var returns *JobV1SpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AzureFile() JobV1SpecTemplateSpecVolumeAzureFileOutputReference {
	var returns JobV1SpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) AzureFileInput() *JobV1SpecTemplateSpecVolumeAzureFile {
	var returns *JobV1SpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) CephFs() JobV1SpecTemplateSpecVolumeCephFsOutputReference {
	var returns JobV1SpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) CephFsInput() *JobV1SpecTemplateSpecVolumeCephFs {
	var returns *JobV1SpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Cinder() JobV1SpecTemplateSpecVolumeCinderOutputReference {
	var returns JobV1SpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) CinderInput() *JobV1SpecTemplateSpecVolumeCinder {
	var returns *JobV1SpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ConfigMap() JobV1SpecTemplateSpecVolumeConfigMapOutputReference {
	var returns JobV1SpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ConfigMapInput() *JobV1SpecTemplateSpecVolumeConfigMap {
	var returns *JobV1SpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Csi() JobV1SpecTemplateSpecVolumeCsiOutputReference {
	var returns JobV1SpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) CsiInput() *JobV1SpecTemplateSpecVolumeCsi {
	var returns *JobV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) DownwardApi() JobV1SpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns JobV1SpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) DownwardApiInput() *JobV1SpecTemplateSpecVolumeDownwardApi {
	var returns *JobV1SpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) EmptyDir() JobV1SpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns JobV1SpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) EmptyDirInput() *JobV1SpecTemplateSpecVolumeEmptyDir {
	var returns *JobV1SpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Fc() JobV1SpecTemplateSpecVolumeFcOutputReference {
	var returns JobV1SpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) FcInput() *JobV1SpecTemplateSpecVolumeFc {
	var returns *JobV1SpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) FlexVolume() JobV1SpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns JobV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *JobV1SpecTemplateSpecVolumeFlexVolume {
	var returns *JobV1SpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Flocker() JobV1SpecTemplateSpecVolumeFlockerOutputReference {
	var returns JobV1SpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) FlockerInput() *JobV1SpecTemplateSpecVolumeFlocker {
	var returns *JobV1SpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GcePersistentDisk() JobV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns JobV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *JobV1SpecTemplateSpecVolumeGcePersistentDisk {
	var returns *JobV1SpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GitRepo() JobV1SpecTemplateSpecVolumeGitRepoOutputReference {
	var returns JobV1SpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GitRepoInput() *JobV1SpecTemplateSpecVolumeGitRepo {
	var returns *JobV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Glusterfs() JobV1SpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns JobV1SpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GlusterfsInput() *JobV1SpecTemplateSpecVolumeGlusterfs {
	var returns *JobV1SpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) HostPath() JobV1SpecTemplateSpecVolumeHostPathOutputReference {
	var returns JobV1SpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) HostPathInput() *JobV1SpecTemplateSpecVolumeHostPath {
	var returns *JobV1SpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Iscsi() JobV1SpecTemplateSpecVolumeIscsiOutputReference {
	var returns JobV1SpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) IscsiInput() *JobV1SpecTemplateSpecVolumeIscsi {
	var returns *JobV1SpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Local() JobV1SpecTemplateSpecVolumeLocalOutputReference {
	var returns JobV1SpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) LocalInput() *JobV1SpecTemplateSpecVolumeLocal {
	var returns *JobV1SpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Nfs() JobV1SpecTemplateSpecVolumeNfsOutputReference {
	var returns JobV1SpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) NfsInput() *JobV1SpecTemplateSpecVolumeNfs {
	var returns *JobV1SpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() JobV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns JobV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *JobV1SpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *JobV1SpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() JobV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns JobV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *JobV1SpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *JobV1SpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Projected() JobV1SpecTemplateSpecVolumeProjectedList {
	var returns JobV1SpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Quobyte() JobV1SpecTemplateSpecVolumeQuobyteOutputReference {
	var returns JobV1SpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) QuobyteInput() *JobV1SpecTemplateSpecVolumeQuobyte {
	var returns *JobV1SpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Rbd() JobV1SpecTemplateSpecVolumeRbdOutputReference {
	var returns JobV1SpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) RbdInput() *JobV1SpecTemplateSpecVolumeRbd {
	var returns *JobV1SpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Secret() JobV1SpecTemplateSpecVolumeSecretOutputReference {
	var returns JobV1SpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) SecretInput() *JobV1SpecTemplateSpecVolumeSecret {
	var returns *JobV1SpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) VsphereVolume() JobV1SpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns JobV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *JobV1SpecTemplateSpecVolumeVsphereVolume {
	var returns *JobV1SpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewJobV1SpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobV1SpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecVolumeOutputReference_Override(j JobV1SpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *JobV1SpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := j.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *JobV1SpecTemplateSpecVolumeAzureDisk) {
	if err := j.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutAzureFile(value *JobV1SpecTemplateSpecVolumeAzureFile) {
	if err := j.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutCephFs(value *JobV1SpecTemplateSpecVolumeCephFs) {
	if err := j.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCephFs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutCinder(value *JobV1SpecTemplateSpecVolumeCinder) {
	if err := j.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCinder",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutConfigMap(value *JobV1SpecTemplateSpecVolumeConfigMap) {
	if err := j.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutCsi(value *JobV1SpecTemplateSpecVolumeCsi) {
	if err := j.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCsi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *JobV1SpecTemplateSpecVolumeDownwardApi) {
	if err := j.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *JobV1SpecTemplateSpecVolumeEmptyDir) {
	if err := j.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutFc(value *JobV1SpecTemplateSpecVolumeFc) {
	if err := j.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFc",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *JobV1SpecTemplateSpecVolumeFlexVolume) {
	if err := j.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutFlocker(value *JobV1SpecTemplateSpecVolumeFlocker) {
	if err := j.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFlocker",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *JobV1SpecTemplateSpecVolumeGcePersistentDisk) {
	if err := j.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutGitRepo(value *JobV1SpecTemplateSpecVolumeGitRepo) {
	if err := j.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *JobV1SpecTemplateSpecVolumeGlusterfs) {
	if err := j.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutHostPath(value *JobV1SpecTemplateSpecVolumeHostPath) {
	if err := j.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHostPath",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutIscsi(value *JobV1SpecTemplateSpecVolumeIscsi) {
	if err := j.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putIscsi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutLocal(value *JobV1SpecTemplateSpecVolumeLocal) {
	if err := j.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLocal",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutNfs(value *JobV1SpecTemplateSpecVolumeNfs) {
	if err := j.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *JobV1SpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := j.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *JobV1SpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := j.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := j.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putProjected",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutQuobyte(value *JobV1SpecTemplateSpecVolumeQuobyte) {
	if err := j.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutRbd(value *JobV1SpecTemplateSpecVolumeRbd) {
	if err := j.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRbd",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutSecret(value *JobV1SpecTemplateSpecVolumeSecret) {
	if err := j.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecret",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *JobV1SpecTemplateSpecVolumeVsphereVolume) {
	if err := j.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		j,
		"resetCephFs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		j,
		"resetCinder",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		j,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		j,
		"resetCsi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		j,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		j,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		j,
		"resetFc",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		j,
		"resetFlocker",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		j,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		j,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		j,
		"resetHostPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		j,
		"resetIscsi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		j,
		"resetLocal",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		j,
		"resetName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		j,
		"resetNfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		j,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		j,
		"resetProjected",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		j,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		j,
		"resetRbd",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		j,
		"resetSecret",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

