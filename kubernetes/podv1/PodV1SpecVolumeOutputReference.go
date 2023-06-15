package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v7/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v7/podv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodV1SpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() PodV1SpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *PodV1SpecVolumeAwsElasticBlockStore
	AzureDisk() PodV1SpecVolumeAzureDiskOutputReference
	AzureDiskInput() *PodV1SpecVolumeAzureDisk
	AzureFile() PodV1SpecVolumeAzureFileOutputReference
	AzureFileInput() *PodV1SpecVolumeAzureFile
	CephFs() PodV1SpecVolumeCephFsOutputReference
	CephFsInput() *PodV1SpecVolumeCephFs
	Cinder() PodV1SpecVolumeCinderOutputReference
	CinderInput() *PodV1SpecVolumeCinder
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
	ConfigMap() PodV1SpecVolumeConfigMapOutputReference
	ConfigMapInput() *PodV1SpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() PodV1SpecVolumeCsiOutputReference
	CsiInput() *PodV1SpecVolumeCsi
	DownwardApi() PodV1SpecVolumeDownwardApiOutputReference
	DownwardApiInput() *PodV1SpecVolumeDownwardApi
	EmptyDir() PodV1SpecVolumeEmptyDirOutputReference
	EmptyDirInput() *PodV1SpecVolumeEmptyDir
	Fc() PodV1SpecVolumeFcOutputReference
	FcInput() *PodV1SpecVolumeFc
	FlexVolume() PodV1SpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *PodV1SpecVolumeFlexVolume
	Flocker() PodV1SpecVolumeFlockerOutputReference
	FlockerInput() *PodV1SpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() PodV1SpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *PodV1SpecVolumeGcePersistentDisk
	GitRepo() PodV1SpecVolumeGitRepoOutputReference
	GitRepoInput() *PodV1SpecVolumeGitRepo
	Glusterfs() PodV1SpecVolumeGlusterfsOutputReference
	GlusterfsInput() *PodV1SpecVolumeGlusterfs
	HostPath() PodV1SpecVolumeHostPathOutputReference
	HostPathInput() *PodV1SpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() PodV1SpecVolumeIscsiOutputReference
	IscsiInput() *PodV1SpecVolumeIscsi
	Local() PodV1SpecVolumeLocalOutputReference
	LocalInput() *PodV1SpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() PodV1SpecVolumeNfsOutputReference
	NfsInput() *PodV1SpecVolumeNfs
	PersistentVolumeClaim() PodV1SpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *PodV1SpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() PodV1SpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *PodV1SpecVolumePhotonPersistentDisk
	Projected() PodV1SpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() PodV1SpecVolumeQuobyteOutputReference
	QuobyteInput() *PodV1SpecVolumeQuobyte
	Rbd() PodV1SpecVolumeRbdOutputReference
	RbdInput() *PodV1SpecVolumeRbd
	Secret() PodV1SpecVolumeSecretOutputReference
	SecretInput() *PodV1SpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() PodV1SpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *PodV1SpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *PodV1SpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *PodV1SpecVolumeAzureDisk)
	PutAzureFile(value *PodV1SpecVolumeAzureFile)
	PutCephFs(value *PodV1SpecVolumeCephFs)
	PutCinder(value *PodV1SpecVolumeCinder)
	PutConfigMap(value *PodV1SpecVolumeConfigMap)
	PutCsi(value *PodV1SpecVolumeCsi)
	PutDownwardApi(value *PodV1SpecVolumeDownwardApi)
	PutEmptyDir(value *PodV1SpecVolumeEmptyDir)
	PutFc(value *PodV1SpecVolumeFc)
	PutFlexVolume(value *PodV1SpecVolumeFlexVolume)
	PutFlocker(value *PodV1SpecVolumeFlocker)
	PutGcePersistentDisk(value *PodV1SpecVolumeGcePersistentDisk)
	PutGitRepo(value *PodV1SpecVolumeGitRepo)
	PutGlusterfs(value *PodV1SpecVolumeGlusterfs)
	PutHostPath(value *PodV1SpecVolumeHostPath)
	PutIscsi(value *PodV1SpecVolumeIscsi)
	PutLocal(value *PodV1SpecVolumeLocal)
	PutNfs(value *PodV1SpecVolumeNfs)
	PutPersistentVolumeClaim(value *PodV1SpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *PodV1SpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *PodV1SpecVolumeQuobyte)
	PutRbd(value *PodV1SpecVolumeRbd)
	PutSecret(value *PodV1SpecVolumeSecret)
	PutVsphereVolume(value *PodV1SpecVolumeVsphereVolume)
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

// The jsii proxy struct for PodV1SpecVolumeOutputReference
type jsiiProxy_PodV1SpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AwsElasticBlockStore() PodV1SpecVolumeAwsElasticBlockStoreOutputReference {
	var returns PodV1SpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AwsElasticBlockStoreInput() *PodV1SpecVolumeAwsElasticBlockStore {
	var returns *PodV1SpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AzureDisk() PodV1SpecVolumeAzureDiskOutputReference {
	var returns PodV1SpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AzureDiskInput() *PodV1SpecVolumeAzureDisk {
	var returns *PodV1SpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AzureFile() PodV1SpecVolumeAzureFileOutputReference {
	var returns PodV1SpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) AzureFileInput() *PodV1SpecVolumeAzureFile {
	var returns *PodV1SpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) CephFs() PodV1SpecVolumeCephFsOutputReference {
	var returns PodV1SpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) CephFsInput() *PodV1SpecVolumeCephFs {
	var returns *PodV1SpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Cinder() PodV1SpecVolumeCinderOutputReference {
	var returns PodV1SpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) CinderInput() *PodV1SpecVolumeCinder {
	var returns *PodV1SpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) ConfigMap() PodV1SpecVolumeConfigMapOutputReference {
	var returns PodV1SpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) ConfigMapInput() *PodV1SpecVolumeConfigMap {
	var returns *PodV1SpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Csi() PodV1SpecVolumeCsiOutputReference {
	var returns PodV1SpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) CsiInput() *PodV1SpecVolumeCsi {
	var returns *PodV1SpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) DownwardApi() PodV1SpecVolumeDownwardApiOutputReference {
	var returns PodV1SpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) DownwardApiInput() *PodV1SpecVolumeDownwardApi {
	var returns *PodV1SpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) EmptyDir() PodV1SpecVolumeEmptyDirOutputReference {
	var returns PodV1SpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) EmptyDirInput() *PodV1SpecVolumeEmptyDir {
	var returns *PodV1SpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Fc() PodV1SpecVolumeFcOutputReference {
	var returns PodV1SpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) FcInput() *PodV1SpecVolumeFc {
	var returns *PodV1SpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) FlexVolume() PodV1SpecVolumeFlexVolumeOutputReference {
	var returns PodV1SpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) FlexVolumeInput() *PodV1SpecVolumeFlexVolume {
	var returns *PodV1SpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Flocker() PodV1SpecVolumeFlockerOutputReference {
	var returns PodV1SpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) FlockerInput() *PodV1SpecVolumeFlocker {
	var returns *PodV1SpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) GcePersistentDisk() PodV1SpecVolumeGcePersistentDiskOutputReference {
	var returns PodV1SpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) GcePersistentDiskInput() *PodV1SpecVolumeGcePersistentDisk {
	var returns *PodV1SpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) GitRepo() PodV1SpecVolumeGitRepoOutputReference {
	var returns PodV1SpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) GitRepoInput() *PodV1SpecVolumeGitRepo {
	var returns *PodV1SpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Glusterfs() PodV1SpecVolumeGlusterfsOutputReference {
	var returns PodV1SpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) GlusterfsInput() *PodV1SpecVolumeGlusterfs {
	var returns *PodV1SpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) HostPath() PodV1SpecVolumeHostPathOutputReference {
	var returns PodV1SpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) HostPathInput() *PodV1SpecVolumeHostPath {
	var returns *PodV1SpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Iscsi() PodV1SpecVolumeIscsiOutputReference {
	var returns PodV1SpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) IscsiInput() *PodV1SpecVolumeIscsi {
	var returns *PodV1SpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Local() PodV1SpecVolumeLocalOutputReference {
	var returns PodV1SpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) LocalInput() *PodV1SpecVolumeLocal {
	var returns *PodV1SpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Nfs() PodV1SpecVolumeNfsOutputReference {
	var returns PodV1SpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) NfsInput() *PodV1SpecVolumeNfs {
	var returns *PodV1SpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) PersistentVolumeClaim() PodV1SpecVolumePersistentVolumeClaimOutputReference {
	var returns PodV1SpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) PersistentVolumeClaimInput() *PodV1SpecVolumePersistentVolumeClaim {
	var returns *PodV1SpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) PhotonPersistentDisk() PodV1SpecVolumePhotonPersistentDiskOutputReference {
	var returns PodV1SpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) PhotonPersistentDiskInput() *PodV1SpecVolumePhotonPersistentDisk {
	var returns *PodV1SpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Projected() PodV1SpecVolumeProjectedList {
	var returns PodV1SpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Quobyte() PodV1SpecVolumeQuobyteOutputReference {
	var returns PodV1SpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) QuobyteInput() *PodV1SpecVolumeQuobyte {
	var returns *PodV1SpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Rbd() PodV1SpecVolumeRbdOutputReference {
	var returns PodV1SpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) RbdInput() *PodV1SpecVolumeRbd {
	var returns *PodV1SpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) Secret() PodV1SpecVolumeSecretOutputReference {
	var returns PodV1SpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) SecretInput() *PodV1SpecVolumeSecret {
	var returns *PodV1SpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) VsphereVolume() PodV1SpecVolumeVsphereVolumeOutputReference {
	var returns PodV1SpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference) VsphereVolumeInput() *PodV1SpecVolumeVsphereVolume {
	var returns *PodV1SpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewPodV1SpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PodV1SpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPodV1SpecVolumeOutputReference_Override(p PodV1SpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutAwsElasticBlockStore(value *PodV1SpecVolumeAwsElasticBlockStore) {
	if err := p.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutAzureDisk(value *PodV1SpecVolumeAzureDisk) {
	if err := p.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutAzureFile(value *PodV1SpecVolumeAzureFile) {
	if err := p.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutCephFs(value *PodV1SpecVolumeCephFs) {
	if err := p.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCephFs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutCinder(value *PodV1SpecVolumeCinder) {
	if err := p.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCinder",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutConfigMap(value *PodV1SpecVolumeConfigMap) {
	if err := p.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutCsi(value *PodV1SpecVolumeCsi) {
	if err := p.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutDownwardApi(value *PodV1SpecVolumeDownwardApi) {
	if err := p.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutEmptyDir(value *PodV1SpecVolumeEmptyDir) {
	if err := p.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutFc(value *PodV1SpecVolumeFc) {
	if err := p.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFc",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutFlexVolume(value *PodV1SpecVolumeFlexVolume) {
	if err := p.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutFlocker(value *PodV1SpecVolumeFlocker) {
	if err := p.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlocker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutGcePersistentDisk(value *PodV1SpecVolumeGcePersistentDisk) {
	if err := p.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutGitRepo(value *PodV1SpecVolumeGitRepo) {
	if err := p.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutGlusterfs(value *PodV1SpecVolumeGlusterfs) {
	if err := p.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutHostPath(value *PodV1SpecVolumeHostPath) {
	if err := p.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPath",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutIscsi(value *PodV1SpecVolumeIscsi) {
	if err := p.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIscsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutLocal(value *PodV1SpecVolumeLocal) {
	if err := p.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLocal",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutNfs(value *PodV1SpecVolumeNfs) {
	if err := p.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutPersistentVolumeClaim(value *PodV1SpecVolumePersistentVolumeClaim) {
	if err := p.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutPhotonPersistentDisk(value *PodV1SpecVolumePhotonPersistentDisk) {
	if err := p.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := p.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProjected",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutQuobyte(value *PodV1SpecVolumeQuobyte) {
	if err := p.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutRbd(value *PodV1SpecVolumeRbd) {
	if err := p.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRbd",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutSecret(value *PodV1SpecVolumeSecret) {
	if err := p.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecret",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) PutVsphereVolume(value *PodV1SpecVolumeVsphereVolume) {
	if err := p.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		p,
		"resetCephFs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		p,
		"resetCinder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		p,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		p,
		"resetCsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		p,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		p,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		p,
		"resetFc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		p,
		"resetFlocker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		p,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		p,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		p,
		"resetIscsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		p,
		"resetLocal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		p,
		"resetNfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		p,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		p,
		"resetProjected",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		p,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		p,
		"resetRbd",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		p,
		"resetSecret",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

