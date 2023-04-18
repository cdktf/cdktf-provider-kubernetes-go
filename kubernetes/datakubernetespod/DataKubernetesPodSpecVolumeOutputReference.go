package datakubernetespod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v6/datakubernetespod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesPodSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DataKubernetesPodSpecVolumeAwsElasticBlockStoreList
	AzureDisk() DataKubernetesPodSpecVolumeAzureDiskList
	AzureFile() DataKubernetesPodSpecVolumeAzureFileList
	CephFs() DataKubernetesPodSpecVolumeCephFsList
	Cinder() DataKubernetesPodSpecVolumeCinderList
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
	ConfigMap() DataKubernetesPodSpecVolumeConfigMapList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DataKubernetesPodSpecVolumeCsiList
	DownwardApi() DataKubernetesPodSpecVolumeDownwardApiList
	EmptyDir() DataKubernetesPodSpecVolumeEmptyDirList
	Fc() DataKubernetesPodSpecVolumeFcList
	FlexVolume() DataKubernetesPodSpecVolumeFlexVolumeList
	Flocker() DataKubernetesPodSpecVolumeFlockerList
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DataKubernetesPodSpecVolumeGcePersistentDiskList
	GitRepo() DataKubernetesPodSpecVolumeGitRepoList
	Glusterfs() DataKubernetesPodSpecVolumeGlusterfsList
	HostPath() DataKubernetesPodSpecVolumeHostPathList
	InternalValue() *DataKubernetesPodSpecVolume
	SetInternalValue(val *DataKubernetesPodSpecVolume)
	Iscsi() DataKubernetesPodSpecVolumeIscsiList
	Local() DataKubernetesPodSpecVolumeLocalList
	Name() *string
	Nfs() DataKubernetesPodSpecVolumeNfsList
	PersistentVolumeClaim() DataKubernetesPodSpecVolumePersistentVolumeClaimList
	PhotonPersistentDisk() DataKubernetesPodSpecVolumePhotonPersistentDiskList
	Projected() DataKubernetesPodSpecVolumeProjectedList
	Quobyte() DataKubernetesPodSpecVolumeQuobyteList
	Rbd() DataKubernetesPodSpecVolumeRbdList
	Secret() DataKubernetesPodSpecVolumeSecretList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DataKubernetesPodSpecVolumeVsphereVolumeList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataKubernetesPodSpecVolumeOutputReference
type jsiiProxy_DataKubernetesPodSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) AwsElasticBlockStore() DataKubernetesPodSpecVolumeAwsElasticBlockStoreList {
	var returns DataKubernetesPodSpecVolumeAwsElasticBlockStoreList
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) AzureDisk() DataKubernetesPodSpecVolumeAzureDiskList {
	var returns DataKubernetesPodSpecVolumeAzureDiskList
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) AzureFile() DataKubernetesPodSpecVolumeAzureFileList {
	var returns DataKubernetesPodSpecVolumeAzureFileList
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) CephFs() DataKubernetesPodSpecVolumeCephFsList {
	var returns DataKubernetesPodSpecVolumeCephFsList
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Cinder() DataKubernetesPodSpecVolumeCinderList {
	var returns DataKubernetesPodSpecVolumeCinderList
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) ConfigMap() DataKubernetesPodSpecVolumeConfigMapList {
	var returns DataKubernetesPodSpecVolumeConfigMapList
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Csi() DataKubernetesPodSpecVolumeCsiList {
	var returns DataKubernetesPodSpecVolumeCsiList
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) DownwardApi() DataKubernetesPodSpecVolumeDownwardApiList {
	var returns DataKubernetesPodSpecVolumeDownwardApiList
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) EmptyDir() DataKubernetesPodSpecVolumeEmptyDirList {
	var returns DataKubernetesPodSpecVolumeEmptyDirList
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Fc() DataKubernetesPodSpecVolumeFcList {
	var returns DataKubernetesPodSpecVolumeFcList
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) FlexVolume() DataKubernetesPodSpecVolumeFlexVolumeList {
	var returns DataKubernetesPodSpecVolumeFlexVolumeList
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Flocker() DataKubernetesPodSpecVolumeFlockerList {
	var returns DataKubernetesPodSpecVolumeFlockerList
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GcePersistentDisk() DataKubernetesPodSpecVolumeGcePersistentDiskList {
	var returns DataKubernetesPodSpecVolumeGcePersistentDiskList
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GitRepo() DataKubernetesPodSpecVolumeGitRepoList {
	var returns DataKubernetesPodSpecVolumeGitRepoList
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Glusterfs() DataKubernetesPodSpecVolumeGlusterfsList {
	var returns DataKubernetesPodSpecVolumeGlusterfsList
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) HostPath() DataKubernetesPodSpecVolumeHostPathList {
	var returns DataKubernetesPodSpecVolumeHostPathList
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) InternalValue() *DataKubernetesPodSpecVolume {
	var returns *DataKubernetesPodSpecVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Iscsi() DataKubernetesPodSpecVolumeIscsiList {
	var returns DataKubernetesPodSpecVolumeIscsiList
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Local() DataKubernetesPodSpecVolumeLocalList {
	var returns DataKubernetesPodSpecVolumeLocalList
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Nfs() DataKubernetesPodSpecVolumeNfsList {
	var returns DataKubernetesPodSpecVolumeNfsList
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) PersistentVolumeClaim() DataKubernetesPodSpecVolumePersistentVolumeClaimList {
	var returns DataKubernetesPodSpecVolumePersistentVolumeClaimList
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) PhotonPersistentDisk() DataKubernetesPodSpecVolumePhotonPersistentDiskList {
	var returns DataKubernetesPodSpecVolumePhotonPersistentDiskList
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Projected() DataKubernetesPodSpecVolumeProjectedList {
	var returns DataKubernetesPodSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Quobyte() DataKubernetesPodSpecVolumeQuobyteList {
	var returns DataKubernetesPodSpecVolumeQuobyteList
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Rbd() DataKubernetesPodSpecVolumeRbdList {
	var returns DataKubernetesPodSpecVolumeRbdList
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Secret() DataKubernetesPodSpecVolumeSecretList {
	var returns DataKubernetesPodSpecVolumeSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) VsphereVolume() DataKubernetesPodSpecVolumeVsphereVolumeList {
	var returns DataKubernetesPodSpecVolumeVsphereVolumeList
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}


func NewDataKubernetesPodSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataKubernetesPodSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDataKubernetesPodSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataKubernetesPodSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataKubernetesPodSpecVolumeOutputReference_Override(d DataKubernetesPodSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.dataKubernetesPod.DataKubernetesPodSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference)SetInternalValue(val *DataKubernetesPodSpecVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataKubernetesPodSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

