// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a managed instance.
func LookupManagedInstance(ctx *pulumi.Context, args *LookupManagedInstanceArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedInstanceResult
	err := ctx.Invoke("azure-native:sql/v20211101:getManagedInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceArgs struct {
	// The child resources to include in the response.
	Expand *string `pulumi:"expand"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure SQL managed instance.
type LookupManagedInstanceResult struct {
	// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The Azure Active Directory administrator of the server.
	Administrators *ManagedInstanceExternalAdministratorResponse `pulumi:"administrators"`
	// Collation of the managed instance.
	Collation *string `pulumi:"collation"`
	// The storage account type used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	CurrentBackupStorageRedundancy string `pulumi:"currentBackupStorageRedundancy"`
	// The Dns Zone that the managed instance is in.
	DnsZone string `pulumi:"dnsZone"`
	// The fully qualified domain name of the managed instance.
	FullyQualifiedDomainName string `pulumi:"fullyQualifiedDomainName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the managed instance.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// The Id of the instance pool this managed server belongs to.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// A CMK URI of the key to use for encryption.
	KeyId *string `pulumi:"keyId"`
	// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location string `pulumi:"location"`
	// Specifies maintenance configuration id to apply to this managed instance.
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name string `pulumi:"name"`
	// The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `pulumi:"primaryUserAssignedIdentityId"`
	// List of private endpoint connections on a managed instance.
	PrivateEndpointConnections []ManagedInstancePecPropertyResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                               `pulumi:"provisioningState"`
	// Connection type used for connecting to the instance.
	ProxyOverride *string `pulumi:"proxyOverride"`
	// Whether or not the public data endpoint is enabled.
	PublicDataEndpointEnabled *bool `pulumi:"publicDataEndpointEnabled"`
	// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
	RequestedBackupStorageRedundancy *string `pulumi:"requestedBackupStorageRedundancy"`
	// The managed instance's service principal.
	ServicePrincipal *ServicePrincipalResponse `pulumi:"servicePrincipal"`
	// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
	Sku *SkuResponse `pulumi:"sku"`
	// The state of the managed instance.
	State string `pulumi:"state"`
	// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
	StorageSizeInGB *int `pulumi:"storageSizeInGB"`
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
	// Resource type.
	Type string `pulumi:"type"`
	// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
	VCores *int `pulumi:"vCores"`
	// Whether or not the multi-az is enabled.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

func LookupManagedInstanceOutput(ctx *pulumi.Context, args LookupManagedInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceResult, error) {
			args := v.(LookupManagedInstanceArgs)
			r, err := LookupManagedInstance(ctx, &args, opts...)
			var s LookupManagedInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceResultOutput)
}

type LookupManagedInstanceOutputArgs struct {
	// The child resources to include in the response.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceArgs)(nil)).Elem()
}

// An Azure SQL managed instance.
type LookupManagedInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceResult)(nil)).Elem()
}

func (o LookupManagedInstanceResultOutput) ToLookupManagedInstanceResultOutput() LookupManagedInstanceResultOutput {
	return o
}

func (o LookupManagedInstanceResultOutput) ToLookupManagedInstanceResultOutputWithContext(ctx context.Context) LookupManagedInstanceResultOutput {
	return o
}

// Administrator username for the managed instance. Can only be specified when the managed instance is being created (and is required for creation).
func (o LookupManagedInstanceResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory administrator of the server.
func (o LookupManagedInstanceResultOutput) Administrators() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *ManagedInstanceExternalAdministratorResponse {
		return v.Administrators
	}).(ManagedInstanceExternalAdministratorResponsePtrOutput)
}

// Collation of the managed instance.
func (o LookupManagedInstanceResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

// The storage account type used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
func (o LookupManagedInstanceResultOutput) CurrentBackupStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.CurrentBackupStorageRedundancy }).(pulumi.StringOutput)
}

// The Dns Zone that the managed instance is in.
func (o LookupManagedInstanceResultOutput) DnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.DnsZone }).(pulumi.StringOutput)
}

// The fully qualified domain name of the managed instance.
func (o LookupManagedInstanceResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupManagedInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the managed instance.
func (o LookupManagedInstanceResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The Id of the instance pool this managed server belongs to.
func (o LookupManagedInstanceResultOutput) InstancePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.InstancePoolId }).(pulumi.StringPtrOutput)
}

// A CMK URI of the key to use for encryption.
func (o LookupManagedInstanceResultOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.KeyId }).(pulumi.StringPtrOutput)
}

// The license type. Possible values are 'LicenseIncluded' (regular price inclusive of a new SQL license) and 'BasePrice' (discounted AHB price for bringing your own SQL licenses).
func (o LookupManagedInstanceResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupManagedInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Specifies maintenance configuration id to apply to this managed instance.
func (o LookupManagedInstanceResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Minimal TLS version. Allowed values: 'None', '1.0', '1.1', '1.2'
func (o LookupManagedInstanceResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupManagedInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource id of a user assigned identity to be used by default.
func (o LookupManagedInstanceResultOutput) PrimaryUserAssignedIdentityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.PrimaryUserAssignedIdentityId }).(pulumi.StringPtrOutput)
}

// List of private endpoint connections on a managed instance.
func (o LookupManagedInstanceResultOutput) PrivateEndpointConnections() ManagedInstancePecPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) []ManagedInstancePecPropertyResponse {
		return v.PrivateEndpointConnections
	}).(ManagedInstancePecPropertyResponseArrayOutput)
}

func (o LookupManagedInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Connection type used for connecting to the instance.
func (o LookupManagedInstanceResultOutput) ProxyOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.ProxyOverride }).(pulumi.StringPtrOutput)
}

// Whether or not the public data endpoint is enabled.
func (o LookupManagedInstanceResultOutput) PublicDataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *bool { return v.PublicDataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// The storage account type to be used to store backups for this instance. The options are Local (LocallyRedundantStorage), Zone (ZoneRedundantStorage), Geo (GeoRedundantStorage) and GeoZone(GeoZoneRedundantStorage)
func (o LookupManagedInstanceResultOutput) RequestedBackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.RequestedBackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

// The managed instance's service principal.
func (o LookupManagedInstanceResultOutput) ServicePrincipal() ServicePrincipalResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *ServicePrincipalResponse { return v.ServicePrincipal }).(ServicePrincipalResponsePtrOutput)
}

// Managed instance SKU. Allowed values for sku.name: GP_Gen5, GP_G8IM, GP_G8IH, BC_Gen5, BC_G8IM, BC_G8IH
func (o LookupManagedInstanceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The state of the managed instance.
func (o LookupManagedInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// Storage size in GB. Minimum value: 32. Maximum value: 16384. Increments of 32 GB allowed only. Maximum value depends on the selected hardware family and number of vCores.
func (o LookupManagedInstanceResultOutput) StorageSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *int { return v.StorageSizeInGB }).(pulumi.IntPtrOutput)
}

// Subnet resource ID for the managed instance.
func (o LookupManagedInstanceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupManagedInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Id of the timezone. Allowed values are timezones supported by Windows.
// Windows keeps details on supported timezones, including the id, in registry under
// KEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones.
// You can get those registry values via SQL Server by querying SELECT name AS timezone_id FROM sys.time_zone_info.
// List of Ids can also be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
// An example of valid timezone id is "Pacific Standard Time" or "W. Europe Standard Time".
func (o LookupManagedInstanceResultOutput) TimezoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.TimezoneId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupManagedInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The number of vCores. Allowed values: 8, 16, 24, 32, 40, 64, 80.
func (o LookupManagedInstanceResultOutput) VCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *int { return v.VCores }).(pulumi.IntPtrOutput)
}

// Whether or not the multi-az is enabled.
func (o LookupManagedInstanceResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceResultOutput{})
}
