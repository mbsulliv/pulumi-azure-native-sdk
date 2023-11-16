// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210615privatepreview

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
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210615privatepreview:getServer", args, &rv, opts...)
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
	// availability zone information of the server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Backup properties of a server.
	Backup *BackupResponse `pulumi:"backup"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName string `pulumi:"fullyQualifiedDomainName"`
	// High availability properties of a server.
	HighAvailability *HighAvailabilityResponse `pulumi:"highAvailability"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The Azure Active Directory identity of the server.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The minor version of the server.
	MinorVersion string `pulumi:"minorVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network properties of a server.
	Network *NetworkResponse `pulumi:"network"`
	// The SKU (pricing tier) of the server.
	Sku *SkuResponse `pulumi:"sku"`
	// A state of a server that is visible to user.
	State string `pulumi:"state"`
	// Storage properties of a server.
	Storage *StorageResponse `pulumi:"storage"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
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

// availability zone information of the server.
func (o LookupServerResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Backup properties of a server.
func (o LookupServerResultOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *BackupResponse { return v.Backup }).(BackupResponsePtrOutput)
}

// The fully qualified domain name of a server.
func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

// High availability properties of a server.
func (o LookupServerResultOutput) HighAvailability() HighAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *HighAvailabilityResponse { return v.HighAvailability }).(HighAvailabilityResponsePtrOutput)
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

// Maintenance window properties of a server.
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

// Network properties of a server.
func (o LookupServerResultOutput) Network() NetworkResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *NetworkResponse { return v.Network }).(NetworkResponsePtrOutput)
}

// The SKU (pricing tier) of the server.
func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// A state of a server that is visible to user.
func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

// Storage properties of a server.
func (o LookupServerResultOutput) Storage() StorageResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageResponse { return v.Storage }).(StorageResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o LookupServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
