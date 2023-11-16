// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of the specified volume
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20221101preview:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Volume resource
type LookupVolumeResult struct {
	// Actual throughput in MiB/s for auto qosType volumes calculated based on size and serviceLevel
	ActualThroughputMibps float64 `pulumi:"actualThroughputMibps"`
	// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
	AvsDataStore *string `pulumi:"avsDataStore"`
	// UUID v4 or resource identifier used to identify the Backup.
	BackupId *string `pulumi:"backupId"`
	// Unique Baremetal Tenant Identifier.
	BaremetalTenantId string `pulumi:"baremetalTenantId"`
	// Pool Resource Id used in case of creating a volume through volume group
	CapacityPoolResourceId *string `pulumi:"capacityPoolResourceId"`
	// When a volume is being restored from another volume's snapshot, will show the percentage completion of this cloning process. When this value is empty/null there is no cloning process currently happening on this volume. This value will update every 5 minutes during cloning.
	CloneProgress int `pulumi:"cloneProgress"`
	// Specifies whether Cool Access(tiering) is enabled for the volume.
	CoolAccess *bool `pulumi:"coolAccess"`
	// Specifies the number of days after which data that is not accessed by clients will be tiered.
	CoolnessPeriod *int `pulumi:"coolnessPeriod"`
	// A unique file path for the volume. Used when creating mount targets
	CreationToken string `pulumi:"creationToken"`
	// DataProtection type volumes include an object containing details of the replication
	DataProtection *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	// Data store resource unique identifier
	DataStoreResourceId []string `pulumi:"dataStoreResourceId"`
	// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
	DefaultGroupQuotaInKiBs *float64 `pulumi:"defaultGroupQuotaInKiBs"`
	// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
	DefaultUserQuotaInKiBs *float64 `pulumi:"defaultUserQuotaInKiBs"`
	// If enabled (true) the snapshot the volume was created from will be automatically deleted after the volume create operation has finished.  Defaults to false
	DeleteBaseSnapshot *bool `pulumi:"deleteBaseSnapshot"`
	// Flag indicating whether subvolume operations are enabled on the volume
	EnableSubvolumes *string `pulumi:"enableSubvolumes"`
	// Specifies if the volume is encrypted or not. Only available on volumes created or updated after 2022-01-01.
	Encrypted bool `pulumi:"encrypted"`
	// Source of key used to encrypt data in volume. Applicable if NetApp account has encryption.keySource = 'Microsoft.KeyVault'. Possible values (case-insensitive) are: 'Microsoft.NetApp, Microsoft.KeyVault'
	EncryptionKeySource *string `pulumi:"encryptionKeySource"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Set of export policy rules
	ExportPolicy *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	// Flag indicating whether file access logs are enabled for the volume, based on active diagnostic settings present on the volume.
	FileAccessLogs string `pulumi:"fileAccessLogs"`
	// Unique FileSystem Identifier.
	FileSystemId string `pulumi:"fileSystemId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Specifies if default quota is enabled for the volume.
	IsDefaultQuotaEnabled *bool `pulumi:"isDefaultQuotaEnabled"`
	// Specifies whether volume is a Large Volume or Regular Volume.
	IsLargeVolume *bool `pulumi:"isLargeVolume"`
	// Restoring
	IsRestoring *bool `pulumi:"isRestoring"`
	// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
	KerberosEnabled *bool `pulumi:"kerberosEnabled"`
	// The resource ID of private endpoint for KeyVault. It must reside in the same VNET as the volume. Only applicable if encryptionKeySource = 'Microsoft.KeyVault'.
	KeyVaultPrivateEndpointResourceId *string `pulumi:"keyVaultPrivateEndpointResourceId"`
	// Specifies whether LDAP is enabled or not for a given NFS volume.
	LdapEnabled *bool `pulumi:"ldapEnabled"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maximum number of files allowed. Needs a service request in order to be changed. Only allowed to be changed if volume quota is more than 4TiB.
	MaximumNumberOfFiles float64 `pulumi:"maximumNumberOfFiles"`
	// List of mount targets
	MountTargets []MountTargetPropertiesResponse `pulumi:"mountTargets"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Basic network, or Standard features available to the volume.
	NetworkFeatures *string `pulumi:"networkFeatures"`
	// Network Sibling Set ID for the the group of volumes sharing networking resources.
	NetworkSiblingSetId string `pulumi:"networkSiblingSetId"`
	// Id of the snapshot or backup that the volume is restored from.
	OriginatingResourceId string `pulumi:"originatingResourceId"`
	// Application specific placement rules for the particular volume
	PlacementRules []PlacementKeyValuePairsResponse `pulumi:"placementRules"`
	// Set of protocol types, default NFSv3, CIFS for SMB protocol
	ProtocolTypes []string `pulumi:"protocolTypes"`
	// The availability zone where the volume is provisioned. This refers to the logical availability zone where the volume resides.
	ProvisionedAvailabilityZone string `pulumi:"provisionedAvailabilityZone"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Proximity placement group associated with the volume
	ProximityPlacementGroup *string `pulumi:"proximityPlacementGroup"`
	// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
	SecurityStyle *string `pulumi:"securityStyle"`
	// The service level of the file system
	ServiceLevel *string `pulumi:"serviceLevel"`
	// Enables access based enumeration share property for SMB Shares. Only applicable for SMB/DualProtocol volume
	SmbAccessBasedEnumeration *string `pulumi:"smbAccessBasedEnumeration"`
	// Enables continuously available share property for smb volume. Only applicable for SMB volume
	SmbContinuouslyAvailable *bool `pulumi:"smbContinuouslyAvailable"`
	// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
	SmbEncryption *bool `pulumi:"smbEncryption"`
	// Enables non browsable property for SMB Shares. Only applicable for SMB/DualProtocol volume
	SmbNonBrowsable *string `pulumi:"smbNonBrowsable"`
	// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (defaults to true).
	SnapshotDirectoryVisible *bool `pulumi:"snapshotDirectoryVisible"`
	// UUID v4 or resource identifier used to identify the Snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// Provides storage to network proximity information for the volume.
	StorageToNetworkProximity string `pulumi:"storageToNetworkProximity"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId string `pulumi:"subnetId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// T2 network information
	T2Network string `pulumi:"t2Network"`
	// Resource tags.
	Tags            map[string]string `pulumi:"tags"`
	ThroughputMibps *float64          `pulumi:"throughputMibps"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
	UnixPermissions *string `pulumi:"unixPermissions"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB, 500Tib for LargeVolume. Specified in bytes.
	UsageThreshold float64 `pulumi:"usageThreshold"`
	// Volume Group Name
	VolumeGroupName string `pulumi:"volumeGroupName"`
	// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
	VolumeSpecName *string `pulumi:"volumeSpecName"`
	// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
	VolumeType *string `pulumi:"volumeType"`
	// Availability Zone
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupVolumeResult
func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AvsDataStore == nil {
		avsDataStore_ := "Disabled"
		tmp.AvsDataStore = &avsDataStore_
	}
	if tmp.CoolAccess == nil {
		coolAccess_ := false
		tmp.CoolAccess = &coolAccess_
	}
	if tmp.DefaultGroupQuotaInKiBs == nil {
		defaultGroupQuotaInKiBs_ := 0.0
		tmp.DefaultGroupQuotaInKiBs = &defaultGroupQuotaInKiBs_
	}
	if tmp.DefaultUserQuotaInKiBs == nil {
		defaultUserQuotaInKiBs_ := 0.0
		tmp.DefaultUserQuotaInKiBs = &defaultUserQuotaInKiBs_
	}
	if tmp.EnableSubvolumes == nil {
		enableSubvolumes_ := "Disabled"
		tmp.EnableSubvolumes = &enableSubvolumes_
	}
	if tmp.EncryptionKeySource == nil {
		encryptionKeySource_ := "Microsoft.NetApp"
		tmp.EncryptionKeySource = &encryptionKeySource_
	}
	if utilities.IsZero(tmp.FileAccessLogs) {
		tmp.FileAccessLogs = "Disabled"
	}
	if tmp.IsDefaultQuotaEnabled == nil {
		isDefaultQuotaEnabled_ := false
		tmp.IsDefaultQuotaEnabled = &isDefaultQuotaEnabled_
	}
	if tmp.IsLargeVolume == nil {
		isLargeVolume_ := false
		tmp.IsLargeVolume = &isLargeVolume_
	}
	if tmp.KerberosEnabled == nil {
		kerberosEnabled_ := false
		tmp.KerberosEnabled = &kerberosEnabled_
	}
	if tmp.LdapEnabled == nil {
		ldapEnabled_ := false
		tmp.LdapEnabled = &ldapEnabled_
	}
	if tmp.NetworkFeatures == nil {
		networkFeatures_ := "Basic"
		tmp.NetworkFeatures = &networkFeatures_
	}
	if tmp.SecurityStyle == nil {
		securityStyle_ := "unix"
		tmp.SecurityStyle = &securityStyle_
	}
	if tmp.SmbContinuouslyAvailable == nil {
		smbContinuouslyAvailable_ := false
		tmp.SmbContinuouslyAvailable = &smbContinuouslyAvailable_
	}
	if tmp.SmbEncryption == nil {
		smbEncryption_ := false
		tmp.SmbEncryption = &smbEncryption_
	}
	if tmp.SnapshotDirectoryVisible == nil {
		snapshotDirectoryVisible_ := true
		tmp.SnapshotDirectoryVisible = &snapshotDirectoryVisible_
	}
	if tmp.UnixPermissions == nil {
		unixPermissions_ := "0770"
		tmp.UnixPermissions = &unixPermissions_
	}
	if utilities.IsZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

// Volume resource
type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

// Actual throughput in MiB/s for auto qosType volumes calculated based on size and serviceLevel
func (o LookupVolumeResultOutput) ActualThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.ActualThroughputMibps }).(pulumi.Float64Output)
}

// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
func (o LookupVolumeResultOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

// UUID v4 or resource identifier used to identify the Backup.
func (o LookupVolumeResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

// Unique Baremetal Tenant Identifier.
func (o LookupVolumeResultOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

// Pool Resource Id used in case of creating a volume through volume group
func (o LookupVolumeResultOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

// When a volume is being restored from another volume's snapshot, will show the percentage completion of this cloning process. When this value is empty/null there is no cloning process currently happening on this volume. This value will update every 5 minutes during cloning.
func (o LookupVolumeResultOutput) CloneProgress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeResult) int { return v.CloneProgress }).(pulumi.IntOutput)
}

// Specifies whether Cool Access(tiering) is enabled for the volume.
func (o LookupVolumeResultOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

// Specifies the number of days after which data that is not accessed by clients will be tiered.
func (o LookupVolumeResultOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *int { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

// A unique file path for the volume. Used when creating mount targets
func (o LookupVolumeResultOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.CreationToken }).(pulumi.StringOutput)
}

// DataProtection type volumes include an object containing details of the replication
func (o LookupVolumeResultOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseDataProtection { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
}

// Data store resource unique identifier
func (o LookupVolumeResultOutput) DataStoreResourceId() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.DataStoreResourceId }).(pulumi.StringArrayOutput)
}

// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
func (o LookupVolumeResultOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
func (o LookupVolumeResultOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

// If enabled (true) the snapshot the volume was created from will be automatically deleted after the volume create operation has finished.  Defaults to false
func (o LookupVolumeResultOutput) DeleteBaseSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.DeleteBaseSnapshot }).(pulumi.BoolPtrOutput)
}

// Flag indicating whether subvolume operations are enabled on the volume
func (o LookupVolumeResultOutput) EnableSubvolumes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.EnableSubvolumes }).(pulumi.StringPtrOutput)
}

// Specifies if the volume is encrypted or not. Only available on volumes created or updated after 2022-01-01.
func (o LookupVolumeResultOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVolumeResult) bool { return v.Encrypted }).(pulumi.BoolOutput)
}

// Source of key used to encrypt data in volume. Applicable if NetApp account has encryption.keySource = 'Microsoft.KeyVault'. Possible values (case-insensitive) are: 'Microsoft.NetApp, Microsoft.KeyVault'
func (o LookupVolumeResultOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVolumeResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Set of export policy rules
func (o LookupVolumeResultOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseExportPolicy { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

// Flag indicating whether file access logs are enabled for the volume, based on active diagnostic settings present on the volume.
func (o LookupVolumeResultOutput) FileAccessLogs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.FileAccessLogs }).(pulumi.StringOutput)
}

// Unique FileSystem Identifier.
func (o LookupVolumeResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies if default quota is enabled for the volume.
func (o LookupVolumeResultOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether volume is a Large Volume or Regular Volume.
func (o LookupVolumeResultOutput) IsLargeVolume() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsLargeVolume }).(pulumi.BoolPtrOutput)
}

// Restoring
func (o LookupVolumeResultOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
func (o LookupVolumeResultOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

// The resource ID of private endpoint for KeyVault. It must reside in the same VNET as the volume. Only applicable if encryptionKeySource = 'Microsoft.KeyVault'.
func (o LookupVolumeResultOutput) KeyVaultPrivateEndpointResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.KeyVaultPrivateEndpointResourceId }).(pulumi.StringPtrOutput)
}

// Specifies whether LDAP is enabled or not for a given NFS volume.
func (o LookupVolumeResultOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum number of files allowed. Needs a service request in order to be changed. Only allowed to be changed if volume quota is more than 4TiB.
func (o LookupVolumeResultOutput) MaximumNumberOfFiles() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.MaximumNumberOfFiles }).(pulumi.Float64Output)
}

// List of mount targets
func (o LookupVolumeResultOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []MountTargetPropertiesResponse { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

// The name of the resource
func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Basic network, or Standard features available to the volume.
func (o LookupVolumeResultOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

// Network Sibling Set ID for the the group of volumes sharing networking resources.
func (o LookupVolumeResultOutput) NetworkSiblingSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.NetworkSiblingSetId }).(pulumi.StringOutput)
}

// Id of the snapshot or backup that the volume is restored from.
func (o LookupVolumeResultOutput) OriginatingResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.OriginatingResourceId }).(pulumi.StringOutput)
}

// Application specific placement rules for the particular volume
func (o LookupVolumeResultOutput) PlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []PlacementKeyValuePairsResponse { return v.PlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

// Set of protocol types, default NFSv3, CIFS for SMB protocol
func (o LookupVolumeResultOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

// The availability zone where the volume is provisioned. This refers to the logical availability zone where the volume resides.
func (o LookupVolumeResultOutput) ProvisionedAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisionedAvailabilityZone }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Proximity placement group associated with the volume
func (o LookupVolumeResultOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
func (o LookupVolumeResultOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

// The service level of the file system
func (o LookupVolumeResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

// Enables access based enumeration share property for SMB Shares. Only applicable for SMB/DualProtocol volume
func (o LookupVolumeResultOutput) SmbAccessBasedEnumeration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SmbAccessBasedEnumeration }).(pulumi.StringPtrOutput)
}

// Enables continuously available share property for smb volume. Only applicable for SMB volume
func (o LookupVolumeResultOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
func (o LookupVolumeResultOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

// Enables non browsable property for SMB Shares. Only applicable for SMB/DualProtocol volume
func (o LookupVolumeResultOutput) SmbNonBrowsable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SmbNonBrowsable }).(pulumi.StringPtrOutput)
}

// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (defaults to true).
func (o LookupVolumeResultOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

// UUID v4 or resource identifier used to identify the Snapshot.
func (o LookupVolumeResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// Provides storage to network proximity information for the volume.
func (o LookupVolumeResultOutput) StorageToNetworkProximity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.StorageToNetworkProximity }).(pulumi.StringOutput)
}

// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
func (o LookupVolumeResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// T2 network information
func (o LookupVolumeResultOutput) T2Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.T2Network }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeResultOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
func (o LookupVolumeResultOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB, 500Tib for LargeVolume. Specified in bytes.
func (o LookupVolumeResultOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.UsageThreshold }).(pulumi.Float64Output)
}

// Volume Group Name
func (o LookupVolumeResultOutput) VolumeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeGroupName }).(pulumi.StringOutput)
}

// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
func (o LookupVolumeResultOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
func (o LookupVolumeResultOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

// Availability Zone
func (o LookupVolumeResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
