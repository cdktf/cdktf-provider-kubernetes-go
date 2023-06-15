package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v7/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v7/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() JobSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *JobSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() JobSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *JobSpecTemplateSpecVolumeAzureDisk
	AzureFile() JobSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *JobSpecTemplateSpecVolumeAzureFile
	CephFs() JobSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *JobSpecTemplateSpecVolumeCephFs
	Cinder() JobSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *JobSpecTemplateSpecVolumeCinder
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
	ConfigMap() JobSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *JobSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() JobSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *JobSpecTemplateSpecVolumeCsi
	DownwardApi() JobSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *JobSpecTemplateSpecVolumeDownwardApi
	EmptyDir() JobSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *JobSpecTemplateSpecVolumeEmptyDir
	Fc() JobSpecTemplateSpecVolumeFcOutputReference
	FcInput() *JobSpecTemplateSpecVolumeFc
	FlexVolume() JobSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *JobSpecTemplateSpecVolumeFlexVolume
	Flocker() JobSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *JobSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() JobSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *JobSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() JobSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *JobSpecTemplateSpecVolumeGitRepo
	Glusterfs() JobSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *JobSpecTemplateSpecVolumeGlusterfs
	HostPath() JobSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *JobSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() JobSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *JobSpecTemplateSpecVolumeIscsi
	Local() JobSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *JobSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() JobSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *JobSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() JobSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *JobSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() JobSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *JobSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() JobSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() JobSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *JobSpecTemplateSpecVolumeQuobyte
	Rbd() JobSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *JobSpecTemplateSpecVolumeRbd
	Secret() JobSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *JobSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() JobSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *JobSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *JobSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *JobSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *JobSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *JobSpecTemplateSpecVolumeCephFs)
	PutCinder(value *JobSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *JobSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *JobSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *JobSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *JobSpecTemplateSpecVolumeEmptyDir)
	PutFc(value *JobSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *JobSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *JobSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *JobSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *JobSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *JobSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *JobSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *JobSpecTemplateSpecVolumeIscsi)
	PutLocal(value *JobSpecTemplateSpecVolumeLocal)
	PutNfs(value *JobSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *JobSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *JobSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *JobSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *JobSpecTemplateSpecVolumeRbd)
	PutSecret(value *JobSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *JobSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for JobSpecTemplateSpecVolumeOutputReference
type jsiiProxy_JobSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() JobSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns JobSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *JobSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *JobSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AzureDisk() JobSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns JobSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *JobSpecTemplateSpecVolumeAzureDisk {
	var returns *JobSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AzureFile() JobSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns JobSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) AzureFileInput() *JobSpecTemplateSpecVolumeAzureFile {
	var returns *JobSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) CephFs() JobSpecTemplateSpecVolumeCephFsOutputReference {
	var returns JobSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) CephFsInput() *JobSpecTemplateSpecVolumeCephFs {
	var returns *JobSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Cinder() JobSpecTemplateSpecVolumeCinderOutputReference {
	var returns JobSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) CinderInput() *JobSpecTemplateSpecVolumeCinder {
	var returns *JobSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ConfigMap() JobSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns JobSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *JobSpecTemplateSpecVolumeConfigMap {
	var returns *JobSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Csi() JobSpecTemplateSpecVolumeCsiOutputReference {
	var returns JobSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) CsiInput() *JobSpecTemplateSpecVolumeCsi {
	var returns *JobSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) DownwardApi() JobSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns JobSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *JobSpecTemplateSpecVolumeDownwardApi {
	var returns *JobSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) EmptyDir() JobSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns JobSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *JobSpecTemplateSpecVolumeEmptyDir {
	var returns *JobSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Fc() JobSpecTemplateSpecVolumeFcOutputReference {
	var returns JobSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) FcInput() *JobSpecTemplateSpecVolumeFc {
	var returns *JobSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) FlexVolume() JobSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns JobSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *JobSpecTemplateSpecVolumeFlexVolume {
	var returns *JobSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Flocker() JobSpecTemplateSpecVolumeFlockerOutputReference {
	var returns JobSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) FlockerInput() *JobSpecTemplateSpecVolumeFlocker {
	var returns *JobSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() JobSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns JobSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *JobSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *JobSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GitRepo() JobSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns JobSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GitRepoInput() *JobSpecTemplateSpecVolumeGitRepo {
	var returns *JobSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Glusterfs() JobSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns JobSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *JobSpecTemplateSpecVolumeGlusterfs {
	var returns *JobSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) HostPath() JobSpecTemplateSpecVolumeHostPathOutputReference {
	var returns JobSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) HostPathInput() *JobSpecTemplateSpecVolumeHostPath {
	var returns *JobSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Iscsi() JobSpecTemplateSpecVolumeIscsiOutputReference {
	var returns JobSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) IscsiInput() *JobSpecTemplateSpecVolumeIscsi {
	var returns *JobSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Local() JobSpecTemplateSpecVolumeLocalOutputReference {
	var returns JobSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) LocalInput() *JobSpecTemplateSpecVolumeLocal {
	var returns *JobSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Nfs() JobSpecTemplateSpecVolumeNfsOutputReference {
	var returns JobSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) NfsInput() *JobSpecTemplateSpecVolumeNfs {
	var returns *JobSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() JobSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns JobSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *JobSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *JobSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() JobSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns JobSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *JobSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *JobSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Projected() JobSpecTemplateSpecVolumeProjectedList {
	var returns JobSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Quobyte() JobSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns JobSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) QuobyteInput() *JobSpecTemplateSpecVolumeQuobyte {
	var returns *JobSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Rbd() JobSpecTemplateSpecVolumeRbdOutputReference {
	var returns JobSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) RbdInput() *JobSpecTemplateSpecVolumeRbd {
	var returns *JobSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Secret() JobSpecTemplateSpecVolumeSecretOutputReference {
	var returns JobSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) SecretInput() *JobSpecTemplateSpecVolumeSecret {
	var returns *JobSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) VsphereVolume() JobSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns JobSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *JobSpecTemplateSpecVolumeVsphereVolume {
	var returns *JobSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecVolumeOutputReference_Override(j JobSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *JobSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := j.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *JobSpecTemplateSpecVolumeAzureDisk) {
	if err := j.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *JobSpecTemplateSpecVolumeAzureFile) {
	if err := j.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutCephFs(value *JobSpecTemplateSpecVolumeCephFs) {
	if err := j.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCephFs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutCinder(value *JobSpecTemplateSpecVolumeCinder) {
	if err := j.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCinder",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *JobSpecTemplateSpecVolumeConfigMap) {
	if err := j.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutCsi(value *JobSpecTemplateSpecVolumeCsi) {
	if err := j.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCsi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *JobSpecTemplateSpecVolumeDownwardApi) {
	if err := j.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *JobSpecTemplateSpecVolumeEmptyDir) {
	if err := j.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutFc(value *JobSpecTemplateSpecVolumeFc) {
	if err := j.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFc",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *JobSpecTemplateSpecVolumeFlexVolume) {
	if err := j.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutFlocker(value *JobSpecTemplateSpecVolumeFlocker) {
	if err := j.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFlocker",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *JobSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := j.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *JobSpecTemplateSpecVolumeGitRepo) {
	if err := j.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *JobSpecTemplateSpecVolumeGlusterfs) {
	if err := j.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutHostPath(value *JobSpecTemplateSpecVolumeHostPath) {
	if err := j.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHostPath",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutIscsi(value *JobSpecTemplateSpecVolumeIscsi) {
	if err := j.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putIscsi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutLocal(value *JobSpecTemplateSpecVolumeLocal) {
	if err := j.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLocal",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutNfs(value *JobSpecTemplateSpecVolumeNfs) {
	if err := j.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putNfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *JobSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := j.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *JobSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := j.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := j.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putProjected",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *JobSpecTemplateSpecVolumeQuobyte) {
	if err := j.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutRbd(value *JobSpecTemplateSpecVolumeRbd) {
	if err := j.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putRbd",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutSecret(value *JobSpecTemplateSpecVolumeSecret) {
	if err := j.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecret",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *JobSpecTemplateSpecVolumeVsphereVolume) {
	if err := j.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		j,
		"resetCephFs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		j,
		"resetCinder",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		j,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		j,
		"resetCsi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		j,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		j,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		j,
		"resetFc",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		j,
		"resetFlocker",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		j,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		j,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		j,
		"resetHostPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		j,
		"resetIscsi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		j,
		"resetLocal",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		j,
		"resetName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		j,
		"resetNfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		j,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		j,
		"resetProjected",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		j,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		j,
		"resetRbd",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		j,
		"resetSecret",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

