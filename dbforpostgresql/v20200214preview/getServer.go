// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200214preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a server.
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20200214preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server.
type LookupServerResult struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// availability Zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Status showing whether the data encryption is enabled with customer-managed keys.
	ByokEnforcement          string                                            `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments *ServerPropertiesResponseDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	// The display name of a server.
	DisplayName *string `pulumi:"displayName"`
	// The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate string `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName string `pulumi:"fullyQualifiedDomainName"`
	// stand by count value can be either enabled or disabled
	HaEnabled *string `pulumi:"haEnabled"`
	// A state of a HA server that is visible to user.
	HaState string `pulumi:"haState"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the server.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The log backup storage sku of the server.
	LogBackupStorageSku *string `pulumi:"logBackupStorageSku"`
	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The minor version of the server.
	MinorVersion string `pulumi:"minorVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	PointInTimeUTC          *string                                          `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments *ServerPropertiesResponsePrivateDnsZoneArguments `pulumi:"privateDnsZoneArguments"`
	// public network access is enabled or not
	PublicNetworkAccess string `pulumi:"publicNetworkAccess"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// The resource group name of source serve PostgreSQL server name to restore from.
	SourceResourceGroupName *string `pulumi:"sourceResourceGroupName"`
	// The source PostgreSQL server name to restore from.
	SourceServerName *string `pulumi:"sourceServerName"`
	// The subscription id of source serve PostgreSQL server name to restore from.
	SourceSubscriptionId *string `pulumi:"sourceSubscriptionId"`
	// availability Zone information of the server.
	StandbyAvailabilityZone string `pulumi:"standbyAvailabilityZone"`
	// The number of standbys.
	StandbyCount *int `pulumi:"standbyCount"`
	// A state of a server that is visible to user.
	State string `pulumi:"state"`
	// Storage profile of a server.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// PostgreSQL Server version.
	Version *string `pulumi:"version"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

// Represents a server.
type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// availability Zone information of the server.
func (o LookupServerResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Status showing whether the data encryption is enabled with customer-managed keys.
func (o LookupServerResultOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) DelegatedSubnetArguments() ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerPropertiesResponseDelegatedSubnetArguments {
		return v.DelegatedSubnetArguments
	}).(ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

// The display name of a server.
func (o LookupServerResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The earliest restore point time (ISO8601 format) for server.
func (o LookupServerResultOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

// The fully qualified domain name of a server.
func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// stand by count value can be either enabled or disabled
func (o LookupServerResultOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

// A state of a HA server that is visible to user.
func (o LookupServerResultOutput) HaState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.HaState }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Azure Active Directory identity of the server.
func (o LookupServerResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The log backup storage sku of the server.
func (o LookupServerResultOutput) LogBackupStorageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.LogBackupStorageSku }).(pulumi.StringPtrOutput)
}

// Maintenance window of a server.
func (o LookupServerResultOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

// The minor version of the server.
func (o LookupServerResultOutput) MinorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.MinorVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Restore point creation time (ISO8601 format), specifying the time to restore from.
func (o LookupServerResultOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) PrivateDnsZoneArguments() ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerPropertiesResponsePrivateDnsZoneArguments {
		return v.PrivateDnsZoneArguments
	}).(ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

// public network access is enabled or not
func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// The SKU (pricing tier) of the server.
func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The resource group name of source serve PostgreSQL server name to restore from.
func (o LookupServerResultOutput) SourceResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceResourceGroupName }).(pulumi.StringPtrOutput)
}

// The source PostgreSQL server name to restore from.
func (o LookupServerResultOutput) SourceServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceServerName }).(pulumi.StringPtrOutput)
}

// The subscription id of source serve PostgreSQL server name to restore from.
func (o LookupServerResultOutput) SourceSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceSubscriptionId }).(pulumi.StringPtrOutput)
}

// availability Zone information of the server.
func (o LookupServerResultOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
}

// The number of standbys.
func (o LookupServerResultOutput) StandbyCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.StandbyCount }).(pulumi.IntPtrOutput)
}

// A state of a server that is visible to user.
func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

// Storage profile of a server.
func (o LookupServerResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags.
func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// PostgreSQL Server version.
func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
