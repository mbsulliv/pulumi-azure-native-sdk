// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL managed instance.
type ManagedInstance struct {
	pulumi.CustomResourceState

	// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// The Azure Active Directory administrator of the instance. This can only be used at instance create time. If used for instance update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators ManagedInstanceExternalAdministratorResponsePtrOutput `pulumi:"administrators"`
	// The managed instance's authentication metadata lookup mode.
	AuthenticationMetadata pulumi.StringPtrOutput `pulumi:"authenticationMetadata"`
	// Collation of the managed instance.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// Specifies the point in time (ISO8601 format) of the Managed Instance creation.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The storage account type used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	CurrentBackupStorageRedundancy pulumi.StringOutput `pulumi:"currentBackupStorageRedundancy"`
	// Specifies the internal format of instance databases specific to the SQL engine version.
	DatabaseFormat pulumi.StringPtrOutput `pulumi:"databaseFormat"`
	// The Dns Zone that the managed instance is in.
	DnsZone pulumi.StringOutput `pulumi:"dnsZone"`
	// Status of external governance.
	ExternalGovernanceStatus pulumi.StringOutput `pulumi:"externalGovernanceStatus"`
	// The fully qualified domain name of the managed instance.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// Hybrid secondary usage. Possible values are 'Active' (default value) and 'Passive' (customer uses the secondary as Passive DR).
	HybridSecondaryUsage pulumi.StringPtrOutput `pulumi:"hybridSecondaryUsage"`
	// Hybrid secondary usage detected. Possible values are 'Active' (customer does not meet the requirements to use the secondary as Passive DR) and 'Passive' (customer meets the requirements to use the secondary as Passive DR).
	HybridSecondaryUsageDetected pulumi.StringOutput `pulumi:"hybridSecondaryUsageDetected"`
	// The Azure Active Directory identity of the managed instance.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// The Id of the instance pool this managed server belongs to.
	InstancePoolId pulumi.StringPtrOutput `pulumi:"instancePoolId"`
	// Whether or not this is a GPv2 variant of General Purpose edition.
	IsGeneralPurposeV2 pulumi.BoolPtrOutput `pulumi:"isGeneralPurposeV2"`
	// A CMK URI of the key to use for encryption.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies maintenance configuration id to apply to this managed instance.
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Weather or not Managed Instance is freemium.
	PricingModel pulumi.StringPtrOutput `pulumi:"pricingModel"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId pulumi.StringPtrOutput `pulumi:"primaryUserAssignedIdentityId"`
	// List of private endpoint connections on a managed instance.
	PrivateEndpointConnections ManagedInstancePecPropertyResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of managed instance.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Connection type used for connecting to the instance.
	ProxyOverride pulumi.StringPtrOutput `pulumi:"proxyOverride"`
	// Whether or not the public data endpoint is enabled.
	PublicDataEndpointEnabled pulumi.BoolPtrOutput `pulumi:"publicDataEndpointEnabled"`
	// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	RequestedBackupStorageRedundancy pulumi.StringPtrOutput `pulumi:"requestedBackupStorageRedundancy"`
	// The managed instance's service principal.
	ServicePrincipal ServicePrincipalResponsePtrOutput `pulumi:"servicePrincipal"`
	// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The state of the managed instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Storage IOps. Minimum value: 120. Maximum value: 120000. Increments of 1 IOps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageIOps pulumi.IntPtrOutput `pulumi:"storageIOps"`
	// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageSizeInGB pulumi.IntPtrOutput `pulumi:"storageSizeInGB"`
	// Storage throughput in MBps. Minimum value: 25. Maximum value: 4000. Increments of 1 MBps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageThroughputMBps pulumi.IntPtrOutput `pulumi:"storageThroughputMBps"`
	// Subnet resource ID for the managed instance.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Id of the timezone. Allowed values are timezones supported by Windows.
	// Windows keeps details on supported timezones, including the id, in registry under
	// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
	// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
	// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
	TimezoneId pulumi.StringPtrOutput `pulumi:"timezoneId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
	VCores pulumi.IntPtrOutput `pulumi:"vCores"`
	// Virtual cluster resource id for the Managed Instance.
	VirtualClusterId pulumi.StringOutput `pulumi:"virtualClusterId"`
	// Whether or not the multi-az is enabled.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewManagedInstance registers a new resource with the given unique name, arguments, and options.
func NewManagedInstance(ctx *pulumi.Context,
	name string, args *ManagedInstanceArgs, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedInstance
	err := ctx.RegisterResource("azure-native:sql/v20230501preview:ManagedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstance gets an existing ManagedInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceState, opts ...pulumi.ResourceOption) (*ManagedInstance, error) {
	var resource ManagedInstance
	err := ctx.ReadResource("azure-native:sql/v20230501preview:ManagedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstance resources.
type managedInstanceState struct {
}

type ManagedInstanceState struct {
}

func (ManagedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceState)(nil)).Elem()
}

type managedInstanceArgs struct {
	// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The administrator login password (required for managed instance creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The Azure Active Directory administrator of the instance. This can only be used at instance create time. If used for instance update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators *ManagedInstanceExternalAdministrator `pulumi:"administrators"`
	// The managed instance's authentication metadata lookup mode.
	AuthenticationMetadata *string `pulumi:"authenticationMetadata"`
	// Collation of the managed instance.
	Collation *string `pulumi:"collation"`
	// Specifies the internal format of instance databases specific to the SQL engine version.
	DatabaseFormat *string `pulumi:"databaseFormat"`
	// The resource id of another managed instance whose DNS zone this managed instance will share after creation.
	DnsZonePartner *string `pulumi:"dnsZonePartner"`
	// Hybrid secondary usage. Possible values are 'Active' (default value) and 'Passive' (customer uses the secondary as Passive DR).
	HybridSecondaryUsage *string `pulumi:"hybridSecondaryUsage"`
	// The Azure Active Directory identity of the managed instance.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The Id of the instance pool this managed server belongs to.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// Whether or not this is a GPv2 variant of General Purpose edition.
	IsGeneralPurposeV2 *bool `pulumi:"isGeneralPurposeV2"`
	// A CMK URI of the key to use for encryption.
	KeyId *string `pulumi:"keyId"`
	// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Specifies maintenance configuration id to apply to this managed instance.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Specifies the mode of database creation.
	//
	// Default: Regular instance creation.
	//
	// Restore: Creates an instance by restoring a set of backups to specific point in time. RestorePointInTime and SourceManagedInstanceId must be specified.
	ManagedInstanceCreateMode *string `pulumi:"managedInstanceCreateMode"`
	// The name of the managed instance.
	ManagedInstanceName *string `pulumi:"managedInstanceName"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Weather or not Managed Instance is freemium.
	PricingModel *string `pulumi:"pricingModel"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `pulumi:"primaryUserAssignedIdentityId"`
	// Connection type used for connecting to the instance.
	ProxyOverride *string `pulumi:"proxyOverride"`
	// Whether or not the public data endpoint is enabled.
	PublicDataEndpointEnabled *bool `pulumi:"publicDataEndpointEnabled"`
	// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	RequestedBackupStorageRedundancy *string `pulumi:"requestedBackupStorageRedundancy"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The managed instance's service principal.
	ServicePrincipal *ServicePrincipal `pulumi:"servicePrincipal"`
	// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
	Sku *Sku `pulumi:"sku"`
	// The resource identifier of the source managed instance associated with create operation of this instance.
	SourceManagedInstanceId *string `pulumi:"sourceManagedInstanceId"`
	// Storage IOps. Minimum value: 120. Maximum value: 120000. Increments of 1 IOps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageIOps *int `pulumi:"storageIOps"`
	// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageSizeInGB *int `pulumi:"storageSizeInGB"`
	// Storage throughput in MBps. Minimum value: 25. Maximum value: 4000. Increments of 1 MBps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageThroughputMBps *int `pulumi:"storageThroughputMBps"`
	// Subnet resource ID for the managed instance.
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Id of the timezone. Allowed values are timezones supported by Windows.
	// Windows keeps details on supported timezones, including the id, in registry under
	// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
	// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
	// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
	TimezoneId *string `pulumi:"timezoneId"`
	// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
	VCores *int `pulumi:"vCores"`
	// Whether or not the multi-az is enabled.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a ManagedInstance resource.
type ManagedInstanceArgs struct {
	// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrInput
	// The administrator login password (required for managed instance creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// The Azure Active Directory administrator of the instance. This can only be used at instance create time. If used for instance update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
	Administrators ManagedInstanceExternalAdministratorPtrInput
	// The managed instance's authentication metadata lookup mode.
	AuthenticationMetadata pulumi.StringPtrInput
	// Collation of the managed instance.
	Collation pulumi.StringPtrInput
	// Specifies the internal format of instance databases specific to the SQL engine version.
	DatabaseFormat pulumi.StringPtrInput
	// The resource id of another managed instance whose DNS zone this managed instance will share after creation.
	DnsZonePartner pulumi.StringPtrInput
	// Hybrid secondary usage. Possible values are 'Active' (default value) and 'Passive' (customer uses the secondary as Passive DR).
	HybridSecondaryUsage pulumi.StringPtrInput
	// The Azure Active Directory identity of the managed instance.
	Identity ResourceIdentityPtrInput
	// The Id of the instance pool this managed server belongs to.
	InstancePoolId pulumi.StringPtrInput
	// Whether or not this is a GPv2 variant of General Purpose edition.
	IsGeneralPurposeV2 pulumi.BoolPtrInput
	// A CMK URI of the key to use for encryption.
	KeyId pulumi.StringPtrInput
	// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Specifies maintenance configuration id to apply to this managed instance.
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: Regular instance creation.
	//
	// Restore: Creates an instance by restoring a set of backups to specific point in time. RestorePointInTime and SourceManagedInstanceId must be specified.
	ManagedInstanceCreateMode pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringPtrInput
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrInput
	// Weather or not Managed Instance is freemium.
	PricingModel pulumi.StringPtrInput
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId pulumi.StringPtrInput
	// Connection type used for connecting to the instance.
	ProxyOverride pulumi.StringPtrInput
	// Whether or not the public data endpoint is enabled.
	PublicDataEndpointEnabled pulumi.BoolPtrInput
	// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	RequestedBackupStorageRedundancy pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The managed instance's service principal.
	ServicePrincipal ServicePrincipalPtrInput
	// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
	Sku SkuPtrInput
	// The resource identifier of the source managed instance associated with create operation of this instance.
	SourceManagedInstanceId pulumi.StringPtrInput
	// Storage IOps. Minimum value: 120. Maximum value: 120000. Increments of 1 IOps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageIOps pulumi.IntPtrInput
	// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageSizeInGB pulumi.IntPtrInput
	// Storage throughput in MBps. Minimum value: 25. Maximum value: 4000. Increments of 1 MBps allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageThroughputMBps pulumi.IntPtrInput
	// Subnet resource ID for the managed instance.
	SubnetId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Id of the timezone. Allowed values are timezones supported by Windows.
	// Windows keeps details on supported timezones, including the id, in registry under
	// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
	// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
	// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
	TimezoneId pulumi.StringPtrInput
	// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
	VCores pulumi.IntPtrInput
	// Whether or not the multi-az is enabled.
	ZoneRedundant pulumi.BoolPtrInput
}

func (ManagedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceArgs)(nil)).Elem()
}

type ManagedInstanceInput interface {
	pulumi.Input

	ToManagedInstanceOutput() ManagedInstanceOutput
	ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput
}

func (*ManagedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstance)(nil)).Elem()
}

func (i *ManagedInstance) ToManagedInstanceOutput() ManagedInstanceOutput {
	return i.ToManagedInstanceOutputWithContext(context.Background())
}

func (i *ManagedInstance) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceOutput)
}

type ManagedInstanceOutput struct{ *pulumi.OutputState }

func (ManagedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstance)(nil)).Elem()
}

func (o ManagedInstanceOutput) ToManagedInstanceOutput() ManagedInstanceOutput {
	return o
}

func (o ManagedInstanceOutput) ToManagedInstanceOutputWithContext(ctx context.Context) ManagedInstanceOutput {
	return o
}

// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
func (o ManagedInstanceOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory administrator of the instance. This can only be used at instance create time. If used for instance update, it will be ignored or it will result in an error. For updates individual APIs will need to be used.
func (o ManagedInstanceOutput) Administrators() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) ManagedInstanceExternalAdministratorResponsePtrOutput {
		return v.Administrators
	}).(ManagedInstanceExternalAdministratorResponsePtrOutput)
}

// The managed instance's authentication metadata lookup mode.
func (o ManagedInstanceOutput) AuthenticationMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.AuthenticationMetadata }).(pulumi.StringPtrOutput)
}

// Collation of the managed instance.
func (o ManagedInstanceOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

// Specifies the point in time (ISO8601 format) of the Managed Instance creation.
func (o ManagedInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The storage account type used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
func (o ManagedInstanceOutput) CurrentBackupStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.CurrentBackupStorageRedundancy }).(pulumi.StringOutput)
}

// Specifies the internal format of instance databases specific to the SQL engine version.
func (o ManagedInstanceOutput) DatabaseFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.DatabaseFormat }).(pulumi.StringPtrOutput)
}

// The Dns Zone that the managed instance is in.
func (o ManagedInstanceOutput) DnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.DnsZone }).(pulumi.StringOutput)
}

// Status of external governance.
func (o ManagedInstanceOutput) ExternalGovernanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.ExternalGovernanceStatus }).(pulumi.StringOutput)
}

// The fully qualified domain name of the managed instance.
func (o ManagedInstanceOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// Hybrid secondary usage. Possible values are 'Active' (default value) and 'Passive' (customer uses the secondary as Passive DR).
func (o ManagedInstanceOutput) HybridSecondaryUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.HybridSecondaryUsage }).(pulumi.StringPtrOutput)
}

// Hybrid secondary usage detected. Possible values are 'Active' (customer does not meet the requirements to use the secondary as Passive DR) and 'Passive' (customer meets the requirements to use the secondary as Passive DR).
func (o ManagedInstanceOutput) HybridSecondaryUsageDetected() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.HybridSecondaryUsageDetected }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the managed instance.
func (o ManagedInstanceOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The Id of the instance pool this managed server belongs to.
func (o ManagedInstanceOutput) InstancePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.InstancePoolId }).(pulumi.StringPtrOutput)
}

// Whether or not this is a GPv2 variant of General Purpose edition.
func (o ManagedInstanceOutput) IsGeneralPurposeV2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.BoolPtrOutput { return v.IsGeneralPurposeV2 }).(pulumi.BoolPtrOutput)
}

// A CMK URI of the key to use for encryption.
func (o ManagedInstanceOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
func (o ManagedInstanceOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ManagedInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Specifies maintenance configuration id to apply to this managed instance.
func (o ManagedInstanceOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
func (o ManagedInstanceOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ManagedInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Weather or not Managed Instance is freemium.
func (o ManagedInstanceOutput) PricingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.PricingModel }).(pulumi.StringPtrOutput)
}

// The resource id of a user assigned identity to be used by default.
func (o ManagedInstanceOutput) PrimaryUserAssignedIdentityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.PrimaryUserAssignedIdentityId }).(pulumi.StringPtrOutput)
}

// List of private endpoint connections on a managed instance.
func (o ManagedInstanceOutput) PrivateEndpointConnections() ManagedInstancePecPropertyResponseArrayOutput {
	return o.ApplyT(func(v *ManagedInstance) ManagedInstancePecPropertyResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(ManagedInstancePecPropertyResponseArrayOutput)
}

// Provisioning state of managed instance.
func (o ManagedInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Connection type used for connecting to the instance.
func (o ManagedInstanceOutput) ProxyOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.ProxyOverride }).(pulumi.StringPtrOutput)
}

// Whether or not the public data endpoint is enabled.
func (o ManagedInstanceOutput) PublicDataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.BoolPtrOutput { return v.PublicDataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
func (o ManagedInstanceOutput) RequestedBackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.RequestedBackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

// The managed instance's service principal.
func (o ManagedInstanceOutput) ServicePrincipal() ServicePrincipalResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) ServicePrincipalResponsePtrOutput { return v.ServicePrincipal }).(ServicePrincipalResponsePtrOutput)
}

// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
func (o ManagedInstanceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstance) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the managed instance.
func (o ManagedInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Storage IOps. Minimum value: 120. Maximum value: 120000. Increments of 1 IOps allowed only. Maximum value depends on the selected hardware family and number of vCores.
func (o ManagedInstanceOutput) StorageIOps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.StorageIOps }).(pulumi.IntPtrOutput)
}

// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
func (o ManagedInstanceOutput) StorageSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.StorageSizeInGB }).(pulumi.IntPtrOutput)
}

// Storage throughput in MBps. Minimum value: 25. Maximum value: 4000. Increments of 1 MBps allowed only. Maximum value depends on the selected hardware family and number of vCores.
func (o ManagedInstanceOutput) StorageThroughputMBps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.StorageThroughputMBps }).(pulumi.IntPtrOutput)
}

// Subnet resource ID for the managed instance.
func (o ManagedInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o ManagedInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Id of the timezone. Allowed values are timezones supported by Windows.
// Windows keeps details on supported timezones, including the id, in registry under
// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
func (o ManagedInstanceOutput) TimezoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringPtrOutput { return v.TimezoneId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ManagedInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
func (o ManagedInstanceOutput) VCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.IntPtrOutput { return v.VCores }).(pulumi.IntPtrOutput)
}

// Virtual cluster resource id for the Managed Instance.
func (o ManagedInstanceOutput) VirtualClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.StringOutput { return v.VirtualClusterId }).(pulumi.StringOutput)
}

// Whether or not the multi-az is enabled.
func (o ManagedInstanceOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstance) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceOutput{})
}
