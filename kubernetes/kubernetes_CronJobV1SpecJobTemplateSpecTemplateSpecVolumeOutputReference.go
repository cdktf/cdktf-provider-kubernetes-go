// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/jsii"

	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk
	AzureFile() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile
	CephFs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs
	Cinder() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder
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
	ConfigMap() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi
	DownwardApi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi
	EmptyDir() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir
	Fc() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFcOutputReference
	FcInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc
	FlexVolume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume
	Flocker() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo
	Glusterfs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs
	HostPath() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi
	Local() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte
	Rbd() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd
	Secret() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs)
	PutCinder(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir)
	PutFc(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi)
	PutLocal(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal)
	PutNfs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd)
	PutSecret(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureFile() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) AzureFileInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) CephFs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) CephFsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Cinder() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) CinderInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ConfigMap() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Csi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) CsiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) DownwardApi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) EmptyDir() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Fc() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFcOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) FcInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlexVolume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Flocker() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) FlockerInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GitRepo() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GitRepoInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Glusterfs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) HostPath() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) HostPathInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Iscsi() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) IscsiInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Local() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) LocalInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Nfs() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) NfsInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Projected() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Quobyte() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) QuobyteInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Rbd() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) RbdInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Secret() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) SecretInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) VsphereVolume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := c.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureDisk) {
	if err := c.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile) {
	if err := c.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCephFs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCephFs) {
	if err := c.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCephFs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCinder(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCinder) {
	if err := c.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCinder",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeConfigMap) {
	if err := c.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutCsi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeCsi) {
	if err := c.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCsi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeDownwardApi) {
	if err := c.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEmptyDir) {
	if err := c.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFc(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFc) {
	if err := c.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlexVolume) {
	if err := c.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutFlocker(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeFlocker) {
	if err := c.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFlocker",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := c.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGitRepo) {
	if err := c.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeGlusterfs) {
	if err := c.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutHostPath(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath) {
	if err := c.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutIscsi(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi) {
	if err := c.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIscsi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutLocal(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeLocal) {
	if err := c.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocal",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutNfs(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeNfs) {
	if err := c.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := c.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := c.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := c.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProjected",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeQuobyte) {
	if err := c.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutRbd(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeRbd) {
	if err := c.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRbd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutSecret(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeSecret) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeVsphereVolume) {
	if err := c.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		c,
		"resetCephFs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		c,
		"resetCinder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		c,
		"resetCsi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		c,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		c,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		c,
		"resetFc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		c,
		"resetFlocker",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		c,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		c,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		c,
		"resetHostPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		c,
		"resetIscsi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		c,
		"resetLocal",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		c,
		"resetNfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		c,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		c,
		"resetProjected",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		c,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		c,
		"resetRbd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

